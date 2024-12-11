package day6

import (
	"fmt"
	"image"
	"sync"
)

var directions = []image.Point{{0, -1}, {1, 0}, {0, 1}, {-1, 0}}

func PartOne(input []string) int {
	var curr image.Point
	for y, line := range input {
		for x, char := range line {
			if char == '^' {
				curr = image.Pt(x, y)
				break
			}
		}

		if curr != (image.Point{}) {
			break
		}
	}

	seen := make(map[image.Point]struct{})

	var turn int
	for {
		seen[curr] = struct{}{}
		next := curr.Add(directions[turn])

		if next.Y < 0 || next.Y >= len(input) || next.X < 0 || next.X >= len(input[next.Y]) {
			break
		}

		if input[next.Y][next.X] == '#' {
			turn = (turn + 1) % 4
			continue
		}

		curr = next
	}

	return len(seen)
}

func PartTwo(input []string) int {
	var start image.Point
	for j, line := range input {
		for i, char := range line {
			if char == '^' {
				start = image.Pt(i, j)
				break
			}
		}

		if start != (image.Point{}) {
			break
		}
	}

	var (
		curr         = start
		originalSeen = make(map[image.Point]struct{})
		turn         = 0
	)

	for {
		originalSeen[curr] = struct{}{}
		next := curr.Add(directions[turn])

		if next.Y < 0 || next.Y >= len(input) || next.X < 0 || next.X >= len(input[next.Y]) {
			break
		}

		if input[next.Y][next.X] == '#' {
			turn = (turn + 1) % 4
			continue
		}

		curr = next
	}

	willLoop := func(grid []string) bool {
		var (
			curr = start
			seen = make(map[string]struct{})
			turn = 0
		)

		for {
			pos := fmt.Sprintf("%d,%d,%d", curr.X, curr.Y, turn)
			if _, ok := seen[pos]; ok {
				return true
			}

			seen[pos] = struct{}{}
			next := curr.Add(directions[turn])

			if next.Y < 0 || next.Y >= len(grid) || next.X < 0 || next.X >= len(grid[next.Y]) {
				return false
			}

			if grid[next.Y][next.X] == '#' {
				turn = (turn + 1) % 4
				continue
			}

			curr = next
		}
	}

	var total int

	wg := sync.WaitGroup{}
	for pos := range originalSeen {
		if pos.X == start.X && pos.X == start.Y {
			continue
		}

		wg.Add(1)
		go func() {
			grid := make([]string, len(input))
			copy(grid, input)
			grid[pos.Y] = grid[pos.Y][:pos.X] + "#" + grid[pos.Y][pos.X+1:]

			if willLoop(grid) {
				total++
			}

			wg.Done()
		}()
	}

	wg.Wait()
	return total
}
