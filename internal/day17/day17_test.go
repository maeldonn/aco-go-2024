package day17

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var input = []string{
	"Register A: 729",
	"Register B: 0",
	"Register C: 0",
	"",
	"Program: 0,1,5,4,3,0",
}

func TestPartOne(t *testing.T) {
	got := PartOne(input)
	assert.Equal(t, "4,6,3,5,6,3,5,2,1,0", got)
}

func TestPartTwo(t *testing.T) {
	got := PartTwo(input)
	assert.Equal(t, "", got)
}
