package main

import (
	"fmt"

	. "github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	var start, end Point2D
	grid := in.ToGrid2D(func(x, y int, s string) string {
		if s == "S" {
			start = Point2D{X: x, Y: y}
			return "T"
		}
		if s == "E" {
			end = Point2D{X: x, Y: y}
			return "T"
		}
		return s
	})

	children := func(p Point2D) []Point2D {
		var children []Point2D
		for _, n := range Neighbors(p, grid) {
			if grid.GetPoint(n) == "T" {
				children = append(children, n)
			}
		}
		return children
	}

	goal := func(p Point2D) bool {
		return p == end
	}

	path, _ := BreadthFirstSearch(start, children, goal)
	fmt.Println(len(path) - 1)
}

func Neighbors(p Point2D, grid Grid2D[string]) []Point2D {
	var ns []Point2D

	// We are always neighbors to the left and right
	if q := p.Left(); grid.InBoundsPoint(q) {
		ns = append(ns, q)
	}
	if q := p.Right(); grid.InBoundsPoint(q) {
		ns = append(ns, q)
	}

	// The remaining neighbor depends on if the triangle is pointing upwards
	// or downwards.
	isPointingDown := p.X%2 == p.Y%2
	if isPointingDown {
		if q := p.Up(); grid.InBoundsPoint(q) {
			ns = append(ns, q)
		}
	} else {
		if q := p.Down(); grid.InBoundsPoint(q) {
			ns = append(ns, q)
		}
	}

	return ns
}
