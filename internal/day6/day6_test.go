package day6

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var input = []string{
	"....#.....",
	".........#",
	"..........",
	"..#.......",
	".......#..",
	"..........",
	".#..^.....",
	"........#.",
	"#.........",
	"......#...",
}

func TestPartOne(t *testing.T) {
	got, err := PartOne(input)
	assert.NoError(t, err)
	assert.Equal(t, 41, got)
}

func TestPartTwo(t *testing.T) {
	got, err := PartTwo(input)
	assert.NoError(t, err)
	assert.Equal(t, 6, got)
}
