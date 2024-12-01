package day1

import (
	"slices"
	"strconv"
	"strings"
)

func PartOne(input []string) (int, error) {
	var (
		first  []int
		second []int
	)

	for _, line := range input {
		split := strings.Split(line, "   ")

		n1, err := strconv.Atoi(split[0])
		if err != nil {
			return 0, err
		}
		first = append(first, n1)

		n2, err := strconv.Atoi(split[1])
		if err != nil {
			return 0, err
		}
		second = append(second, n2)
	}

	slices.Sort(first)
	slices.Sort(second)

	var score int
	for i := 0; i < len(first); i++ {
		if first[i] > second[i] {
			score += first[i] - second[i]
		} else {
			score += second[i] - first[i]
		}
	}

	return score, nil
}

func PartTwo(input []string) (int, error) {
	var (
		first  []int
		second []int
	)

	for _, line := range input {
		in := strings.Split(line, "   ")

		n1, err := strconv.Atoi(in[0])
		if err != nil {
			return 0, err
		}
		first = append(first, n1)

		n2, err := strconv.Atoi(in[1])
		if err != nil {
			return 0, err
		}
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

	return score, nil
}
