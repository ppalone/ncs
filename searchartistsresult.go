package ncs

import (
	"context"
	"fmt"
)

// Search Artists Result
type SearchArtistsResult struct {
	// Keep reference for Next()
	c    *Client
	q    string
	opts *searchArtistOptions

	Size    int
	Page    int
	Artists []ArtistInfo
	HasNext bool
}

func (r *SearchArtistsResult) Next(ctx context.Context) (SearchArtistsResult, error) {
	if !r.HasNext {
		return SearchArtistsResult{}, fmt.Errorf("no further results")
	}

	r.opts.page += 1
	return r.c.searchArtists(ctx, r.q, r.opts)
}
