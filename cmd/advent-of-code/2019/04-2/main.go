package main

import (
	"fmt"
	"strings"

	"github.com/bbeck/advent-of-code/lib"
)

func main() {
	low, high := InputToRange(2019, 4)

	var count int
	for pw := low; pw <= high; pw++ {
		if IsPossiblePassword(lib.Digits(pw)) {
			count++
		}
	}
	fmt.Println(count)
}

func IsPossiblePassword(digits []int) bool {
	d := func(index int) int {
		if index < 0 || index >= len(digits) {
			return -1
		}
		return digits[index]
	}

	hasDouble, isNonDecreasing := false, true
	for i := 1; i < len(digits); i++ {
		hasDouble = hasDouble || (d(i-2) != d(i-1) && d(i-1) == d(i) && d(i) != d(i+1))
		isNonDecreasing = isNonDecreasing && (digits[i-1] <= digits[i])
	}

	return len(digits) == 6 && hasDouble && isNonDecreasing
}

func InputToRange(year, day int) (int, int) {
	parts := strings.Split(lib.InputToString(), "-")
	return lib.ParseInt(parts[0]), lib.ParseInt(parts[1])
}
