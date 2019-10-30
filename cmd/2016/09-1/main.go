package main

import (
	"fmt"
	"strings"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	var sum int
	for _, line := range aoc.InputToLines(2016, 9) {
		sum += DecompressedLength(line)
	}

	fmt.Printf("total: %d\n", sum)
}

func DecompressedLength(s string) int {
	var length int

	for len(s) > 0 {
		switch s[0] {
		case '(':
			var chars, times int
			chars, times, s = ParseMarker(s)
			length += times * DecompressedLength(s[:chars])
			if chars <= len(s) {
				s = s[chars:]
			} else {
				s = ""
			}

		default:
			length++
			s = s[1:]
		}
	}

	return length
}

// ParseMarker takes a string that has a marker at the beginning, removes it
// and returns the sizes and the string without the marker.
func ParseMarker(s string) (int, int, string) {
	var pieces []string
	pieces = strings.SplitN(s[1:], "x", 2)
	length := aoc.ParseInt(pieces[0]) // the number of characters to copy

	pieces = strings.SplitN(pieces[1], ")", 2)
	times := aoc.ParseInt(pieces[0]) // the number of times to paste

	return length, times, pieces[1]
}
