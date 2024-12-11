package day3

import (
	"regexp"
	"strconv"
	"strings"
)

func PartOne(input []string) int {
	text := strings.Join(input, "")

	re := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	matches := re.FindAllStringSubmatch(text, -1)

	var sum int
	for _, match := range matches {
		x, _ := strconv.Atoi(match[1])
		y, _ := strconv.Atoi(match[2])
		sum += x * y
	}

	return sum
}

func PartTwo(input []string) int {
	text := strings.Join(input, "")

	re := regexp.MustCompile(`don't\(\).*?do\(\)`)
	text = re.ReplaceAllString(text, "")

	text = strings.Split(text, "don't()")[0]

	re = regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	matches := re.FindAllStringSubmatch(text, -1)

	var sum int
	for _, match := range matches {
		x, _ := strconv.Atoi(match[1])
		y, _ := strconv.Atoi(match[2])
		sum += x * y
	}

	return sum
}
