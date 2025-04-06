package main

import (
	"fmt"
	"github.com/bbeck/puzzles/lib/in"

	. "github.com/bbeck/puzzles/lib"
)

func main() {
	deltas := in.Ints()

	var seen Set[int]
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
