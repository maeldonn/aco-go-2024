package day9

import (
	"strconv"
)

func PartOne(input []string) (int, error) {
	var size int
	for _, c := range input[0] {
		n, _ := strconv.Atoi(string(c))
		size += n
	}

	memory := make([]int, size)

	var pos, id int
	for i, c := range input[0] {
		n, _ := strconv.Atoi(string(c))
		for j := range n {
			if i%2 == 0 {
				memory[pos+j] = id
			} else {
				memory[pos+j] = -1
			}
		}

		if i%2 == 0 {
			id++
		}

		pos += n
	}

	i, j := 0, len(memory)-1
	for i < j {
		if memory[j] == -1 {
			j--
			continue
		}

		if memory[i] != -1 {
			i++
			continue
		}

		memory[i], memory[j] = memory[j], memory[i]
		i++
		j--
	}

	var checksum int
	for i, n := range memory {
		if n == -1 {
			break
		}
		checksum += i * n
	}

	return checksum, nil
}

func PartTwo(input []string) (int, error) {
	var size int
	for _, c := range input[0] {
		n, _ := strconv.Atoi(string(c))
		size += n
	}

	memory := make([]int, size)

	var pos, id int
	for i, c := range input[0] {
		n, _ := strconv.Atoi(string(c))

		for j := range n {
			if i%2 == 0 {
				memory[pos+j] = id
			} else {
				memory[pos+j] = -1
			}
		}

		if i%2 == 0 {
			id++
		}

		pos += n
	}

	i, j := 0, len(memory)-1
	for range id + 1 {
		var slotSize, freeSpace int
		for i < j {
			if memory[j] == -1 {
				j--
				continue
			}

			if memory[i] != -1 {
				i++
				continue
			}

			slotSize = 1
			for n := 1; j-n >= 0 && memory[j-n] == memory[j]; n++ {
				slotSize++
			}

			freeSpace = 1
			for n := 1; i+n < len(memory) && memory[i+n] == memory[i]; n++ {
				freeSpace++
			}

			if freeSpace < slotSize {
				i++
				continue
			}

			for n := 0; n < slotSize; n++ {
				memory[i+n], memory[j-n] = memory[j-n], memory[i+n]
			}

			break
		}
		i, j = 0, j-slotSize
	}

	var checksum int
	for i, n := range memory {
		if n != -1 {
			checksum += i * n
		}
	}

	return checksum, nil
}
