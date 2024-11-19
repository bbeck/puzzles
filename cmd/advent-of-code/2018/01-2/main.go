package main

import (
	"fmt"

	"github.com/bbeck/advent-of-code/lib"
)

func main() {
	deltas := lib.InputToInts()

	var seen lib.Set[int]
	var frequency int

outer:
	for {
		for _, delta := range deltas {
			frequency += delta
			if !seen.Add(frequency) {
				break outer
			}
		}
	}

	fmt.Println(frequency)
}
