package main

import (
	"fmt"

	. "github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	grid, current, bone := InputToGrid()

	var count int
	for current != bone {
		count++
		current = Step(current, grid)
		grid.Add(current)
	}
	fmt.Println(count)
}

var index = -1

func Step(p Point2D, grid Set[Point2D]) Point2D {
	for range 4 {
		index = (index + 1) % 4
		if index == 0 && !grid.Contains(p.Up()) {
			return p.Up()
		}
		if index == 1 && !grid.Contains(p.Right()) {
			return p.Right()
		}
		if index == 2 && !grid.Contains(p.Down()) {
			return p.Down()
		}
		if index == 3 && !grid.Contains(p.Left()) {
			return p.Left()
		}
	}

	panic("no move")
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

	return SetFrom(current), current, bone
}
