package day5

import (
	"slices"
	"strconv"
	"strings"
)

func PartOne(input []string) (int, error) {
	var pageNumbers int

	rules := make(map[string][]string)
	for _, line := range input {
		if strings.Contains(line, "|") {
			parts := strings.Split(line, "|")
			page1, page2 := parts[0], parts[1]
			rules[page1] = append(rules[page1], page2)
			continue
		}

		if len(line) == 0 {
			continue
		}

		pages := strings.Split(line, ",")

		var badOrder bool
		for i := range pages {
			for j := range pages[0:i] {
				if slices.Contains(rules[pages[i]], pages[j]) {
					badOrder = true
					break
				}
			}
		}

		if !badOrder {
			middle, _ := strconv.Atoi(pages[len(pages)/2])
			pageNumbers += middle
		}
	}

	return pageNumbers, nil
}

func PartTwo(input []string) (int, error) {
	var pageNumbers int

	rules := make(map[string][]string)
	for _, line := range input {
		parts := strings.Split(line, "|")
		if len(parts) == 2 {
			page1, page2 := parts[0], parts[1]
			rules[page1] = append(rules[page1], page2)
			continue
		}

		if len(line) == 0 {
			continue
		}

		pages := strings.Split(line, ",")

		var badOrder bool
		for i := range pages {
			for j := range pages[0:i] {
				if slices.Contains(rules[pages[i]], pages[j]) {
					badOrder = true
					pages[i], pages[j] = pages[j], pages[i]
				}
			}
		}

		if badOrder {
			middle, _ := strconv.Atoi(pages[len(pages)/2])
			pageNumbers += middle
		}
	}

	return pageNumbers, nil
}
