package day12

import "image"

var directions = []image.Point{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

func PartOne(input []string) int {
	visited := make(map[image.Point]struct{})

	var total int
	for y, line := range input {
		for x := range line {
			curr := image.Pt(x, y)
			if _, ok := visited[curr]; ok {
				continue
			}

			region := make(map[image.Point]struct{})

			var walk func(curr, prev image.Point)
			walk = func(curr, prev image.Point) {
				if curr.X < 0 || curr.Y < 0 || curr.X >= len(input) || curr.Y >= len(input[curr.X]) {
					return
				}

				if input[curr.Y][curr.X] != input[prev.Y][prev.X] {
					return
				}

				if _, ok := region[curr]; ok {
					return
				}

				visited[curr] = struct{}{}
				region[curr] = struct{}{}

				for _, dir := range directions {
					walk(curr.Add(dir), curr)
				}
			}

			visited[curr], region[curr] = struct{}{}, struct{}{}
			for _, dir := range directions {
				walk(curr.Add(dir), curr)
			}

			var perimeter int
			for plant := range region {
				for _, dir := range directions {
					if _, ok := region[plant.Add(dir)]; !ok {
						perimeter++
					}
				}
			}

			total += perimeter * len(region)
		}
	}

	return total
}

func PartTwo(input []string) int {
	visited := make(map[image.Point]struct{})

	var total int
	for y, line := range input {
		for x := range line {
			curr := image.Pt(x, y)
			if _, ok := visited[curr]; ok {
				continue
			}

			region := make(map[image.Point]struct{})

			var walk func(curr, prev image.Point)
			walk = func(curr, prev image.Point) {
				if curr.X < 0 || curr.Y < 0 || curr.X >= len(input) || curr.Y >= len(input[curr.X]) {
					return
				}

				if input[curr.Y][curr.X] != input[prev.Y][prev.X] {
					return
				}

				if _, ok := region[curr]; ok {
					return
				}

				visited[curr] = struct{}{}
				region[curr] = struct{}{}

				for _, dir := range directions {
					walk(curr.Add(dir), curr)
				}
			}

			visited[curr], region[curr] = struct{}{}, struct{}{}
			for _, dir := range directions {
				walk(curr.Add(dir), curr)
			}

			var corners int
			for plant := range region {
				isCorner := func(dir1, dir2 image.Point) {
					_, p1Exists := region[plant.Add(dir1)]
					_, p2Exists := region[plant.Add(dir2)]

					_, ok := region[plant.Add(dir1.Add(dir2))]
					if (!p1Exists && !p2Exists) || (p1Exists && p2Exists && !ok) {
						corners++
					}
				}

				right, left, down, up := directions[0], directions[1], directions[2], directions[3]

				isCorner(right, up)
				isCorner(right, down)
				isCorner(left, up)
				isCorner(left, down)
			}

			total += corners * len(region)
		}
	}

	return total
}
