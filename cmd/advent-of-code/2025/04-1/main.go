package main

import (
	"fmt"

	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	grid := in.ToStringGrid2D()

	var sum int
	grid.ForEach(func(x int, y int, s string) {
		if s != "@" {
			return
		}

		var count int
		grid.ForEachNeighbor(x, y, func(_ int, _ int, t string) {
			if t == "@" {
				count++
			}
		})
		if count < 4 {
			sum++
		}
	})
	fmt.Println(sum)
}
