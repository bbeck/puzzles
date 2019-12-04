package main

import (
	"fmt"
	"strings"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	low, high := InputToRange(2019, 4)

	var count int
	for pw := low; pw <= high; pw++ {
		if IsPossiblePassword(pw) {
			count++
		}
	}

	fmt.Printf("number of possible passwords: %d\n", count)
}

func IsPossiblePassword(pw int) bool {
	if pw < 100000 || pw > 999999 {
		return false
	}

	digits := []int{
		pw / 100000 % 10,
		pw / 10000 % 10,
		pw / 1000 % 10,
		pw / 100 % 10,
		pw / 10 % 10,
		pw % 10,
	}

	return IsNonDecreasing(digits) && HasRunOfLength2(digits)
}

func HasRunOfLength2(digits []int) bool {
	for n := 0; n < 10; n++ {
		// Compute a map indicating if the digit at position i is the digit we're
		// working with.
		indices := make(map[int]bool)
		for index, digit := range digits {
			if digit != n {
				continue
			}

			indices[index] = true
		}

		// Look for a run of length 2 in the indices map
		for start := 0; start < len(digits)-1; start++ {
			if !indices[start-1] && indices[start] && indices[start+1] && !indices[start+2] {
				return true
			}
		}
	}

	return false
}

func IsNonDecreasing(digits []int) bool {
	last := digits[0]
	for i := 1; i < len(digits); i++ {
		if digits[i] < last {
			return false
		}
		last = digits[i]
	}

	return true
}

func InputToRange(year, day int) (int, int) {
	parts := strings.Split(aoc.InputToString(year, day), "-")
	return aoc.ParseInt(parts[0]), aoc.ParseInt(parts[1])
}
