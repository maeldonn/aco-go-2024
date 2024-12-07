package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/maeldonn/aoc-go-2024/internal/day1"
	"github.com/maeldonn/aoc-go-2024/internal/day2"
	"github.com/maeldonn/aoc-go-2024/internal/day3"
	"github.com/maeldonn/aoc-go-2024/internal/day4"
	"github.com/maeldonn/aoc-go-2024/internal/day5"
	"github.com/maeldonn/aoc-go-2024/internal/day6"
	"github.com/maeldonn/aoc-go-2024/internal/day7"
	"github.com/maeldonn/aoc-go-2024/internal/day8"
	aocgoclient "github.com/maeldonn/aoc-go-client"
)

func main() {
	dayStr, ok := os.LookupEnv("DAY")
	if !ok {
		panic("DAY env var not set")
	}

	day, err := strconv.Atoi(dayStr)
	if err != nil {
		panic(err)
	}

	client, err := aocgoclient.NewClient()
	if err != nil {
		panic(err)
	}

	input, err := client.GetInput(2024, day)
	if err != nil {
		panic(err)
	}

	var partOne, partTwo func([]string) (int, error)
	switch day {
	case 1:
		partOne = day1.PartOne
		partTwo = day1.PartTwo
	case 2:
		partOne = day2.PartOne
		partTwo = day2.PartTwo
	case 3:
		partOne = day3.PartOne
		partTwo = day3.PartTwo
	case 4:
		partOne = day4.PartOne
		partTwo = day4.PartTwo
	case 5:
		partOne = day5.PartOne
		partTwo = day5.PartTwo
	case 6:
		partOne = day6.PartOne
		partTwo = day6.PartTwo
	case 7:
		partOne = day7.PartOne
		partTwo = day7.PartTwo
	case 8:
		partOne = day8.PartOne
		partTwo = day8.PartTwo
	}

	fmt.Printf("########## Day %d ##########\n", day)

	start1 := time.Now()
	part1, _ := partOne(input)
	fmt.Printf("Solution of part 1: %d (took %s)\n", part1, time.Since(start1))

	start2 := time.Now()
	part2, _ := partTwo(input)
	fmt.Printf("Solution of part 2: %d (took %s)\n", part2, time.Since(start2))
}
