package main

import (
	"fmt"
	"github.com/bbeck/puzzles/lib"
)

func main() {
	ps := lib.SetFrom(InputToPoints()...)
	outside := Outside(ps)

	var count int
	for p := range ps {
		for _, n := range p.OrthogonalNeighbors() {
			if !ps.Contains(n) && outside.Contains(n) {
				count++
			}
		}
	}
	fmt.Println(count)
}

func Outside(ps lib.Set[lib.Point3D]) lib.Set[lib.Point3D] {
	// Determine the bounding box of the points (expanded by 1 unit)
	tl, br := lib.GetBounds3D(ps.Entries())
	tl = lib.Point3D{X: tl.X - 1, Y: tl.Y - 1, Z: tl.Z - 1}
	br = lib.Point3D{X: br.X + 1, Y: br.Y + 2, Z: br.Z + 1}

	inRange := func(v, min, max int) bool {
		return min <= v && v <= max
	}

	// Perform a flood fill from the corner of the bounding box and see which
	// points are reachable.
	children := func(p lib.Point3D) []lib.Point3D {
		var children []lib.Point3D
		for _, n := range p.OrthogonalNeighbors() {
			if inRange(n.X, tl.X, br.X) && inRange(n.Y, tl.Y, br.Y) && inRange(n.Z, tl.Z, br.Z) && !ps.Contains(n) {
				children = append(children, n)
			}
		}
		return children
	}

	var outside lib.Set[lib.Point3D]
	goal := func(p lib.Point3D) bool {
		outside.Add(p)
		return false
	}

	lib.BreadthFirstSearch(tl, children, goal)
	return outside
}

func InputToPoints() []lib.Point3D {
	return lib.InputLinesTo(func(line string) lib.Point3D {
		var p lib.Point3D
		fmt.Sscanf(line, "%d,%d,%d", &p.X, &p.Y, &p.Z)
		return p
	})
}
