package main

import (
	"fmt"
	"github.com/bbeck/puzzles/lib"
)

func main() {
	dish := lib.InputToStringGrid2D()
	dish = lib.WalkCycleWithIdentity(dish, 1_000_000_000, Cycle, ID)

	var load int
	dish.ForEach(func(x int, y int, s string) {
		if s == "O" {
			load += dish.Height - y
		}
	})
	fmt.Println(load)
}

func Cycle(dish lib.Grid2D[string]) lib.Grid2D[string] {
	up := func(x, y int) int {
		for dish.InBounds(x, y-1) && dish.Get(x, y-1) == "." {
			y--
		}
		return y
	}

	for n := 0; n < 4; n++ {
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

		dish = dish.RotateRight()
	}

	return dish
}

func ID(dish lib.Grid2D[string]) string {
	return dish.String()
}
