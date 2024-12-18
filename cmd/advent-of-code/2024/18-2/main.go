package main

import (
	"fmt"
	"sort"
	"strings"

	. "github.com/bbeck/puzzles/lib"
)

func main() {
	pts := InputToPoints()

	idx := sort.Search(len(pts), func(n int) bool {
		grid := NewGrid2D[string](71, 71)
		for i := 0; i <= n; i++ {
			grid.SetPoint(pts[i], "#")
		}

		return !HasPath(grid)
	})
	fmt.Printf("%d,%d\n", pts[idx].X, pts[idx].Y)
}

func HasPath(grid Grid2D[string]) bool {
	_, ok := BreadthFirstSearch(
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
	return ok
}

func InputToPoints() []Point2D {
	return InputLinesTo(func(s string) Point2D {
		lhs, rhs, _ := strings.Cut(s, ",")
		return Point2D{X: ParseInt(lhs), Y: ParseInt(rhs)}
	})
}
