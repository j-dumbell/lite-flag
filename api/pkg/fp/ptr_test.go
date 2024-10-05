package fp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToPtr(t *testing.T) {
	input := "foobar"
	assert.Equal(t, input, *ToPtr(input))
}

func TestFromPtr(t *testing.T) {
	input := 12
	assert.Equal(t, input, FromPtr(&input))
}
