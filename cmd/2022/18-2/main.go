package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	ps := aoc.SetFrom(InputToPoints()...)
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

func Outside(ps aoc.Set[aoc.Point3D]) aoc.Set[aoc.Point3D] {
	// Determine the bounding box of the points (expanded by 1 unit)
	tl, br := aoc.GetBounds3D(ps.Entries())
	tl = aoc.Point3D{X: tl.X - 1, Y: tl.Y - 1, Z: tl.Z - 1}
	br = aoc.Point3D{X: br.X + 1, Y: br.Y + 2, Z: br.Z + 1}

	inRange := func(v, min, max int) bool {
		return min <= v && v <= max
	}

	// Perform a flood fill from the corner of the bounding box and see which
	// points are reachable.
	children := func(p aoc.Point3D) []aoc.Point3D {
		var children []aoc.Point3D
		for _, n := range p.OrthogonalNeighbors() {
			if inRange(n.X, tl.X, br.X) && inRange(n.Y, tl.Y, br.Y) && inRange(n.Z, tl.Z, br.Z) && !ps.Contains(n) {
				children = append(children, n)
			}
		}
		return children
	}

	var outside aoc.Set[aoc.Point3D]
	goal := func(p aoc.Point3D) bool {
		outside.Add(p)
		return false
	}

	aoc.BreadthFirstSearch(tl, children, goal)
	return outside
}

func InputToPoints() []aoc.Point3D {
	return aoc.InputLinesTo(2022, 18, func(line string) (aoc.Point3D, error) {
		var p aoc.Point3D
		_, err := fmt.Sscanf(line, "%d,%d,%d", &p.X, &p.Y, &p.Z)
		return p, err
	})
}
