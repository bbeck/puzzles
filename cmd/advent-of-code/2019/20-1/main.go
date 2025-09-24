package main

import (
	"fmt"

	. "github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	grid, portals, start, goal := InputToMaze()

	children := func(p Point2D) []Point2D {
		var children []Point2D
		if other, found := portals[p]; found {
			children = append(children, other)
		}
		grid.ForEachOrthogonalNeighborPoint(p, func(q Point2D, isOpen bool) {
			if isOpen {
				children = append(children, q)
			}
		})
		return children
	}

	isGoal := func(p Point2D) bool {
		return p == goal
	}

	path, _ := BreadthFirstSearch(start, children, isGoal)
	fmt.Println(len(path) - 1) // the path includes the starting point
}

var Letters = SetFrom("A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z")

func InputToMaze() (Grid2D[bool], map[Point2D]Point2D, Point2D, Point2D) {
	grid := in.ToGrid2D(func(x, y int, s string) string {
		return s
	})

	labels := make(map[string][]Point2D)
	grid.ForEachPoint(func(p Point2D, s string) {
		if s != "." {
			return
		}

		if u, uu := grid.GetPoint(p.Up()), grid.GetPoint(p.Up().Up()); Letters.Contains(u) && Letters.Contains(uu) {
			labels[uu+u] = append(labels[uu+u], p)
		}

		if r, rr := grid.GetPoint(p.Right()), grid.GetPoint(p.Right().Right()); Letters.Contains(r) && Letters.Contains(rr) {
			labels[r+rr] = append(labels[r+rr], p)
		}

		if d, dd := grid.GetPoint(p.Down()), grid.GetPoint(p.Down().Down()); Letters.Contains(d) && Letters.Contains(dd) {
			labels[d+dd] = append(labels[d+dd], p)
		}

		if l, ll := grid.GetPoint(p.Left()), grid.GetPoint(p.Left().Left()); Letters.Contains(l) && Letters.Contains(ll) {
			labels[ll+l] = append(labels[ll+l], p)
		}
	})

	portals := make(map[Point2D]Point2D)
	var start, end Point2D
	for label, ps := range labels {
		switch label {
		case "AA":
			start = ps[0]
		case "ZZ":
			end = ps[0]
		default:
			portals[ps[0]] = ps[1]
			portals[ps[1]] = ps[0]
		}
	}

	open := NewGrid2D[bool](grid.Width, grid.Height)
	grid.ForEach(func(x, y int, s string) {
		open.Set(x, y, s == ".")
	})

	return open, portals, start, end
}
