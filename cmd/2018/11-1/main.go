package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	serial := aoc.InputToInt(2018, 11)

	grid := aoc.NewGrid2D[int](301, 301)
	for x := 1; x < grid.Width; x++ {
		rack := x + 10

		for y := 1; y < grid.Height; y++ {
			power := (rack*y + serial) * rack
			power = (power%1000)/100 - 5
			grid.Add(x, y, power)
		}
	}

	var best int
	var p aoc.Point2D
	for x := 1; x < grid.Width-3; x++ {
		for y := 1; y < grid.Height-3; y++ {
			var total int
			for dx := 0; dx < 3; dx++ {
				for dy := 0; dy < 3; dy++ {
					total += grid.Get(x+dx, y+dy)
				}
			}

			if total > best {
				best = total
				p = aoc.Point2D{X: x, Y: y}
			}
		}
	}

	fmt.Printf("%d,%d\n", p.X, p.Y)
}
