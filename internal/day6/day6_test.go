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
	got := PartOne(input)
	assert.Equal(t, 41, got)
}

func TestPartTwo(t *testing.T) {
	got := PartTwo(input)
	assert.Equal(t, 6, got)
}
