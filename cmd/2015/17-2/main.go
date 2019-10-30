package main

import (
	"fmt"
	"math"
	"math/bits"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	sizes := aoc.InputToInts(2015, 17)

	counts := make(map[int]int)
	EnumerateWays(sizes, 150, func(n uint) {
		counts[bits.OnesCount(n)]++
	})

	bestNumContainers := math.MaxInt64
	bestCount := 0

	for numContainers, count := range counts {
		if numContainers < bestNumContainers {
			bestNumContainers = numContainers
			bestCount = count
		}
	}

	fmt.Printf("count: %d\n", bestCount)
}

func EnumerateWays(sizes []int, amount int, fn func(n uint)) {
	valid := func(n uint) bool {
		var size int
		for i := 0; i < len(sizes); i++ {
			if n&(1<<uint(i)) != 0 {
				size += sizes[i]
			}
		}

		return size == amount
	}

	for n := uint(0); n < 1<<uint(len(sizes)); n++ {
		if valid(n) {
			fn(n)
		}
	}
}
