package day4

import "image"

func PartOne(input []string) int {
	directions := []image.Point{{0, 1}, {1, 0}, {0, -1}, {-1, 0}, {1, 1}, {-1, 1}, {-1, -1}, {1, -1}}

	var occurence int
	for y, line := range input {
		for x, char := range line {
			if char != 'X' {
				continue
			}

			for _, dir := range directions {
				if y+dir.Y*3 >= len(input) || y+dir.Y*3 < 0 {
					continue
				}

				if x+dir.X*3 >= len(input[y]) || x+dir.X*3 < 0 {
					continue
				}

				if input[y+dir.Y][x+dir.X] == 'M' && input[y+dir.Y*2][x+dir.X*2] == 'A' && input[y+dir.Y*3][x+dir.X*3] == 'S' {
					occurence++
				}
			}
		}
	}

	return occurence
}

func PartTwo(input []string) int {
	var occurence int
	for y, line := range input {
		for x, char := range line {
			if char != 'A' {
				continue
			}

			if y+1 >= len(input) || y-1 < 0 {
				continue
			}

			if x+1 >= len(input[y]) || x-1 < 0 {
				continue
			}

			isMas := func(r1, r2 byte) bool {
				return (r1 == 'M' || r1 == 'S') && (r2 == 'M' || r2 == 'S') && r1 != r2
			}

			if isMas(input[y-1][x-1], input[y+1][x+1]) && isMas(input[y+1][x-1], input[y-1][x+1]) {
				occurence++
			}
		}
	}

	return occurence
}
