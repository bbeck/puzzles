package main

import (
	"fmt"

	. "github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	grid, portals, depths, start, goal := InputToMaze()
	start3d := Point3D{X: start.X, Y: start.Y, Z: 0}
	goal3d := Point3D{X: goal.X, Y: goal.Y, Z: 0}

	children := func(p3d Point3D) []Point3D {
		p := Point2D{X: p3d.X, Y: p3d.Y}

		var children []Point3D
		if other, found := portals[p]; found {
			if depth := p3d.Z + depths[p]; depth >= 0 {
				children = append(children, Point3D{X: other.X, Y: other.Y, Z: p3d.Z + depths[p]})
			}
		}

		for _, child := range p.OrthogonalNeighbors() {
			if grid.GetPoint(child) {
				children = append(children, Point3D{X: child.X, Y: child.Y, Z: p3d.Z})
			}
		}

		return children
	}

	isGoal := func(p Point3D) bool {
		return p == goal3d
	}

	path, _ := BreadthFirstSearch(start3d, children, isGoal)
	fmt.Println(len(path) - 1) // the path includes the starting point
}

var Letters = SetFrom("A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z")

func InputToMaze() (Grid2D[bool], map[Point2D]Point2D, map[Point2D]int, Point2D, Point2D) {
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
	depths := make(map[Point2D]int)
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

			for _, p := range ps {
				if p.X == 2 || p.X == grid.Width-3 || p.Y == 2 || p.Y == grid.Height-3 {
					depths[p] = -1
				} else {
					depths[p] = 1
				}
			}
		}
	}

	open := NewGrid2D[bool](grid.Width, grid.Height)
	grid.ForEach(func(x, y int, s string) {
		open.Set(x, y, s == ".")
	})

	return open, portals, depths, start, end
}
