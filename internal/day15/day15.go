package day15

import (
	"image"
	"slices"
	"strings"
)

func PartOne(input []string) any {
	var (
		robot image.Point
		boxes = make(map[image.Point]struct{})
		walls = make(map[image.Point]struct{})
		moves = make([]image.Point, 0)
	)

	for y, line := range input {
		for x, char := range line {
			if len(line) != 0 && line[0] != '#' {
				switch char {
				case 'v':
					moves = append(moves, image.Pt(0, 1))
				case '^':
					moves = append(moves, image.Pt(0, -1))
				case '<':
					moves = append(moves, image.Pt(-1, 0))
				case '>':
					moves = append(moves, image.Pt(1, 0))
				}

				continue
			}

			curr := image.Pt(x, y)
			switch char {
			case '@':
				robot = curr
			case '#':
				walls[curr] = struct{}{}
			case 'O':
				boxes[curr] = struct{}{}
			}
		}
	}

	for _, dir := range moves {
		next := robot.Add(dir)

		if _, ok := walls[next]; ok {
			continue
		}

		if _, ok := boxes[next]; ok {
			n := next.Add(dir)
			for {
				if _, ok := boxes[n]; !ok {
					break
				}
				n = n.Add(dir)
			}

			if _, ok := walls[n]; ok {
				continue
			}

			delete(boxes, next)
			boxes[n] = struct{}{}
		}

		robot = next
	}

	var sum int
	for box := range boxes {
		sum += box.X + box.Y*100
	}

	return sum
}

func PartTwo(input []string) any {
	var (
		robot image.Point
		moves = make([]image.Point, 0)
		grid  = make([][]string, len(input[0]))
	)

	for y, line := range input {
		var wideLine string
		for x, char := range line {
			if len(line) != 0 && line[0] != '#' {
				switch char {
				case 'v':
					moves = append(moves, image.Pt(0, 1))
				case '^':
					moves = append(moves, image.Pt(0, -1))
				case '<':
					moves = append(moves, image.Pt(-1, 0))
				case '>':
					moves = append(moves, image.Pt(1, 0))
				}

				continue
			}

			switch char {
			case '@':
				robot = image.Pt(x*2, y)
				wideLine += "@."
			case '#':
				wideLine += "##"
			case 'O':
				wideLine += "[]"
			default:
				wideLine += ".."
			}
		}

		if wideLine != "" {
			grid[y] = strings.Split(wideLine, "")
		}
	}

	for _, dir := range moves {
		next := robot.Add(dir)

		if grid[next.Y][next.X] == "#" {
			continue
		}

		if grid[next.Y][next.X] == "[" || grid[next.Y][next.X] == "]" {
			if dir.X != 0 {
				var moveIfPossible func(box image.Point) bool
				moveIfPossible = func(box image.Point) bool {
					next := box.Add(dir.Mul(2))

					switch grid[next.Y][next.X] {
					case "#":
						return false
					case "[", "]":
						if !moveIfPossible(next) {
							return false
						}
					}

					grid[next.Y][next.X] = grid[box.Add(dir).Y][box.Add(dir).X]
					grid[box.Add(dir).Y][box.Add(dir).X] = grid[box.Y][box.X]
					grid[box.Y][box.X] = "."
					return true
				}

				if !moveIfPossible(next) {
					continue
				}
			} else if dir.Y != 0 {
				virtualGrid := make([][]string, len(grid))
				for i := range grid {
					virtualGrid[i] = slices.Clone(grid[i])
				}

				var moveIfPossible func(box image.Point) bool
				moveIfPossible = func(box image.Point) bool {
					currOpen, currClose := box.Add(image.Pt(-1, 0)), box
					if virtualGrid[box.Y][box.X] == "[" {
						currOpen, currClose = box, box.Add(image.Pt(1, 0))
					}
					nextOpen, nextClose := currOpen.Add(dir), currClose.Add(dir)

					if virtualGrid[nextOpen.Y][nextOpen.X] == "#" || virtualGrid[nextClose.Y][nextClose.X] == "#" {
						return false
					}

					if virtualGrid[nextOpen.Y][nextOpen.X] == "." && virtualGrid[nextClose.Y][nextClose.X] == "." {
						virtualGrid[nextOpen.Y][nextOpen.X] = virtualGrid[currOpen.Y][currOpen.X]
						virtualGrid[nextClose.Y][nextClose.X] = virtualGrid[currClose.Y][currClose.X]
						virtualGrid[currOpen.Y][currOpen.X] = "."
						virtualGrid[currClose.Y][currClose.X] = "."
						return true
					}

					canMoveOpen := true
					if virtualGrid[nextOpen.Y][nextOpen.X] == "[" || virtualGrid[nextOpen.Y][nextOpen.X] == "]" {
						canMoveOpen = moveIfPossible(nextOpen)
					}

					canMoveClose := true
					if virtualGrid[nextClose.Y][nextClose.X] == "[" {
						canMoveClose = moveIfPossible(nextClose)
					}

					canMove := canMoveOpen && canMoveClose
					if canMove {
						virtualGrid[nextOpen.Y][nextOpen.X] = virtualGrid[currOpen.Y][currOpen.X]
						virtualGrid[nextClose.Y][nextClose.X] = virtualGrid[currClose.Y][currClose.X]
						virtualGrid[currOpen.Y][currOpen.X] = "."
						virtualGrid[currClose.Y][currClose.X] = "."
					}

					return canMove
				}

				if !moveIfPossible(next) {
					continue
				}
				grid = virtualGrid
			}
		}

		grid[next.Y][next.X] = "@"
		grid[robot.Y][robot.X] = "."
		robot = next
	}

	var sum int
	for y, line := range grid {
		for x, char := range line {
			if char == "[" {
				sum += x + y*100
			}
		}
	}

	return sum
}
