package main

import (
	"fmt"

	. "github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	grid := in.ToStringGrid2D()

	var count int
	grid.ForEachPoint(func(p Point2D, s string) {
		if s != "#" {
			return
		}
		for _, q := range p.OrthogonalNeighbors() {
			if !grid.InBoundsPoint(q) || grid.GetPoint(q) == "." {
				count++
			}
		}
	})
	fmt.Println(count)
}
