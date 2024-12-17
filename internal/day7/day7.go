package day7

import (
	"fmt"
	"strconv"
	"strings"
)

func PartOne(input []string) any {
	var totalCalibration int
	for _, line := range input {
		split := strings.Split(line, " ")

		var (
			testValue, _ = strconv.Atoi(split[0][:len(split[0])-1])
			nums         = split[1:]
		)

		var checkCombo func(slice []string, length int) bool
		checkCombo = func(combo []string, length int) bool {
			if length == len(nums)-1 {
				return eval(nums, combo) == testValue
			}

			for _, operator := range []string{"*", "+"} {
				if checkCombo(append(combo, operator), length+1) {
					return true
				}
			}

			return false
		}

		if checkCombo([]string{}, 0) {
			totalCalibration += testValue
		}
	}

	return totalCalibration
}

func PartTwo(input []string) any {
	var totalCalibration int
	for _, line := range input {
		split := strings.Split(line, " ")

		var (
			testValue, _ = strconv.Atoi(split[0][:len(split[0])-1])
			nums         = split[1:]
		)

		var checkCombo func(slice []string, length int) bool
		checkCombo = func(combo []string, length int) bool {
			if length == len(nums)-1 {
				return eval(nums, combo) == testValue
			}

			for _, operator := range []string{"*", "+", "||"} {
				if checkCombo(append(combo, operator), length+1) {
					return true
				}
			}

			return false
		}

		if checkCombo([]string{}, 0) {
			totalCalibration += testValue
		}
	}

	return totalCalibration
}

func eval(nums, combo []string) int {
	total, _ := strconv.Atoi(nums[0])

	for i := 1; i < len(nums); i++ {
		n, _ := strconv.Atoi(nums[i])
		switch combo[i-1] {
		case "*":
			total *= n
		case "+":
			total += n
		case "||":
			conc, _ := strconv.Atoi(fmt.Sprintf("%d%d", total, n))
			total = conc
		}
	}

	return total
}
