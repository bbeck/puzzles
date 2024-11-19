package main

import (
	"fmt"
	"github.com/bbeck/puzzles/lib"
)

func main() {
	offsets := lib.InputToInts()

	var steps int
	for pc := 0; pc >= 0 && pc < len(offsets); steps++ {
		offset := offsets[pc]
		if offset >= 3 {
			offsets[pc]--
		} else {
			offsets[pc]++
		}
		pc += offset
	}

	fmt.Println(steps)
}
