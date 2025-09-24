package main

import (
	"fmt"

	. "github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	digits := InputToDigits()

	// Precompute the patterns we'll need.
	var patterns [][]int
	for digit := 1; digit < len(digits)+1; digit++ {
		patterns = append(patterns, BuildPattern([]int{0, 1, 0, -1}, digit, len(digits)))
	}

	for range 100 {
		digits = FFT(digits, patterns)
	}

	for n := range 8 {
		fmt.Print(digits[n])
	}
	fmt.Println()
}

func FFT(digits []int, patterns [][]int) []int {
	output := make([]int, len(digits))
	for i := range digits {
		var sum int
		for j := range digits {
			sum += digits[j] * patterns[i][j]
		}
		output[i] = Abs(sum) % 10
	}

	return output
}

func BuildPattern(base []int, n int, length int) []int {
	var pattern []int
	for len(pattern) < length+1 {
		for _, digit := range base {
			for range n {
				pattern = append(pattern, digit)
			}
		}
	}

	return pattern[1 : length+1]
}

func InputToDigits() []int {
	var digits []int
	for in.HasNext() {
		digits = append(digits, ParseInt(string(in.Byte())))
	}
	return digits
}
