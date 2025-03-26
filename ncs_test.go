package ncs_test

import (
	"context"
	"testing"

	"github.com/ppalone/ncs"
	"github.com/stretchr/testify/assert"
)

func TestNewClient(t *testing.T) {
	c := ncs.NewClient(nil)
	assert.NotNil(t, c)
}

// TODO: add better tests
func TestSearch(t *testing.T) {
	t.Run("with search options", func(t *testing.T) {
		c := ncs.NewClient(nil)

		opts := make([]ncs.SearchOption, 0)
		opts = append(opts, ncs.WithGenre(ncs.Dubstep))

		res, err := c.Search(context.Background(), "", opts...)
		assert.NoError(t, err)
		assert.NotNil(t, res)
		assert.True(t, res.HasNext)
		t.Log(len(res.Songs))

		res1, err := res.Next(context.Background())
		assert.NoError(t, err)
		assert.NotNil(t, res1)
		t.Log(len(res1.Songs))
	})

	t.Run("no next result", func(t *testing.T) {
		c := ncs.NewClient(nil)

		res, err := c.Search(context.Background(), "Alan Walker")
		assert.NoError(t, err)
		assert.NotEmpty(t, res.Songs)
		assert.False(t, res.HasNext)

		res1, err := res.Next(context.Background())
		assert.Error(t, err)
		assert.Empty(t, res1.Songs)
	})

	t.Run("no results", func(t *testing.T) {
		c := ncs.NewClient(nil)

		res, err := c.Search(
			context.Background(),
			"Alan Walker",
			ncs.WithGenre(ncs.DrumBass),
		)
		assert.NoError(t, err)
		assert.Empty(t, res.Songs)
		assert.False(t, res.HasNext)
	})

	t.Run("with page search option", func(t *testing.T) {
		c := ncs.NewClient(nil)

		res, err := c.Search(context.Background(), "")
		assert.NoError(t, err)
		assert.NotEmpty(t, res.Songs)
		assert.True(t, res.HasNext)

		nextRes, err := res.Next(context.Background())
		assert.NoError(t, err)
		assert.NotEmpty(t, nextRes.Songs)
		assert.True(t, nextRes.HasNext)

		opts := []ncs.SearchOption{
			ncs.WithPage(res.Page + 1),
		}
		nextResWithPage, err := c.Search(context.Background(), "", opts...)
		assert.NoError(t, err)
		assert.NotEmpty(t, nextResWithPage.Songs)
		assert.True(t, nextResWithPage.HasNext)

		// response from next & page both must match
		assert.Equal(t, nextRes.Songs, nextResWithPage.Songs)
	})
}

func TestReleases(t *testing.T) {
	c := ncs.NewClient(nil)
	res, err := c.Releases(context.TODO())
	assert.NoError(t, err)
	assert.NotEmpty(t, res.Songs)
}

func TestGetSongById(t *testing.T) {
	t.Run("with invalid id", func(t *testing.T) {
		c := ncs.NewClient(nil)

		_, err := c.GetSongById(context.Background(), "DoesNotExist")
		assert.ErrorContains(t, err, "invalid id")
	})

	t.Run("with valid id", func(t *testing.T) {
		c := ncs.NewClient(nil)

		song, err := c.GetSongById(context.Background(), "toburoots")
		assert.NoError(t, err)
		assert.Equal(t, "Roots", song.Title)
		assert.Len(t, song.Artists, 1)
		assert.Equal(t, "Tobu", song.Artists[0].Name)
		assert.NotEmpty(t, song.Downloads)
	})
}
