package day1

import (
	"math"
	"slices"
	"strconv"
	"strings"
)

func PartOne(input []string) int {
	first := make([]int, 0, len(input))
	second := make([]int, 0, len(input))

	for _, line := range input {
		split := strings.Split(line, "   ")

		n1, _ := strconv.Atoi(split[0])
		first = append(first, n1)

		n2, _ := strconv.Atoi(split[1])
		second = append(second, n2)
	}

	slices.Sort(first)
	slices.Sort(second)

	var score int
	for i := 0; i < len(first); i++ {
		diff := float64(first[i] - second[i])
		score += int(math.Abs(diff))
	}

	return score
}

func PartTwo(input []string) int {
	first := make([]int, 0, len(input))
	second := make([]int, 0, len(input))

	for _, line := range input {
		in := strings.Split(line, "   ")

		n1, _ := strconv.Atoi(in[0])
		first = append(first, n1)

		n2, _ := strconv.Atoi(in[1])
		second = append(second, n2)
	}

	slices.Sort(first)

	freq := make(map[int]int)
	for _, v := range second {
		freq[v]++
	}

	var score int
	for i := 0; i < len(first); i++ {
		score += first[i] * freq[first[i]]
	}

	return score
}
