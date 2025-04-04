package ncs

import "time"

// Song
type Song struct {
	Id          string
	Title       string
	CoverURL    string
	MediaURL    string
	Artists     []Artist
	Versions    []string
	Genres      []string
	Moods       []string
	WebURL      string
	ReleaseDate time.Time
	Downloads   []Download
}
