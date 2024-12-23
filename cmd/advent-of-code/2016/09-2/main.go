package main

import (
	"fmt"
	"strings"

	"github.com/bbeck/puzzles/lib"
)

func main() {
	s := lib.InputToString()
	fmt.Println(DecompressedLength(s))
}

func DecompressedLength(s string) int {
	var length int
	for len(s) > 0 {
		if s[0] == '(' {
			chars, times, rest := ParseMarker(s[1:])
			length += times * DecompressedLength(rest[:chars])
			s = rest[chars:]
			continue
		}

		length++
		s = s[1:]
	}

	return length
}

func ParseMarker(s string) (int, int, string) {
	marker, rest, _ := strings.Cut(s, ")")

	var a, b int
	fmt.Sscanf(marker, "%dx%d", &a, &b)
	return a, b, rest
}
