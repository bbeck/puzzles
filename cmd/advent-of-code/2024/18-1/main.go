package main

import (
	"fmt"
	"strings"

	. "github.com/bbeck/puzzles/lib"
)

func main() {
	pts := InputToPoints()

	grid := NewGrid2D[string](71, 71)
	for n := 0; n < 1024; n++ {
		grid.SetPoint(pts[n], "#")
	}

	path, _ := BreadthFirstSearch(
		Origin2D,
		func(p Point2D) []Point2D {
			var children []Point2D
			grid.ForEachOrthogonalNeighborPoint(p, func(q Point2D, s string) {
				if s != "#" {
					children = append(children, q)
				}
			})
			return children
		},
		func(p Point2D) bool {
			return p.X == grid.Width-1 && p.Y == grid.Height-1
		})
	fmt.Println(len(path) - 1)
}

func InputToPoints() []Point2D {
	return InputLinesTo(func(s string) Point2D {
		lhs, rhs, _ := strings.Cut(s, ",")
		return Point2D{
			X: ParseInt(lhs),
			Y: ParseInt(rhs),
		}
	})
}
