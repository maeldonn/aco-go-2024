package day22

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var input = []string{
	"1",
	"10",
	"100",
	"2024",
}

func TestPartOne(t *testing.T) {
	got := PartOne(input)
	assert.Equal(t, 37327623, got)
}

func TestPartTwo(t *testing.T) {
	got := PartTwo(input)
	assert.Equal(t, 0, got)
}
