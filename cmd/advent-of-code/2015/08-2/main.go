package main

import (
	"fmt"

	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	var difference int
	for in.HasNext() {
		line := in.Line()
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
