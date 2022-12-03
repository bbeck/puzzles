package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	var sum int
	for _, group := range aoc.Chunk(aoc.InputToLines(2022, 3), 3) {
		s0, s1, s2 := SetFrom(group[0]), SetFrom(group[1]), SetFrom(group[2])
		common := s0.Intersect(s1).Intersect(s2).Entries()[0]

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
