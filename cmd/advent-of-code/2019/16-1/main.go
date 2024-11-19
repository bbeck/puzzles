package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/lib"
)

func main() {
	digits := InputToDigits()

	// Precompute the patterns we'll need.
	var patterns [][]int
	for digit := 1; digit < len(digits)+1; digit++ {
		patterns = append(patterns, BuildPattern([]int{0, 1, 0, -1}, digit, len(digits)))
	}

	for n := 0; n < 100; n++ {
		digits = FFT(digits, patterns)
	}

	for n := 0; n < 8; n++ {
		fmt.Print(digits[n])
	}
	fmt.Println()
}

func FFT(digits []int, patterns [][]int) []int {
	output := make([]int, len(digits))
	for i := 0; i < len(digits); i++ {
		var sum int
		for j := 0; j < len(digits); j++ {
			sum += digits[j] * patterns[i][j]
		}
		output[i] = lib.Abs(sum) % 10
	}

	return output
}

func BuildPattern(base []int, n int, length int) []int {
	var pattern []int
	for len(pattern) < length+1 {
		for _, digit := range base {
			for c := 0; c < n; c++ {
				pattern = append(pattern, digit)
			}
		}
	}

	return pattern[1 : length+1]
}

func InputToDigits() []int {
	var digits []int
	for _, s := range lib.InputToString() {
		digits = append(digits, lib.ParseInt(string(s)))
	}

	return digits
}
