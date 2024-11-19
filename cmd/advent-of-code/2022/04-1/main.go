package main

import (
	"fmt"
	"strings"

	"github.com/bbeck/puzzles/lib"
)

func main() {
	var count int
	for _, line := range lib.InputToLines() {
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
	return lib.ParseInt(lhs), lib.ParseInt(rhs)
}
