package day16

import (
	"fmt"
	"image"
	"slices"
)

var directions = []image.Point{{0, -1}, {0, 1}, {-1, 0}, {1, 0}}

func PartOne(input []string) int {
	walls := make(map[image.Point]struct{})

	var start, end image.Point
	for y, line := range input {
		for x, char := range line {
			switch char {
			case 'S':
				start = image.Pt(x, y)
			case 'E':
				end = image.Pt(x, y)
			case '#':
				walls[image.Pt(x, y)] = struct{}{}
			}
		}
	}

	// Renvoie le score du plus petit chemin et si possible d'arriver a la fin
	var move func(curr, orientation image.Point, seen map[image.Point]struct{}) (int, bool)
	move = func(curr, orientation image.Point, seen map[image.Point]struct{}) (int, bool) {
		if _, ok := walls[curr]; ok {
			return 0, false
		}

		if curr.Eq(end) {
			return 0, true
		}

		scores := make([]int, 0)
		for _, dir := range directions {
			next := curr.Add(dir)

			if _, ok := seen[next]; ok {
				continue
			}

			m := make(map[image.Point]struct{})
			for k, v := range seen {
				m[k] = v
			}
			m[next] = struct{}{}

			switch {
			case orientation.Add(dir) == (image.Point{}):
				continue
			case !orientation.Eq(dir):
				if score, ok := move(next, dir, m); ok {
					scores = append(scores, score+1000+1)
				}
			default:
				if score, ok := move(next, orientation, m); ok {
					scores = append(scores, score+1)
				}
			}
		}

		if len(scores) == 0 {
			return 0, false
		}

		return slices.Min(scores), true
	}

	score, _ := move(start, image.Pt(1, 0), make(map[image.Point]struct{}))
	return score
}

func PartTwo(input []string) int {
	return 0
}

func debug(curr image.Point, walls map[image.Point]struct{}) {
	grid := make([][]rune, 17)
	for i := range grid {
		grid[i] = make([]rune, 17)
		for j := range grid[i] {
			grid[i][j] = '.'
		}
	}

	for wall := range walls {
		grid[wall.Y][wall.X] = '#'
	}

	grid[curr.Y][curr.X] = 'X'

	for _, line := range grid {
		for _, char := range line {
			fmt.Printf("%c", char)
		}
		fmt.Println()
	}

	fmt.Scanln()
}
