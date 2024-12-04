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
	got, err := PartOne(input)
	assert.NoError(t, err)
	assert.Equal(t, 11, got)
}

func TestPartTwo(t *testing.T) {
	got, err := PartTwo(input)
	assert.NoError(t, err)
	assert.Equal(t, 31, got)
}
