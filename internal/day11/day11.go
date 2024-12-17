package day11

import (
	"fmt"
	"strconv"
	"strings"
)

func PartOne(input []string) any {
	var total int

	var blink func(stone string, depth int)
	blink = func(stone string, depth int) {
		if depth == 25 {
			total++
			return
		}

		switch {
		case stone == "0":
			blink("1", depth+1)
		case len(stone)%2 == 0:
			n1, _ := strconv.Atoi(stone[:len(stone)/2])
			blink(strconv.Itoa(n1), depth+1)

			n2, _ := strconv.Atoi(stone[len(stone)/2:])
			blink(strconv.Itoa(n2), depth+1)
		default:
			n, _ := strconv.Atoi(stone)
			blink(strconv.Itoa(n*2024), depth+1)
		}
	}

	for _, stone := range strings.Split(input[0], " ") {
		blink(stone, 0)
	}

	return total
}

func PartTwo(input []string) any {
	cache := make(map[string]int)

	var blink func(stone string, depth int) int
	blink = func(stone string, depth int) int {
		key := fmt.Sprintf("%s;%d", stone, depth)
		if val, ok := cache[key]; ok {
			return val
		}

		if depth == 75 {
			return 1
		}

		switch {
		case stone == "0":
			cache[key] = blink("1", depth+1)
		case len(stone)%2 == 0:
			n1, _ := strconv.Atoi(stone[:len(stone)/2])
			total := blink(strconv.Itoa(n1), depth+1)

			n2, _ := strconv.Atoi(stone[len(stone)/2:])
			total += blink(strconv.Itoa(n2), depth+1)

			cache[key] = total
		default:
			n, _ := strconv.Atoi(stone)
			cache[key] = blink(strconv.Itoa(n*2024), depth+1)
		}

		return cache[key]
	}

	var total int
	for _, stone := range strings.Split(input[0], " ") {
		total += blink(stone, 0)
	}

	return total
}
