package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/puz"
)

func main() {
	var difference int
	for _, line := range puz.InputToLines() {
		difference += len(line) - DecodeLength(line)
	}

	fmt.Println(difference)
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
