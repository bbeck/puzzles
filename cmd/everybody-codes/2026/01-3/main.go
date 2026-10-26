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
	size := func(k int) int { return 4 * (k*k + k) }
	s := sort.Search(n, func(s int) bool { return n <= size(s) })
	return size(s) - n
}
