package main

import (
	"fmt"
	"strings"

	"github.com/bbeck/puzzles/lib"
)

const AllItems = " abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func main() {
	var sum int
	for _, group := range lib.Chunks(lib.InputToLines(), 3) {
		s0, s1, s2 := SetFrom(group[0]), SetFrom(group[1]), SetFrom(group[2])

		common := s0.Intersect(s1).Intersect(s2).Entries()[0]
		sum += strings.IndexByte(AllItems, common)
	}
	fmt.Println(sum)
}

func SetFrom(s string) lib.Set[byte] {
	return lib.SetFrom([]byte(s)...)
}
