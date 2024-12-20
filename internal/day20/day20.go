package day20

import (
	"container/list"
	"image"
	"math"
	"slices"
)

func PartOne(input []string) any {
	walls := make(map[image.Point]struct{})

	var start, end image.Point
	for y, line := range input {
		for x, char := range line {
			switch char {
			case '#':
				walls[image.Pt(x, y)] = struct{}{}
			case 'S':
				start = image.Pt(x, y)
			case 'E':
				end = image.Pt(x, y)
			}
			if char == '#' {
				walls[image.Pt(x, y)] = struct{}{}
			}
		}
	}

	var path []image.Point

	q := list.New()
	q.PushBack(start)

	parents := make(map[image.Point]image.Point)
	for q.Len() > 0 {
		current := q.Front().Value.(image.Point)
		q.Remove(q.Front())

		if current.Eq(end) {
			for current != start {
				path = append(path, current)
				current = parents[current]
			}
			path = append(path, start)
		}

		dirs := []image.Point{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
		for _, dir := range dirs {
			neighbor := current.Add(dir)

			if _, ok := parents[neighbor]; ok {
				continue
			}

			if _, ok := walls[neighbor]; ok {
				continue
			}

			parents[neighbor] = current
			q.PushBack(neighbor)
		}
	}

	slices.Reverse(path)

	manhattan := func(a, b image.Point) float64 {
		return math.Abs(float64(b.X)-float64(a.X)) + math.Abs(float64(b.Y)-float64(a.Y))
	}

	var (
		cheatDuration = 2.0
		save          = 2
	)

	var total int
	for i, p1 := range path[:len(path)-2-save] {
		for _, p2 := range path[i+1+save:] {
			if manhattan(p1, p2) == cheatDuration {
				total++
			}
		}
	}

	return total
}

func PartTwo(input []string) any {
	walls := make(map[image.Point]struct{})

	var start, end image.Point
	for y, line := range input {
		for x, char := range line {
			switch char {
			case '#':
				walls[image.Pt(x, y)] = struct{}{}
			case 'S':
				start = image.Pt(x, y)
			case 'E':
				end = image.Pt(x, y)
			}
			if char == '#' {
				walls[image.Pt(x, y)] = struct{}{}
			}
		}
	}

	var path []image.Point

	q := list.New()
	q.PushBack(start)

	parents := make(map[image.Point]image.Point)
	for q.Len() > 0 {
		current := q.Front().Value.(image.Point)
		q.Remove(q.Front())

		if current.Eq(end) {
			for current != start {
				path = append(path, current)
				current = parents[current]
			}
			path = append(path, start)
		}

		dirs := []image.Point{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
		for _, dir := range dirs {
			neighbor := current.Add(dir)

			if _, ok := parents[neighbor]; ok {
				continue
			}

			if _, ok := walls[neighbor]; ok {
				continue
			}

			parents[neighbor] = current
			q.PushBack(neighbor)
		}
	}

	slices.Reverse(path)

	var (
		cheatDuration = 20.0
		save          = 50
	)

	manhattan := func(a, b image.Point) float64 {
		return math.Abs(float64(b.X)-float64(a.X)) + math.Abs(float64(b.Y)-float64(a.Y))
	}

	var total int
	for i, p1 := range path[:len(path)-2-save] {
		for j, p2 := range path[i+1+save:] {
			m := manhattan(p1, p2)
			savedTime := save + 1 + j - int(m)

			if m <= cheatDuration && savedTime >= save {
				total++
			}
		}
	}

	return total
}
