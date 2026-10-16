package main

import (
	"fmt"

	. "github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	grid, current, bones := InputToGrid()

	// The only thing that can change the bounds is when we add a new point to the
	// grid.  So we'll manually keep track of our bounds and incrementally update
	// them whenever we take a step.
	tl, br := GetBounds(grid.Entries())

	// The grids sometimes have fully enclosed sections.  Just flood the area
	// around each bone before we start taking steps so that we can fill them in.
	// If we don't do this then those bones will never be able to be surrounded.
	for b := range bones {
		Flood(b, grid, tl, br)
	}

	var count int
	for !Surrounded(bones, grid) {
		current = Step(current, grid)
		grid.Add(current)

		// We just added a point, update our bounds.
		tl = Point2D{X: Min(tl.X, current.X), Y: Min(tl.Y, current.Y)}
		br = Point2D{X: Max(br.X, current.X), Y: Max(br.Y, current.Y)}

		Flood(current, grid, tl, br)

		count++
	}
	fmt.Println(count)
}

var index = -1
var dirs = []Heading{
	Up, Up, Up,
	Right, Right, Right,
	Down, Down, Down,
	Left, Left, Left,
}

func Step(p Point2D, grid Set[Point2D]) Point2D {
	for range len(dirs) {
		index = (index + 1) % len(dirs)
		if q := p.Move(dirs[index]); !grid.Contains(q) {
			return q
		}
	}
	panic("no move")
}

var neighbors Set[Point2D]

func Surrounded(bones Set[Point2D], grid Set[Point2D]) bool {
	if len(neighbors) == 0 {
		for p := range bones {
			neighbors.Add(p.Up(), p.Right(), p.Down(), p.Left())
		}
	}
	for q := range neighbors {
		if !grid.Contains(q) {
			return false
		}
	}
	return true
}

func Flood(p Point2D, grid Set[Point2D], tl, br Point2D) {
	var flood func(p Point2D, seen Set[Point2D]) bool
	flood = func(p Point2D, seen Set[Point2D]) bool {
		if grid.Contains(p) || !seen.Add(p) {
			return true
		}
		if p.X < tl.X || p.X > br.X || p.Y < tl.Y || p.Y > br.Y {
			return false
		}

		for _, q := range p.OrthogonalNeighbors() {
			if !flood(q, seen) {
				return false
			}
		}

		return true
	}

	for _, candidate := range p.OrthogonalNeighbors() {
		var seen = SetFrom[Point2D]()
		if flood(candidate, seen) {
			grid.Add(seen.Entries()...)
		}
	}
}

func InputToGrid() (Set[Point2D], Point2D, Set[Point2D]) {
	grid := in.ToStringGrid2D()

	var current Point2D
	var bones Set[Point2D]
	grid.ForEachPoint(func(p Point2D, s string) {
		if s == "@" {
			current = p
		}
		if s == "#" {
			bones.Add(p)
		}
	})

	return SetFrom(current).Union(bones), current, bones
}
