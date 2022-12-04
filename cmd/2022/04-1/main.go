package main

import (
	"fmt"
	"strings"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	var count int
	for _, line := range aoc.InputToLines(2022, 4) {
		lhs, rhs, _ := strings.Cut(line, ",")
		alo, ahi := ParseRange(lhs)
		blo, bhi := ParseRange(rhs)
		if (alo <= blo && bhi <= ahi) || (blo <= alo && ahi <= bhi) {
			count++
		}
	}
	fmt.Println(count)
}

func ParseRange(s string) (int, int) {
	lhs, rhs, _ := strings.Cut(s, "-")
	return aoc.ParseInt(lhs), aoc.ParseInt(rhs)
}
