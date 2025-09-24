package main

import (
	"fmt"

	"github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	var fuel int
	for _, mass := range in.Ints() {
		fuel += Fuel(mass)
	}

	fmt.Println(fuel)
}

func Fuel(mass int) int {
	var total int
	for mass > 0 {
		fuel := lib.Max(0, mass/3-2)
		total += fuel
		mass = fuel
	}

	return total
}
