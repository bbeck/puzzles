package main

import (
	"fmt"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	s := aoc.InputToString(2015, 10)
	for i := 0; i < 50; i++ {
		s = LookAndSay(s)
	}

	fmt.Printf("length: %d\n", len(s))
}

func LookAndSay(s string) string {
	output := ""

	last, count := s[0], 1
	for i := 1; i < len(s); i++ {
		if s[i] != last {
			output = output + fmt.Sprintf("%d%c", count, last)
			last = s[i]
			count = 1
			continue
		}

		count++
	}

	output = output + fmt.Sprintf("%d%c", count, last)
	return output
}
