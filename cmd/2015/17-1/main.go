package main

import (
	"fmt"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	sizes := aoc.InputToInts(2015, 17)
	count := CountWays(sizes, 150)

	fmt.Printf("count: %d\n", count)
}

func CountWays(sizes []int, amount int) int {
	valid := func(n int) bool {
		var size int
		for i := uint(0); i < uint(len(sizes)); i++ {
			if n&(1<<i) != 0 {
				size += sizes[i]
			}
		}

		return size == amount
	}

	var count int
	for n := 0; n < 1<<uint(len(sizes)); n++ {
		if valid(n) {
			count++
		}
	}

	return count
}
