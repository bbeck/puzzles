package main

import (
	"fmt"
	"sort"

	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	var sum int
	for _, n := range in.Ints() {
		sum += FindSize(n)
	}
	fmt.Println(sum)
}

func FindSize(n int) int {
	size := func(s int) int { return 4*s - 4 }
	s := sort.Search(n+5, func(s int) bool { return n <= size(s) && 2 < s })
	return size(s) - n
}
