package day19

import (
	"fmt"
	"strings"
)

func PartOne(input []string) any {
	patterns := make(map[string]bool)
	for _, color := range strings.Split(input[0], ",") {
		patterns[strings.TrimSpace(color)] = true
	}

	cache := make(map[string]bool)

	var validDesign func(towel string) bool
	validDesign = func(towel string) bool {
		if towel == "" {
			return true
		}

		if v, ok := cache[towel]; ok {
			return v
		}

		for i := 1; i <= len(towel); i++ {
			if !patterns[towel[:i]] {
				continue
			}

			if cache[towel[i:]] = validDesign(towel[i:]); cache[towel[i:]] {
				return true
			}
		}

		return false
	}

	var total int
	for _, towel := range input[2:] {
		fmt.Println(towel)
		if validDesign(towel) {
			total++
		}
	}

	return total
}

func PartTwo(input []string) any {
	patterns := make(map[string]bool)
	for _, color := range strings.Split(input[0], ",") {
		patterns[strings.TrimSpace(color)] = true
	}

	cache := make(map[string]int)

	var countValidDesign func(towel string) int
	countValidDesign = func(towel string) int {
		if towel == "" {
			return 1
		}

		if v, ok := cache[towel]; ok {
			return v
		}

		var total int
		for i := 1; i <= len(towel); i++ {
			if !patterns[towel[:i]] {
				continue
			}

			cache[towel[i:]] = countValidDesign(towel[i:])
			total += cache[towel[i:]]
		}

		return total
	}

	var total int
	for _, towel := range input[2:] {
		total += countValidDesign(towel)
	}

	return total
}
