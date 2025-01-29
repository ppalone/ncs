package ncs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewClient(t *testing.T) {
	c := NewClient(nil)
	assert.NotNil(t, c)
}
