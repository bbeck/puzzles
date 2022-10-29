package main

import (
	"fmt"
	"sort"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	adapters := aoc.InputToInts(2020, 10)
	adapters = append(adapters, 0)
	adapters = append(adapters, aoc.Max(adapters...)+3)
	sort.Ints(adapters)

	var fc aoc.FrequencyCounter[int]
	for i := 1; i < len(adapters); i++ {
		fc.Add(adapters[i] - adapters[i-1])
	}

	fmt.Println(fc.GetCount(1) * fc.GetCount(3))
}
