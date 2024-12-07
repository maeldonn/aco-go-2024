package day8

import (
	"image"
)

func PartOne(input []string) (int, error) {
	antennas := make(map[rune][]image.Point)
	for y, line := range input {
		for x, char := range line {
			if char == '.' {
				continue
			}
			antennas[char] = append(antennas[char], image.Pt(x, y))
		}
	}

	location := make(map[image.Point]struct{})
	for _, points := range antennas {
		for _, p1 := range points {
			for _, p2 := range points {
				if p1.Eq(p2) {
					continue
				}

				antinodes := []image.Point{
					{X: 2*p2.X - p1.X, Y: 2*p2.Y - p1.Y},
					{X: 2*p1.X - p2.X, Y: 2*p1.Y - p2.Y},
				}

				for _, p := range antinodes {
					if p.Y >= 0 && p.Y < len(input) && p.X >= 0 && p.X < len(input[p.Y]) {
						location[p] = struct{}{}
					}
				}

			}
		}
	}

	return len(location), nil
}

func PartTwo(input []string) (int, error) {
	antennas := make(map[rune][]image.Point)
	for y, line := range input {
		for x, char := range line {
			if char == '.' {
				continue
			}
			antennas[char] = append(antennas[char], image.Pt(x, y))
		}
	}

	location := make(map[image.Point]struct{})
	for _, points := range antennas {
		if len(points) == 1 {
			continue
		}

		for _, p1 := range points {
			location[p1] = struct{}{}
			for _, p2 := range points {
				if p1.Eq(p2) {
					continue
				}

				computeAntinodes := func(curr, prev image.Point) {
					for curr.Y >= 0 && curr.Y < len(input) && curr.X >= 0 && curr.X < len(input[curr.Y]) {
						location[curr] = struct{}{}
						curr, prev = image.Pt(2*curr.X-prev.X, 2*curr.Y-prev.Y), curr
					}
				}

				computeAntinodes(p1, p2)
				computeAntinodes(p2, p1)
			}
		}
	}

	return len(location), nil
}
