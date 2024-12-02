package day2

import (
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

		if isIncreasing || isDecreasing {
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

		var (
			isIncreasing = true
			isDecreasing = true
			dampenerUsed = false
		)

		for i := 1; i < len(numbers); i++ {
			diff := numbers[i] - numbers[i-1]
			if diff <= 0 || diff > 3 {
				if dampenerUsed {
					isIncreasing = false
				} else {
					dampenerUsed = true
				}
			}
			if diff >= 0 || -diff > 3 {
				if dampenerUsed {
					isDecreasing = false
				} else {
					dampenerUsed = true
				}
			}
		}

		if isIncreasing || isDecreasing {
			safeReports++
		}
	}

	return safeReports, nil
}
