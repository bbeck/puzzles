package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/puz"
	"strings"
)

const AllItems = " abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func main() {
	var sum int
	for _, sack := range puz.InputToLines() {
		N := len(sack)
		lhs, rhs := SetFrom(sack[:N/2]), SetFrom(sack[N/2:])

		common := lhs.Intersect(rhs).Entries()[0]
		sum += strings.IndexByte(AllItems, common)
	}
	fmt.Println(sum)
}

func SetFrom(s string) puz.Set[byte] {
	return puz.SetFrom([]byte(s)...)
}
