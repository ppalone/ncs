package ncs

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

// constants
const (
	baseURL     string = "https://ncs.io"
	searchURL   string = "music-search"
	downloadURL string = "track/download"
)

// NCS Client
type Client struct {
	httpClient *http.Client
}

// NewClient returns a new NCS client
func NewClient(httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = &http.Client{}
	}

	return &Client{httpClient}
}

// Search music with options
func (c *Client) Search(ctx context.Context, q string, opts ...SearchOption) (Result, error) {
	// Search query can be empty apparently
	q = strings.TrimSpace(q)

	filters := defaultSearchOptions()
	for _, opt := range opts {
		opt(filters)
	}

	return c.search(ctx, q, filters)
}

// Releases returns the latest NCS releases
func (c *Client) Releases(ctx context.Context, opts ...SearchOption) (Result, error) {
	filters := defaultSearchOptions()
	for _, opt := range opts {
		opt(filters)
	}

	// releases is just normal music search without any search filters and query
	return c.search(ctx, "", filters)
}

// GetSongById returns the song info by id
func (c *Client) GetSongById(ctx context.Context, id string) (Song, error) {
	id = strings.TrimSpace(id)
	if len(id) == 0 {
		return Song{}, fmt.Errorf("id cannot be empty")
	}

	req, err := makeRequest(ctx, http.MethodGet, fmt.Sprintf("%s/%s", baseURL, id), "")
	if err != nil {
		return Song{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return Song{}, err
	}
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return Song{}, err
	}

	section := doc.Find(".player-nest")
	if section.Length() == 0 {
		return Song{}, fmt.Errorf("invalid id")
	}

	info := section.Find(".buttons a").First()
	if info.Length() == 0 {
		return Song{}, fmt.Errorf("unable to get song info")
	}

	song := Song{
		Id:     id,
		Title:  info.AttrOr("data-track", ""),
		WebURL: fmt.Sprintf("%s/%s", baseURL, id),
		Genres: strings.Split(info.AttrOr("data-genre", ""), ", "),
	}

	song.MediaURL = section.Find("#player").First().AttrOr("data-url", "")
	song.Artists = extractArtistsFromSelection(section.Find("h2 a"))

	versions := make([]string, 0)
	downloads := make([]Download, 0)
	section.Find(".buttons a.btn").Each(func(i int, s *goquery.Selection) {
		v, ok := s.Attr("data-version")
		if ok {
			versions = append(versions, v)
			if downloadURL, hasDownloadURL := s.Attr("href"); hasDownloadURL {
				downloads = append(downloads, Download{
					Version:     v,
					DownloadURL: fmt.Sprintf("%s%s", baseURL, downloadURL),
				})
			}
		}
	})
	song.Versions = versions
	song.Downloads = downloads

	return song, nil
}

func makeRequest(ctx context.Context, method string, url string, params string) (*http.Request, error) {
	req, err := http.NewRequestWithContext(ctx, method, url, nil)
	if err != nil {
		return nil, err
	}
	req.URL.RawQuery = params
	return req, nil
}

func (c *Client) search(ctx context.Context, q string, opts *searchOptions) (Result, error) {
	params := &url.Values{}

	// search query
	params.Set("q", q)

	// genre
	if g, ok := GenreMap[opts.genre]; ok {
		params.Set("genre", strconv.Itoa(g))
	}

	// mood
	if m, ok := MoodMap[opts.mood]; ok {
		params.Set("mood", strconv.Itoa(m))
	}

	params.Set("version", string(opts.version))
	params.Set("page", strconv.Itoa(opts.page))

	req, err := makeRequest(ctx, http.MethodGet, fmt.Sprintf("%s/%s", baseURL, searchURL), params.Encode())
	if err != nil {
		return Result{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return Result{}, err
	}
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return Result{}, err
	}

	rows := doc.Find(".tablesorter tbody tr")
	if rows.Length() == 0 {
		return Result{
			Size:    0,
			Page:    opts.page,
			Songs:   make([]Song, 0),
			HasNext: false,
		}, nil
	}

	songs := make([]Song, 0)
	rows.Each(func(i int, s *goquery.Selection) {
		var song Song
		title := s.Find("td").Eq(3).Find("a").First()
		if title.Length() == 0 {
			return
		}

		url, ok := title.Attr("href")
		if !ok {
			return
		}

		song = Song{
			Id:     strings.TrimPrefix(url, "/"),
			Title:  strings.TrimSpace(title.Text()),
			WebURL: fmt.Sprintf("%s%s", baseURL, url),
		}

		info := s.Find("td").First().Find("a").First()
		if info.Length() == 0 {
			return
		}

		song.MediaURL = info.AttrOr("data-url", "")
		song.CoverURL = info.AttrOr("data-cover", "")
		song.Genres = strings.Split(info.AttrOr("data-genre", ""), ", ")
		song.Versions = strings.Split(info.AttrOr("data-versions", ""), ", ")

		downloads := make([]Download, 0)
		if id, ok := info.Attr("data-tid"); ok {
			for _, v := range song.Versions {
				if strings.EqualFold(v, string(Regular)) {
					downloads = append(downloads, Download{
						Version:     v,
						DownloadURL: fmt.Sprintf("%s/%s/%s", baseURL, downloadURL, id),
					})
				} else if strings.EqualFold(v, string(Instrumental)) {
					downloads = append(downloads, Download{
						Version:     v,
						DownloadURL: fmt.Sprintf("%s/%s/i_%s", baseURL, downloadURL, id),
					})
				}
			}
		}
		song.Downloads = downloads

		doc, err := goquery.NewDocumentFromReader(strings.NewReader(info.AttrOr("data-artist", "")))
		if err != nil {
			return
		}

		song.Artists = extractArtistsFromSelection(doc.Find("a"))

		moods := make([]string, 0)
		s.Find("td").Eq(4).Find("a").Each(func(i int, s *goquery.Selection) {
			// Ignore, since first entry will be genre
			if i == 0 {
				return
			}

			moods = append(moods, strings.TrimSpace(s.Text()))
		})
		song.Moods = moods

		// ignore error
		t, _ := time.Parse("2 Jan 2006", strings.TrimSpace(s.Find("td").Eq(5).Text()))
		song.ReleaseDate = t

		songs = append(songs, song)
	})

	pagination := doc.Find("ul.pagination")
	if pagination.Length() == 0 {
		return Result{
			Size:    len(songs),
			Songs:   songs,
			Page:    opts.page,
			HasNext: false,
		}, nil
	}

	item := pagination.Find("li").Last()
	if item.Length() == 0 || item.HasClass("disabled") {
		return Result{
			Size:    len(songs),
			Songs:   songs,
			Page:    opts.page,
			HasNext: false,
		}, nil
	}

	return Result{
		c:    c,
		q:    q,
		opts: opts,

		Size:    len(songs),
		Songs:   songs,
		Page:    opts.page,
		HasNext: true,
	}, nil
}

func extractArtistsFromSelection(rows *goquery.Selection) []Artist {
	artists := make([]Artist, 0)

	rows.Each(func(i int, s *goquery.Selection) {
		u, ok := s.Attr("href")
		if !ok {
			return
		}

		t := strings.Split(u, "/")
		if len(t) < 3 {
			return
		}

		artists = append(artists, Artist{
			Id:        t[2],
			Name:      strings.TrimSpace(s.Text()),
			ArtistURL: fmt.Sprintf("%s%s", baseURL, u),
		})
	})

	return artists
}
