package day2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var input = []string{
	"7 6 4 2 1",
	"1 2 7 8 9",
	"9 7 6 2 1",
	"1 3 2 4 5",
	"8 6 4 4 1",
	"1 3 6 7 9",
}

func TestPartOne(t *testing.T) {
	got := PartOne(input)
	assert.Equal(t, 2, got)
}

func TestPartTwo(t *testing.T) {
	got := PartTwo(input)
	assert.Equal(t, 4, got)
}
