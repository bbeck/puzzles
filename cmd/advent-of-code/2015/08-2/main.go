package main

import (
	"fmt"

	"github.com/bbeck/puzzles/lib"
)

func main() {
	var difference int
	for _, line := range lib.InputToLines() {
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
