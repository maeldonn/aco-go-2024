package day11

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var input = []string{
	"125 17",
}

func TestPartOne(t *testing.T) {
	got := PartOne(input)
	assert.Equal(t, 55312, got)
}

func TestPartTwo(t *testing.T) {
	got := PartTwo(input)
	assert.Equal(t, 65601038650482, got)
}
