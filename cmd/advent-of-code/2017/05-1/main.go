package main

import (
	"fmt"

	"github.com/bbeck/advent-of-code/puz"
)

func main() {
	offsets := puz.InputToInts()

	var steps int
	for pc := 0; pc >= 0 && pc < len(offsets); steps++ {
		offset := offsets[pc]
		offsets[pc]++
		pc += offset
	}

	fmt.Println(steps)
}
