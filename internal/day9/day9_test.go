package day9

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var input = []string{
	"2333133121414131402",
}

func TestPartOne(t *testing.T) {
	got, err := PartOne(input)
	assert.NoError(t, err)
	assert.Equal(t, 1928, got)
}

func TestPartTwo(t *testing.T) {
	got, err := PartTwo(input)
	assert.NoError(t, err)
	assert.Equal(t, 2858, got)
}
