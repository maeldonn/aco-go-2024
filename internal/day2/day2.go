package day2

import (
	"slices"
	"strconv"
	"strings"
)

func PartOne(input []string) int {
	var safeReports int
	for _, line := range input {
		numbers := strings.Split(line, " ")
		if isSafe(numbers) {
			safeReports++
		}
	}

	return safeReports
}

func PartTwo(input []string) int {
	var safeReports int
	for _, line := range input {
		numbers := strings.Split(line, " ")
		for i := 0; i < len(numbers); i++ {
			s := slices.Delete(slices.Clone(numbers), i, i+1)
			if isSafe(s) {
				safeReports++
				break
			}
		}
	}

	return safeReports
}

func isSafe(numbers []string) bool {
	var (
		isIncreasing = true
		isDecreasing = true
	)

	for i := 1; i < len(numbers); i++ {
		curr, _ := strconv.Atoi(numbers[i])
		prev, _ := strconv.Atoi(numbers[i-1])

		diff := curr - prev
		if diff <= 0 || diff > 3 {
			isIncreasing = false
		}
		if diff >= 0 || -diff > 3 {
			isDecreasing = false
		}
	}

	return isIncreasing || isDecreasing
}
