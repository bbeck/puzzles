package main

import (
	"fmt"
	"strings"

	"github.com/bbeck/advent-of-code/puz"
)

func main() {
	var count int
	for _, line := range puz.InputToLines(2022, 4) {
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
	return puz.ParseInt(lhs), puz.ParseInt(rhs)
}
