package day2

import (
	"testing"
)

func IsCorrectV2(numbers []int) bool {
	diff := make(map[int]struct{})
	for i := 1; i < len(numbers); i++ {
		diff[numbers[i]-numbers[i-1]] = struct{}{}
	}

	isIncreasing := true
	for d := range diff {
		if d < 1 || d > 3 {
			isIncreasing = false
		}
	}

	isDecreasing := true
	for d := range diff {
		if -d < 1 || -d > 3 {
			isDecreasing = false
		}
	}

	return isIncreasing && isDecreasing
}

func IsCorrect(numbers []int) bool {
	var (
		isIncreasing = true
		isDecreasing = true
	)

	for i := 1; i < len(numbers); i++ {
		diff := numbers[i] - numbers[i-1]
		if diff <= 0 || diff > 3 {
			isIncreasing = false
		}
		if diff >= 0 || -diff > 3 {
			isDecreasing = false
		}
	}

	return isIncreasing || isDecreasing
}

func Test_IsCorrect(t *testing.T) {
	tests := map[string]struct {
		args []int
		want bool
	}{
		"1": {
			args: []int{7, 6, 4, 2, 1},
			want: true,
		},
		"2": {
			args: []int{1, 2, 7, 8, 9},
			want: false,
		},
		"3": {
			args: []int{9, 7, 6, 2, 1},
			want: false,
		},
		"4": {
			args: []int{1, 3, 2, 4, 5},
			want: false,
		},
		"5": {
			args: []int{8, 6, 4, 4, 1},
			want: false,
		},
		"6": {
			args: []int{1, 3, 6, 7, 9},
			want: true,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			if got := IsCorrectV2(tt.args); got != tt.want {
				t.Errorf("IsCorrect() = %v, want %v", got, tt.want)
			}
		})
	}
}
