package main

import (
	"fmt"
	"github.com/bbeck/puzzles/lib"
)

func main() {
	var adapters lib.Set[int]
	adapters.Add(lib.InputToInts()...)

	start, end := 0, lib.Max(adapters.Entries()...)
	adapters.Add(start, end)

	fmt.Println(Count(adapters, start, end))
}

func Count(adapters lib.Set[int], start, end int) int {
	memo := map[int]int{end: 1}

	var helper func(start int) int
	helper = func(start int) int {
		if count, found := memo[start]; found {
			return count
		}

		var ways int
		for adapter := start + 1; adapter <= start+3; adapter++ {
			if adapters.Contains(adapter) {
				ways += helper(adapter)
			}
		}

		memo[start] = ways
		return ways
	}

	return helper(start)
}
