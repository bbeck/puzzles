package main

import (
	"fmt"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	var fuel int
	for _, mass := range aoc.InputToInts(2019, 1) {
		fuel += mass/3 - 2
	}

	fmt.Println(fuel)
}
