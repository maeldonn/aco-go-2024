package day7

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var input = []string{
	"190: 10 19",
	"3267: 81 40 27",
	"83: 17 5",
	"156: 15 6",
	"7290: 6 8 6 15",
	"161011: 16 10 13",
	"192: 17 8 14",
	"21037: 9 7 18 13",
	"292: 11 6 16 20",
}

func TestPartOne(t *testing.T) {
	got := PartOne(input)
	assert.Equal(t, 3749, got)
}

func TestPartTwo(t *testing.T) {
	got := PartTwo(input)
	assert.Equal(t, 11387, got)
}
