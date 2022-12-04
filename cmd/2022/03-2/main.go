package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/aoc"
	"strings"
)

const AllItems = " abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func main() {
	var sum int
	for _, group := range aoc.Chunk(aoc.InputToLines(2022, 3), 3) {
		s0, s1, s2 := SetFrom(group[0]), SetFrom(group[1]), SetFrom(group[2])

		common := s0.Intersect(s1).Intersect(s2).Entries()[0]
		sum += strings.IndexByte(AllItems, common)
	}
	fmt.Println(sum)
}

func SetFrom(s string) aoc.Set[byte] {
	return aoc.SetFrom([]byte(s)...)
}
