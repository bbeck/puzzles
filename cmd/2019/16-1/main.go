package main

import (
	"fmt"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	digits := InputToDigits(2019, 16)

	for n := 0; n < 100; n++ {
		digits = RunPhase(digits)
	}

	var prefix string
	for n := 0; n < 8; n++ {
		prefix = prefix + fmt.Sprintf("%d", digits[n])
	}
	fmt.Println("prefix:", prefix)
}

func RunPhase(inputs []int) []int {
	var outputs []int

	for n := 0; n < len(inputs); n++ {
		pattern := Pattern(len(inputs), n)

		var output int
		for i := 0; i < len(inputs); i++ {
			output += inputs[i] * pattern[i]
		}
		outputs = append(outputs, aoc.AbsInt(output)%10)
	}

	return outputs
}

func Pattern(length int, digit int) []int {
	base := []int{0, 1, 0, -1}

	var pattern = make([]int, 0)
	for len(pattern) < length+1 {
		for _, b := range base {
			for n := 0; n < digit+1; n++ {
				pattern = append(pattern, b)
			}
		}
	}

	return pattern[1 : length+1]
}

func InputToDigits(year, day int) []int {
	var digits []int
	for _, b := range aoc.InputToString(year, day) {
		digits = append(digits, aoc.ParseInt(string(b)))
	}

	return digits
}
