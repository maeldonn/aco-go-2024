package day2

import (
	"slices"
	"strconv"
	"strings"
)

func PartOne(input []string) (int, error) {
	var safeReports int
	for _, line := range input {
		split := strings.Split(line, " ")

		numbers := make([]int, 0, len(split))
		for _, s := range split {
			n, err := strconv.Atoi(s)
			if err != nil {
				return 0, err
			}
			numbers = append(numbers, n)
		}

		if isSafe(numbers) {
			safeReports++
		}
	}

	return safeReports, nil
}

func PartTwo(input []string) (int, error) {
	var safeReports int
	for _, line := range input {
		split := strings.Split(line, " ")

		numbers := make([]int, 0, len(split))
		for _, s := range split {
			n, err := strconv.Atoi(s)
			if err != nil {
				return 0, err
			}
			numbers = append(numbers, n)
		}

		for i := 0; i < len(numbers); i++ {
			var (
				clone = slices.Clone(numbers)
				s     = slices.Delete(clone, i, i+1)
			)

			if isSafe(s) {
				safeReports++
				break
			}
		}
	}

	return safeReports, nil
}

func isSafe(numbers []int) bool {
	var (
		isIncreasing = true
		isDecreasing = true
	)

	for i := 1; i < len(numbers); i++ {
		diff := numbers[i] - numbers[i-1]
		if diff <= 0 || diff > 3 {
			isIncreasing = false
		}
		if diff >= 0 || -diff > 3 {
			isDecreasing = false
		}
	}

	return isIncreasing || isDecreasing
}
