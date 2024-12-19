package day19

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var input = []string{
	"r, wr, b, g, bwu, rb, gb, br",
	"",
	"brwrr",
	"bggr",
	"gbbr",
	"rrbgbr",
	"ubwu",
	"bwurrg",
	"brgr",
	"bbrgwb",
}

func TestPartOne(t *testing.T) {
	got := PartOne(input)
	assert.Equal(t, 6, got)
}

func TestPartTwo(t *testing.T) {
	got := PartTwo(input)
	assert.Equal(t, 16, got)
}
