package day4

func PartOne(input []string) (int, error) {
	var occurence int
	for j, line := range input {
		for i, char := range line {
			if char != 'X' {
				continue
			}

			for x := -1; x <= 1; x++ {
				for y := -1; y <= 1; y++ {
					if x == 0 && y == 0 {
						continue
					}

					if j+y*3 >= len(input) || j+y*3 < 0 {
						continue
					}

					if i+x*3 >= len(input[j]) || i+x*3 < 0 {
						continue
					}

					if input[j+y][i+x] == 'M' && input[j+y*2][i+x*2] == 'A' && input[j+y*3][i+x*3] == 'S' {
						occurence++
					}
				}
			}
		}
	}

	return occurence, nil
}

func PartTwo(input []string) (int, error) {
	var occurence int
	for j, line := range input {
		for i, char := range line {
			if char != 'A' {
				continue
			}

			if j+1 >= len(input) || j-1 < 0 {
				continue
			}

			if i+1 >= len(input[j]) || i-1 < 0 {
				continue
			}

			isMas := func(r1, r2 byte) bool {
				return (r1 == 'M' || r1 == 'S') && (r2 == 'M' || r2 == 'S') && r1 != r2
			}

			if isMas(input[j-1][i-1], input[j+1][i+1]) && isMas(input[j+1][i-1], input[j-1][i+1]) {
				occurence++
			}
		}
	}

	return occurence, nil
}
