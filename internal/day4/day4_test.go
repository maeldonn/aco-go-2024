package day4

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var input = []string{
	"MMMSXXMASM",
	"MSAMXMSMSA",
	"AMXSXMAAMM",
	"MSAMASMSMX",
	"XMASAMXAMM",
	"XXAMMXXAMA",
	"SMSMSASXSS",
	"SAXAMASAAA",
	"MAMMMXMMMM",
	"MXMXAXMASX",
}

func TestPartOne(t *testing.T) {
	got := PartOne(input)
	assert.Equal(t, 18, got)
}

func TestPartTwo(t *testing.T) {
	got := PartTwo(input)
	assert.Equal(t, 9, got)
}
