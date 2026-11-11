package main

import (
	"fmt"

	. "github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	shapes := FindShapes()

	var sum int
	for _, shape := range shapes {
		tl, br := GetBounds(shape)
		W, H := br.X-tl.X+1, br.Y-tl.Y+1

		var full int
		var rCounts, cCounts = make(map[int]int), make(map[int]int)
		for _, p := range shape {
			rCounts[p.X]++
			if rCounts[p.X] == H {
				full++
			}

			cCounts[p.Y]++
			if cCounts[p.Y] == W {
				full++
			}
		}

		sum += Min(W, H, full)
	}
	fmt.Println(sum)
}

func FindShapes() [][]Point2D {
	var shapes [][]Point2D

	var seen Set[Point2D] // Points already included in a shape.

	grid := in.ToStringGrid2D()
	grid.ForEachPoint(func(p Point2D, s string) {
		if s != "#" || !seen.Add(p) {
			return
		}

		var shape = []Point2D{p}
		for i := 0; i < len(shape); i++ {
			grid.ForEachOrthogonalNeighborPoint(shape[i], func(n Point2D, s string) {
				if s == "#" && seen.Add(n) {
					shape = append(shape, n)
				}
			})
		}
		shapes = append(shapes, shape)
	})
	return shapes
}
