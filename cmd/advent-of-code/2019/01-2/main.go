package main

import (
	"fmt"

	"github.com/bbeck/advent-of-code/puz"
)

func main() {
	var fuel int
	for _, mass := range puz.InputToInts() {
		fuel += Fuel(mass)
	}

	fmt.Println(fuel)
}

func Fuel(mass int) int {
	var total int
	for mass > 0 {
		fuel := puz.Max(0, mass/3-2)
		total += fuel
		mass = fuel
	}

	return total
}
