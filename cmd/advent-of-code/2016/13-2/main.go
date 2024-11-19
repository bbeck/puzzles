package main

import (
	"fmt"
	"github.com/bbeck/puzzles/lib"
	"math/bits"
)

func main() {
	m := Maze(lib.InputToInt())
	start := lib.Point2D{X: 1, Y: 1}

	// Remember the distance from the start of each location we visit.
	distances := make(map[lib.Point2D]int)

	children := func(p lib.Point2D) []lib.Point2D {
		// If this node is too far away from the start then don't explore any of
		// its children.
		if distances[p] >= 50 {
			return nil
		}

		var children []lib.Point2D
		for _, n := range p.OrthogonalNeighbors() {
			if n.X < 0 || n.Y < 0 || !m.IsOpen(n) {
				continue
			}

			// Skip over this neighbor if we've previously processed it.  Since we're
			// performing a BFS, if we've previously processed it we've found a path
			// that's either the same length or shorter than this one.
			if _, found := distances[n]; found {
				continue
			}

			children = append(children, n)
			distances[n] = distances[p] + 1
		}

		return children
	}

	goal := func(lib.Point2D) bool {
		return false
	}

	lib.BreadthFirstSearch(start, children, goal)

	var count int
	for _, distance := range distances {
		if distance <= 50 {
			count++
		}
	}
	fmt.Println(count)
}

type Maze int

func (m Maze) IsOpen(p lib.Point2D) bool {
	n := uint(p.X*p.X + 3*p.X + 2*p.X*p.Y + p.Y + p.Y*p.Y + int(m))
	return bits.OnesCount(n)%2 == 0
}
