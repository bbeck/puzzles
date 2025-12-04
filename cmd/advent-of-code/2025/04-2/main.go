package main

import (
	"fmt"

	. "github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	grid := in.ToStringGrid2D()

	var removed Set[Point2D]
	for {
		toRemove := Run(grid)
		if len(toRemove) == 0 {
			break
		}

		for p := range toRemove {
			grid.SetPoint(p, ".")
		}
		removed = removed.Union(toRemove)
	}
	fmt.Println(len(removed))
}

func Run(grid Grid2D[string]) Set[Point2D] {
	var remove Set[Point2D]
	grid.ForEachPoint(func(p Point2D, s string) {
		if s != "@" {
			return
		}

		var count int
		grid.ForEachNeighborPoint(p, func(q Point2D, t string) {
			if t == "@" {
				count++
			}
		})
		if count < 4 {
			remove.Add(p)
		}
	})

	return remove
}
