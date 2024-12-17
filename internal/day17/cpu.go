package day17

import (
	"math"
	"strconv"
	"strings"
)

type CPU struct {
	RegisterA int
	RegisterB int
	RegisterC int
	Program   []int
}

func (c *CPU) Run() string {
	var output []string

	var cursor int
	for cursor < len(c.Program) {
		opcode := c.Program[cursor]
		operand := c.Program[cursor+1]
		combo := c.combo(operand)

		switch opcode {
		case 0:
			c.RegisterA = c.division(combo)
		case 1:
			c.RegisterB ^= operand
		case 2:
			c.RegisterB = combo % 8
		case 3:
			if c.RegisterA != 0 {
				cursor = operand
				continue
			}
		case 4:
			c.RegisterB ^= c.RegisterC
		case 5:
			output = append(output, strconv.Itoa(combo%8))
		case 6:
			c.RegisterB = c.division(combo)
		case 7:
			c.RegisterC = c.division(combo)
		}

		cursor += 2
	}

	return strings.Join(output, ",")
}

func (c *CPU) division(combo int) int {
	return int(math.Trunc(float64(c.RegisterA) / math.Pow(2, float64(combo))))
}

func (c *CPU) combo(n int) int {
	switch n {
	case 0, 1, 2, 3:
		return n
	case 4:
		return c.RegisterA
	case 5:
		return c.RegisterB
	case 6:
		return c.RegisterC
	default:
		panic("invalid combo")
	}
}
