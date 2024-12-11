package day1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var input = []string{
	"3   4",
	"4   3",
	"2   5",
	"1   3",
	"3   9",
	"3   3",
}

func TestPartOne(t *testing.T) {
	got := PartOne(input)
	assert.Equal(t, 11, got)
}

func TestPartTwo(t *testing.T) {
	got := PartTwo(input)
	assert.Equal(t, 31, got)
}
