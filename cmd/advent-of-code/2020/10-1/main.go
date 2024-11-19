package main

import (
	"fmt"
	"sort"

	"github.com/bbeck/puzzles/lib"
)

func main() {
	adapters := lib.InputToInts()
	adapters = append(adapters, 0)
	adapters = append(adapters, lib.Max(adapters...)+3)
	sort.Ints(adapters)

	var fc lib.FrequencyCounter[int]
	for i := 1; i < len(adapters); i++ {
		fc.Add(adapters[i] - adapters[i-1])
	}

	fmt.Println(fc.GetCount(1) * fc.GetCount(3))
}
