package main

import (
	"fmt"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	dish := InputToDish()
	dish = TiltUp(dish)

	var load int
	dish.ForEach(func(x int, y int, s string) {
		if s == "O" {
			load += dish.Height - y
		}
	})
	fmt.Println(load)
}

func TiltUp(dish aoc.Grid2D[string]) aoc.Grid2D[string] {
	up := func(x, y int) int {
		for dish.InBounds(x, y-1) && dish.Get(x, y-1) == "." {
			y--
		}
		return y
	}

	for y := 1; y < dish.Height; y++ {
		for x := 0; x < dish.Width; x++ {
			if dish.Get(x, y) != "O" {
				continue
			}

			newY := up(x, y)
			dish.Set(x, y, ".")
			dish.Set(x, newY, "O")
		}
	}

	return dish
}

func InputToDish() aoc.Grid2D[string] {
	return aoc.InputToGrid2D(2023, 14, func(x int, y int, s string) string {
		return s
	})
}
