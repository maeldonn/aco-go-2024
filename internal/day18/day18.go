package day18

import (
	"container/list"
	"image"
	"strconv"
	"strings"
)

const (
	gridSize = 71
	size     = 1024
)

func PartOne(input []string) any {
	bytes := make(map[image.Point]struct{})
	for _, line := range input[:size] {
		x, _ := strconv.Atoi(strings.Split(line, ",")[0])
		y, _ := strconv.Atoi(strings.Split(line, ",")[1])
		bytes[image.Pt(x, y)] = struct{}{}
	}

	start, end := image.Pt(0, 0), image.Pt(gridSize-1, gridSize-1)

	q := list.New()
	q.PushBack(start)

	parents := make(map[image.Point]image.Point)

	for q.Len() > 0 {
		current := q.Front().Value.(image.Point)
		q.Remove(q.Front())

		if current.Eq(end) {
			var length int
			for current != start {
				length++
				current = parents[current]
			}
			return length
		}

		dirs := []image.Point{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
		for _, dir := range dirs {
			neighbor := current.Add(dir)

			if _, ok := parents[neighbor]; ok {
				continue
			}

			if _, ok := bytes[neighbor]; ok {
				continue
			}

			if neighbor.Y < 0 || neighbor.Y > gridSize-1 {
				continue
			}

			if neighbor.X < 0 || neighbor.X > gridSize-1 {
				continue
			}

			parents[neighbor] = current
			q.PushBack(neighbor)
		}
	}

	return 0
}

func PartTwo(input []string) any {
	bytes := make(map[image.Point]struct{})
	for _, line := range input[:size] {
		x, _ := strconv.Atoi(strings.Split(line, ",")[0])
		y, _ := strconv.Atoi(strings.Split(line, ",")[1])
		bytes[image.Pt(x, y)] = struct{}{}
	}

	start, end := image.Pt(0, 0), image.Pt(gridSize-1, gridSize-1)

	bfs := func() int {
		q := list.New()
		q.PushBack(start)

		parents := make(map[image.Point]image.Point)

		for q.Len() > 0 {
			current := q.Front().Value.(image.Point)
			q.Remove(q.Front())

			if current.Eq(end) {
				var length int
				for current != start {
					length++
					current = parents[current]
				}
				return length
			}

			dirs := []image.Point{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
			for _, dir := range dirs {
				neighbor := current.Add(dir)

				if _, ok := parents[neighbor]; ok {
					continue
				}

				if _, ok := bytes[neighbor]; ok {
					continue
				}

				if neighbor.Y < 0 || neighbor.Y > gridSize-1 {
					continue
				}

				if neighbor.X < 0 || neighbor.X > gridSize-1 {
					continue
				}

				parents[neighbor] = current
				q.PushBack(neighbor)
			}
		}

		return 0
	}

	for _, line := range input[size:] {
		x, _ := strconv.Atoi(strings.Split(line, ",")[0])
		y, _ := strconv.Atoi(strings.Split(line, ",")[1])
		bytes[image.Pt(x, y)] = struct{}{}

		if bfs() == 0 {
			return line
		}
	}

	return ""
}
