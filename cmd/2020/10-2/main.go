package main

import (
	"fmt"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	adapters := aoc.InputToInts(2020, 10)

	start := 0
	adapters = append(adapters, start)

	end := aoc.MaxInt(0, adapters...) + 3
	adapters = append(adapters, end)

	count := Count(adapters, start, end, map[int]int{end: 1})
	fmt.Println(count)
}

func Count(adapters []int, start, end int, memo map[int]int) int {
	if memo[start] != 0 {
		return memo[start]
	}

	var ways int
	for _, adapter := range adapters {
		if adapter > start && adapter-start <= 3 {
			ways += Count(adapters, adapter, end, memo)
		}
	}

	memo[start] = ways
	return ways
}
