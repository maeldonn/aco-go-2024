package template

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var input = []string{
	"",
}

func TestPartOne(t *testing.T) {
	got, err := PartOne(input)
	assert.NoError(t, err)
	assert.Equal(t, 0, got)
}

func TestPartTwo(t *testing.T) {
	got, err := PartTwo(input)
	assert.NoError(t, err)
	assert.Equal(t, 0, got)
}
