package main

import (
	"fmt"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	var difference int
	for _, line := range aoc.InputToLines(2015, 8) {
		difference += EncodeLength(line) - len(line)
	}

	fmt.Println(difference)
}

func EncodeLength(s string) int {
	var length int

	// Starting quote
	length++

	for _, c := range s {
		if c == '"' || c == '\\' {
			length += 2
			continue
		}

		length++
	}

	// Ending quote
	length++

	return length
}
