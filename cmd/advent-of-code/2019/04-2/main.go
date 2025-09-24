package main

import (
	"fmt"

	. "github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	low, high := InputToRange()

	var count int
	for pw := low; pw <= high; pw++ {
		if IsPossiblePassword(Digits(pw)) {
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

func InputToRange() (int, int) {
	var low, high int
	in.Scanf("%d-%d", &low, &high)
	return low, high
}
