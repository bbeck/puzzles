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
		sum += Min(br.X-tl.X+1, br.Y-tl.Y+1)
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
