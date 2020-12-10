package main

import (
	"fmt"
	"sort"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	adapters := aoc.InputToInts(2020, 10)
	adapters = append(adapters, 0)
	adapters = append(adapters, aoc.MaxInt(0, adapters...)+3)
	sort.Ints(adapters)

	differences := make(map[int]int)
	for i := 1; i < len(adapters); i++ {
		differences[adapters[i]-adapters[i-1]]++
	}

	fmt.Println(differences[1] * differences[3])
}
