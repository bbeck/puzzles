package main

import (
	"fmt"

	. "github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	grid, current, bone := InputToGrid()

	var count int
	for !Surrounded(bone, grid) {
		current = Step(current, grid)
		grid.Add(current)
		Flood(current, grid)

		count++
	}
	fmt.Println(count)
}

var index = -1
var dirs = []Heading{Up, Right, Down, Left}

func Step(p Point2D, grid Set[Point2D]) Point2D {
	for range 4 {
		index = (index + 1) % 4
		if q := p.Move(dirs[index]); !grid.Contains(q) {
			return q
		}
	}
	panic("no move")
}

func Surrounded(p Point2D, grid Set[Point2D]) bool {
	for _, q := range p.OrthogonalNeighbors() {
		if !grid.Contains(q) {
			return false
		}
	}
	return true
}

func Flood(p Point2D, grid Set[Point2D]) {
	tl, br := GetBounds(grid.Entries())

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

func InputToGrid() (Set[Point2D], Point2D, Point2D) {
	grid := in.ToStringGrid2D()

	var current Point2D
	var bone Point2D
	grid.ForEachPoint(func(p Point2D, s string) {
		if s == "@" {
			current = p
		}
		if s == "#" {
			bone = p
		}
	})

	return SetFrom(current, bone), current, bone
}
