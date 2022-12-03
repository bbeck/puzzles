package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	var sum int
	for _, sack := range aoc.InputToLines(2022, 3) {
		N := len(sack)
		lhs, rhs := SetFrom(sack[:N/2]), SetFrom(sack[N/2:])

		common := lhs.Intersect(rhs).Entries()[0]
		switch {
		case 'a' <= common && common <= 'z':
			sum += int(common-'a') + 1
		case 'A' <= common && common <= 'Z':
			sum += int(common-'A') + 27
		}
	}
	fmt.Println(sum)
}

func SetFrom(s string) aoc.Set[byte] {
	return aoc.SetFrom([]byte(s)...)
}
