package day8

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var input = []string{
	"............",
	"........0...",
	".....0......",
	".......0....",
	"....0.......",
	"......A.....",
	"............",
	"............",
	"........A...",
	".........A..",
	"............",
	"............",
}

func TestPartOne(t *testing.T) {
	got := PartOne(input)
	assert.Equal(t, 14, got)
}

func TestPartTwo(t *testing.T) {
	got := PartTwo(input)
	assert.Equal(t, 34, got)
}
