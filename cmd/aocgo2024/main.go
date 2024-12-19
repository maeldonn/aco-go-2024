package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/maeldonn/aoc-go-2024/internal/day1"
	"github.com/maeldonn/aoc-go-2024/internal/day10"
	"github.com/maeldonn/aoc-go-2024/internal/day11"
	"github.com/maeldonn/aoc-go-2024/internal/day12"
	"github.com/maeldonn/aoc-go-2024/internal/day13"
	"github.com/maeldonn/aoc-go-2024/internal/day14"
	"github.com/maeldonn/aoc-go-2024/internal/day15"
	"github.com/maeldonn/aoc-go-2024/internal/day17"
	"github.com/maeldonn/aoc-go-2024/internal/day18"
	"github.com/maeldonn/aoc-go-2024/internal/day19"
	"github.com/maeldonn/aoc-go-2024/internal/day2"
	"github.com/maeldonn/aoc-go-2024/internal/day3"
	"github.com/maeldonn/aoc-go-2024/internal/day4"
	"github.com/maeldonn/aoc-go-2024/internal/day5"
	"github.com/maeldonn/aoc-go-2024/internal/day6"
	"github.com/maeldonn/aoc-go-2024/internal/day7"
	"github.com/maeldonn/aoc-go-2024/internal/day8"
	"github.com/maeldonn/aoc-go-2024/internal/day9"
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

	var partOne, partTwo func([]string) any
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
	case 9:
		partOne = day9.PartOne
		partTwo = day9.PartTwo
	case 10:
		partOne = day10.PartOne
		partTwo = day10.PartTwo
	case 11:
		partOne = day11.PartOne
		partTwo = day11.PartTwo
	case 12:
		partOne = day12.PartOne
		partTwo = day12.PartTwo
	case 13:
		partOne = day13.PartOne
		partTwo = day13.PartTwo
	case 14:
		partOne = day14.PartOne
		partTwo = day14.PartTwo
	case 15:
		partOne = day15.PartOne
		partTwo = day15.PartTwo
	case 17:
		partOne = day17.PartOne
		partTwo = day17.PartTwo
	case 18:
		partOne = day18.PartOne
		partTwo = day18.PartTwo
	case 19:
		partOne = day19.PartOne
		partTwo = day19.PartTwo
	}

	fmt.Printf("########## Day %d ##########\n", day)

	start1 := time.Now()
	part1 := partOne(input)
	fmt.Printf("Solution of part 1: %v (took %s)\n", part1, time.Since(start1))

	start2 := time.Now()
	part2 := partTwo(input)
	fmt.Printf("Solution of part 2: %v (took %s)\n", part2, time.Since(start2))
}
