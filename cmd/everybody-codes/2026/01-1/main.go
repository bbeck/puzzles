package main

import (
	"fmt"
	"sort"

	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	var sum int
	for _, n := range in.Ints() {
		sum += FindSquare(n)
	}
	fmt.Println(sum)
}

func FindSquare(n int) int {
	size := func(s int) int { return s * s }
	s := sort.Search(n, func(s int) bool { return n <= size(s) })
	return size(s) - n
}
