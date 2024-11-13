package main

import (
	"fmt"

	"github.com/bbeck/advent-of-code/puz"
)

func main() {
	deltas := puz.InputToInts(2018, 1)

	var seen puz.Set[int]
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
