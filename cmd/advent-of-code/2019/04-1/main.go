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
	hasDouble, isNonDecreasing := false, true
	for i := 1; i < len(digits); i++ {
		hasDouble = hasDouble || (digits[i-1] == digits[i])
		isNonDecreasing = isNonDecreasing && (digits[i-1] <= digits[i])
	}

	return len(digits) == 6 && hasDouble && isNonDecreasing
}

func InputToRange() (int, int) {
	var low, high int
	in.Scanf("%d-%d", &low, &high)
	return low, high
}
