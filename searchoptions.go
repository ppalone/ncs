package ncs

// Search Options
type searchOptions struct {
	genre   Genre
	mood    Mood
	version Version
	page    int
}

// SearchOption
type SearchOption func(opts *searchOptions)

// Default search options
func defaultSearchOptions() *searchOptions {
	return &searchOptions{}
}

// SearchOption with Genre
func WithGenre(genre Genre) SearchOption {
	return func(opts *searchOptions) {
		opts.genre = genre
	}
}

// SearchOption with Mood
func WithMood(mood Mood) SearchOption {
	return func(opts *searchOptions) {
		opts.mood = mood
	}
}

// SearchOption with Version
func WithVersion(version Version) SearchOption {
	return func(opts *searchOptions) {
		opts.version = version
	}
}
