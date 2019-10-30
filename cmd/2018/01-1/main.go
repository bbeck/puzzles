package main

import (
	"fmt"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	var frequency int
	for _, i := range aoc.InputToInts(2018, 1) {
		frequency += i
	}

	fmt.Printf("frequency: %d\n", frequency)
}
