package main

import (
	"fmt"
	"log"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	var sum int
	for _, line := range aoc.InputToLines(2015, 8) {
		sum += len(line) - DecodeLength(line)
	}

	fmt.Printf("size: %d\n", sum)
}

func DecodeLength(s string) int {
	var length int
	for i := 0; i < len(s); i++ {
		if s[i] == '\\' {
			// We have a single character escape sequence
			if s[i+1] == '"' || s[i+1] == '\\' {
				length += 1
				i++
				continue
			}

			// We have a hex character escape sequence
			if s[i+1] == 'x' {
				length += 1
				i += 3
				continue
			}

			log.Fatalf("unrecognized escape sequence at position %d of %s", i, s)
		}

		if s[i] == '"' {
			// This is an unescaped quote, it's either the beginning or end of the
			// string, it should be ignored.
			continue
		}

		// This is just a normal character
		length++
	}

	return length
}
