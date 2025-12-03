package main

import (
	"fmt"

	. "github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	grid := in.ToGrid2D(func(_, _ int, s string) string { return s })

	var count int
	grid.ForEachPoint(func(p Point2D, ps string) {
		if ps != "T" {
			return
		}
		for _, q := range Neighbors(p, grid) {
			if grid.GetPoint(q) != "T" {
				continue
			}

			count++
		}
	})

	// We'll double count since we visit the pair through both trampolines.
	fmt.Println(count / 2)
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
