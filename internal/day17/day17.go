package day17

import (
	"strconv"
	"strings"
)

func PartOne(input []string) any {
	registerA, _ := strconv.Atoi(strings.Split(input[0], " ")[2])
	registerB, _ := strconv.Atoi(strings.Split(input[1], " ")[2])
	registerC, _ := strconv.Atoi(strings.Split(input[2], " ")[2])
	instructions := strings.Split(strings.Split(input[4], " ")[1], ",")

	program := make([]int, 0, len(instructions))
	for _, instruction := range instructions {
		n, _ := strconv.Atoi(instruction)
		program = append(program, n)
	}

	cpu := CPU{
		RegisterA: registerA,
		RegisterB: registerB,
		RegisterC: registerC,
		Program:   program,
	}

	return cpu.Run()
}

func PartTwo(input []string) any {
	return ""
}
