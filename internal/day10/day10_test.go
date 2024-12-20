package day10

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var input = []string{
	"89010123",
	"78121874",
	"87430965",
	"96549874",
	"45678903",
	"32019012",
	"01329801",
	"10456732",
}

func TestPartOne(t *testing.T) {
	got := PartOne(input)
	assert.Equal(t, 36, got)
}

func TestPartTwo(t *testing.T) {
	got := PartTwo(input)
	assert.Equal(t, 81, got)
}
