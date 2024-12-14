package day14

import (
	"image"
	"regexp"
	"strconv"
	"strings"
)

const (
	wide    = 101
	tall    = 103
	seconds = 100
)

var re = regexp.MustCompile(`p=(\d+),(\d+) v=(-?\d+),(-?\d+)`)

func PartOne(input []string) int {
	var a, b, c, d int
	for _, line := range input {
		matches := re.FindStringSubmatch(line)

		posX, _ := strconv.Atoi(matches[1])
		posY, _ := strconv.Atoi(matches[2])

		dirX, _ := strconv.Atoi(matches[3])
		dirY, _ := strconv.Atoi(matches[4])

		posX = (posX + dirX*seconds) % wide
		posY = (posY + dirY*seconds) % tall

		if posX < 0 {
			posX += wide
		}

		if posY < 0 {
			posY += tall
		}

		switch {
		case posX < (wide-1)/2 && posY < (tall-1)/2:
			a++
		case posX < (wide-1)/2 && posY > (tall-1)/2:
			b++
		case posX > (wide-1)/2 && posY < (tall-1)/2:
			c++
		case posX > (wide-1)/2 && posY > (tall-1)/2:
			d++
		}
	}

	return a * b * c * d
}

type Robot struct {
	pos, dir image.Point
}

func (r *Robot) Move(seconds int) {
	r.pos.X = (r.pos.X + r.dir.X*seconds) % wide
	r.pos.Y = (r.pos.Y + r.dir.Y*seconds) % tall

	if r.pos.X < 0 {
		r.pos.X += wide
	}

	if r.pos.Y < 0 {
		r.pos.Y += tall
	}
}

func PartTwo(input []string) int {

	robots := make([]Robot, 0, len(input))
	for _, line := range input {
		matches := re.FindStringSubmatch(line)

		posX, _ := strconv.Atoi(matches[1])
		posY, _ := strconv.Atoi(matches[2])

		dirX, _ := strconv.Atoi(matches[3])
		dirY, _ := strconv.Atoi(matches[4])

		robot := Robot{
			pos: image.Pt(posX, posY),
			dir: image.Pt(dirX, dirY),
		}

		robots = append(robots, robot)

	}

	var (
		seconds int
		found   bool
	)

	for !found {
		for i := range robots {
			robots[i].Move(1)
		}
		seconds++

		grid := make([][]string, tall)
		for i := range grid {
			grid[i] = make([]string, wide)
			for j := range grid[i] {
				grid[i][j] = ""
			}
		}

		for _, robot := range robots {
			grid[robot.pos.Y][robot.pos.X] = "󰄮"
		}

		for _, line := range grid {
			if strings.Contains(strings.Join(line, ""), "󰄮󰄮󰄮󰄮󰄮󰄮󰄮󰄮󰄮") {
				found = true
			}
		}
	}

	return seconds
}
