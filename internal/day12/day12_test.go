package day12

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var input = []string{
	"RRRRIICCFF",
	"RRRRIICCCF",
	"VVRRRCCFFF",
	"VVRCCCJFFF",
	"VVVVCJJCFE",
	"VVIVCCJJEE",
	"VVIIICJJEE",
	"MIIIIIJJEE",
	"MIIISIJEEE",
	"MMMISSJEEE",
}

func TestPartOne(t *testing.T) {
	got := PartOne(input)
	assert.Equal(t, 1930, got)
}

func TestPartTwo(t *testing.T) {
	got := PartTwo(input)
	assert.Equal(t, 1206, got)
}
