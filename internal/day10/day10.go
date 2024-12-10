package day10

import (
	"image"
)

func PartOne(input []string) (int, error) {
	var scores int

	directions := []image.Point{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

	for y, line := range input {
		for x, n := range line {
			if n != '0' {
				continue
			}

			visited := make(map[image.Point]struct{})

			var walk func(curr, prev image.Point)
			walk = func(curr, prev image.Point) {
				if curr.X < 0 || curr.Y < 0 || curr.X >= len(input) || curr.Y >= len(input[curr.X]) {
					return
				}

				if input[curr.Y][curr.X] != input[prev.Y][prev.X]+1 {
					return
				}

				if input[curr.Y][curr.X] == '9' {
					visited[curr] = struct{}{}
					return
				}

				for _, dir := range directions {
					walk(curr.Add(dir), curr)
				}
			}

			curr := image.Pt(x, y)
			for _, dir := range directions {
				walk(curr.Add(dir), curr)
			}

			scores += len(visited)
		}
	}

	return scores, nil
}

func PartTwo(input []string) (int, error) {
	var scores int
	directions := []image.Point{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

	for y, line := range input {
		for x, n := range line {
			if n != '0' {
				continue
			}

			var walk func(curr, prev image.Point)
			walk = func(curr, prev image.Point) {
				if curr.X < 0 || curr.Y < 0 || curr.X >= len(input) || curr.Y >= len(input[curr.X]) {
					return
				}

				if input[curr.Y][curr.X] != input[prev.Y][prev.X]+1 {
					return
				}

				if input[curr.Y][curr.X] == '9' {
					scores++
					return
				}

				for _, dir := range directions {
					walk(curr.Add(dir), curr)
				}
			}

			curr := image.Pt(x, y)
			for _, dir := range directions {
				walk(curr.Add(dir), curr)
			}
		}
	}

	return scores, nil
}
