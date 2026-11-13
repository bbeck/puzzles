package main

import (
	"fmt"

	. "github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	grid := in.ToStringGrid2D()

	var internal int
	grid.ForEachPoint(func(p Point2D, s string) {
		grid.ForEachOrthogonalNeighborPoint(p, func(q Point2D, t string) {
			if s != t {
				internal++
			}
		})
	})

	var external = 2*grid.Width + 2*grid.Height
	fmt.Println(external + internal/2)
}
