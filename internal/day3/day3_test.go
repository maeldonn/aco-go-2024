package day3

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPartOne(t *testing.T) {
	input := []string{"xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"}
	got := PartOne(input)
	assert.Equal(t, 161, got)
}

func TestPartTwo(t *testing.T) {
	input := []string{"xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"}
	got := PartTwo(input)
	assert.Equal(t, 48, got)
}
