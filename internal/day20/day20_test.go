package day20

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var input = []string{
	"###############",
	"#...#...#.....#",
	"#.#.#.#.#.###.#",
	"#S#...#.#.#...#",
	"#######.#.#.###",
	"#######.#.#...#",
	"#######.#.###.#",
	"###..E#...#...#",
	"###.#######.###",
	"#...###...#...#",
	"#.#####.#.###.#",
	"#.#...#.#.#...#",
	"#.#.#.#.#.#.###",
	"#...#...#...###",
	"###############",
}

func TestPartOne(t *testing.T) {
	got := PartOne(input)
	assert.Equal(t, 44, got)
}

func TestPartTwo(t *testing.T) {
	got := PartTwo(input)
	assert.Equal(t, 285, got)
}
