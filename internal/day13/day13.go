package day13

import (
	"math"
	"regexp"
	"strconv"
)

const delta = 10_000_000_000_000

var (
	reButton = regexp.MustCompile(`Button [A|B]: X\+(\d+), Y\+(\d+)`)
	rePrize  = regexp.MustCompile(`Prize: X\=(\d+), Y\=(\d+)`)
)

func PartOne(input []string) any {
	var tokens int

	for i := 0; i < len(input); i += 4 {
		matches := reButton.FindStringSubmatch(input[i])
		a, _ := strconv.ParseFloat(matches[1], 64)
		d, _ := strconv.ParseFloat(matches[2], 64)

		matches = reButton.FindStringSubmatch(input[i+1])
		b, _ := strconv.ParseFloat(matches[1], 64)
		e, _ := strconv.ParseFloat(matches[2], 64)

		matches = rePrize.FindStringSubmatch(input[i+2])
		c, _ := strconv.ParseFloat(matches[1], 64)
		f, _ := strconv.ParseFloat(matches[2], 64)

		y := (f*a - d*c) / (a*e - d*b)
		if y != math.Trunc(y) {
			continue
		}

		x := (c - b*y) / a
		if x != math.Trunc(x) {
			continue
		}

		tokens += int(3*x + y)
	}

	return tokens
}

func PartTwo(input []string) any {
	var tokens int

	for i := 0; i < len(input); i += 4 {
		matches := reButton.FindStringSubmatch(input[i])
		a, _ := strconv.ParseFloat(matches[1], 64)
		d, _ := strconv.ParseFloat(matches[2], 64)

		matches = reButton.FindStringSubmatch(input[i+1])
		b, _ := strconv.ParseFloat(matches[1], 64)
		e, _ := strconv.ParseFloat(matches[2], 64)

		matches = rePrize.FindStringSubmatch(input[i+2])
		c, _ := strconv.ParseFloat(matches[1], 64)
		f, _ := strconv.ParseFloat(matches[2], 64)

		c, f = c+delta, f+delta

		y := (f*a - d*c) / (a*e - d*b)
		if y != math.Trunc(y) {
			continue
		}

		x := (c - b*y) / a
		if x != math.Trunc(x) {
			continue
		}

		tokens += int(3*x + y)
	}

	return tokens
}
