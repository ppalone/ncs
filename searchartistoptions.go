package ncs

import (
	"fmt"
	"net/url"
	"strconv"
	"time"
)

type SearchArtistSort string

// Sort types
const (
	Latest SearchArtistSort = "latest"
	AZ     SearchArtistSort = "az"
	ZA     SearchArtistSort = "za"
)

// Search Artist Options.
type searchArtistOptions struct {
	q    string
	sort SearchArtistSort
	year int
	page int
}

// SearchArtistOption
type SearchArtistOption func(opts *searchArtistOptions)

// default search artist options
func defaultSearchArtistOptions() *searchArtistOptions {
	return &searchArtistOptions{
		q:    "",
		sort: Latest,
		year: 0,
		page: 1,
	}
}

// WithSort option
func WithSort(sort SearchArtistSort) SearchArtistOption {
	return func(opts *searchArtistOptions) {
		opts.sort = sort
	}
}

// WithYear option
func WithYear(year int) SearchArtistOption {
	return func(opts *searchArtistOptions) {
		opts.year = year
	}
}

// WithSearchArtistPage option
func WithSearchArtistPage(page int) SearchArtistOption {
	return func(opts *searchArtistOptions) {
		opts.page = page
	}
}

// validates the search artists options
func (opts *searchArtistOptions) validate() error {
	if opts.year != 0 && (opts.year < 2013 || opts.year > time.Now().Year()) {
		return fmt.Errorf("invalid year")
	}

	return nil
}

// builds the query string for search artists options
func (opts *searchArtistOptions) build() string {
	params := &url.Values{}

	params.Set("q", opts.q)
	params.Set("sort", string(opts.sort))
	params.Set("year", strconv.Itoa(opts.year))
	params.Set("page", strconv.Itoa(opts.page))

	return params.Encode()
}
