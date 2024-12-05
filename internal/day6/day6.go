package day6

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
)

var directions = [][]int{{0, -1}, {1, 0}, {0, 1}, {-1, 0}}

func PartOne(input []string) (int, error) {
	var currX, currY int
	for j, line := range input {
		for i, char := range line {
			if char == '^' {
				currX, currY = i, j
				break
			}
		}

		if currX != 0 && currY != 0 {
			break
		}
	}

	seen := make(map[string]struct{})

	var turn int
	for {
		pos := fmt.Sprintf("%d,%d", currX, currY)
		seen[pos] = struct{}{}

		nextX := currX + directions[turn][0]
		nextY := currY + directions[turn][1]

		if nextY < 0 || nextY >= len(input) || nextX < 0 || nextX >= len(input[nextY]) {
			break
		}

		if input[nextY][nextX] == '#' {
			turn = (turn + 1) % 4
			continue
		}

		currX, currY = nextX, nextY
	}

	return len(seen), nil
}

func PartTwo(input []string) (int, error) {
	var startX, startY int
	for j, line := range input {
		for i, char := range line {
			if char == '^' {
				startX, startY = i, j
				break
			}
		}

		if startX != 0 && startY != 0 {
			break
		}
	}

	var (
		currX, currY = startX, startY
		originalSeen = make(map[string]struct{})
		turn         = 0
	)

	for {
		pos := fmt.Sprintf("%d,%d", currX, currY)
		originalSeen[pos] = struct{}{}

		nextX := currX + directions[turn][0]
		nextY := currY + directions[turn][1]

		if nextY < 0 || nextY >= len(input) || nextX < 0 || nextX >= len(input[nextY]) {
			break
		}

		if input[nextY][nextX] == '#' {
			turn = (turn + 1) % 4
			continue
		}

		currX, currY = nextX, nextY
	}

	willLoop := func(grid []string) bool {
		var (
			currX, currY = startX, startY
			seen         = make(map[string]struct{})
			turn         = 0
		)

		for {
			pos := fmt.Sprintf("%d,%d,%d", currX, currY, turn)
			if _, ok := seen[pos]; ok {
				return true
			}
			seen[pos] = struct{}{}

			nextX := currX + directions[turn][0]
			nextY := currY + directions[turn][1]

			if nextY < 0 || nextY >= len(grid) || nextX < 0 || nextX >= len(grid[nextY]) {
				return false
			}

			if grid[nextY][nextX] == '#' {
				turn = (turn + 1) % 4
				continue
			}

			currX, currY = nextX, nextY
		}
	}

	var total int

	wg := sync.WaitGroup{}
	for pos := range originalSeen {
		split := strings.Split(pos, ",")
		x, _ := strconv.Atoi(split[0])
		y, _ := strconv.Atoi(split[1])

		if x == startX && y == startY {
			continue
		}

		wg.Add(1)
		go func() {
			grid := make([]string, len(input))
			copy(grid, input)
			grid[y] = grid[y][:x] + "#" + grid[y][x+1:]

			if willLoop(grid) {
				total++
			}

			wg.Done()
		}()
	}

	wg.Wait()
	return total, nil
}
