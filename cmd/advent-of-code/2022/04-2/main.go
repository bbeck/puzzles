package main

import (
	"fmt"
	"strings"

	"github.com/bbeck/advent-of-code/lib"
)

func main() {
	var count int
	for _, line := range lib.InputToLines() {
		lhs, rhs, _ := strings.Cut(line, ",")
		alo, ahi := ParseRange(lhs)
		blo, bhi := ParseRange(rhs)
		if ahi >= blo && bhi >= alo {
			count++
		}
	}
	fmt.Println(count)
}

func ParseRange(s string) (int, int) {
	lhs, rhs, _ := strings.Cut(s, "-")
	return lib.ParseInt(lhs), lib.ParseInt(rhs)
}
