package ncs

import (
	"context"
	"fmt"
)

// Search Result
type Result struct {
	// Keep reference for Next()
	c    *Client
	q    string
	opts *searchOptions

	Size    int
	Page    int
	Songs   []Song
	HasNext bool
}

func (r *Result) Next(ctx context.Context) (Result, error) {
	if !r.HasNext {
		return Result{}, fmt.Errorf("doesn't have further results")
	}

	// Next page is available
	// increment the page offset
	r.opts.page += 1

	return r.c.search(ctx, r.q, r.opts)
}
