package day16

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var input = []string{
	"###############",
	"#.......#....E#",
	"#.#.###.#.###.#",
	"#.....#.#...#.#",
	"#.###.#####.#.#",
	"#.#.#.......#.#",
	"#.#.#####.###.#",
	"#...........#.#",
	"###.#.#####.#.#",
	"#...#.....#.#.#",
	"#.#.#.###.#.#.#",
	"#.....#...#.#.#",
	"#.###.#.#.#.#.#",
	"#S..#.....#...#",
	"###############",
}

func TestPartOne(t *testing.T) {
	got := PartOne(input)
	assert.Equal(t, 7036, got)
}

func TestPartTwo(t *testing.T) {
	got := PartTwo(input)
	assert.Equal(t, 0, got)
}
