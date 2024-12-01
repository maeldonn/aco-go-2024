package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("cmd/day1/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var input []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	part1, err := partOne(input)
	if err != nil {
		panic(err)
	}

	part2, err := partTwo(input)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Solution of part 1: %d\n", part1)
	fmt.Printf("Solution of part 2: %d\n", part2)
}

func partOne(input []string) (int, error) {
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

	return score, nil
}

func partTwo(input []string) (int, error) {
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

	return score, nil
}
