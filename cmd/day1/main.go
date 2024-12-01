package main

import (
	"bufio"
	"fmt"
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

func partTwo(input []string) (int, error) {
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
