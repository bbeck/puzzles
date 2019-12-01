package main

import (
	"fmt"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	var fuel int
	for _, line := range aoc.InputToLines(2019, 1) {
		mass := aoc.ParseInt(line)
		fuel += FuelRequired(mass)
	}

	fmt.Printf("total fuel required: %d\n", fuel)
}

func FuelRequired(mass int) int {
	needed := func(m int) int {
		fuel := m/3 - 2
		if fuel < 0 {
			return 0
		}
		return fuel
	}

	var total int
	for mass > 0 {
		fuel := needed(mass)
		total += fuel
		mass = fuel
	}

	return total
}
