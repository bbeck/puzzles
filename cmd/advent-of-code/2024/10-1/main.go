package main

import (
	"fmt"
	. "github.com/bbeck/puzzles/lib"
)

func main() {
	var heads Set[Point2D]
	grid := InputToGrid2D(func(x int, y int, s string) int {
		n := ParseInt(s)
		if n == 0 {
			heads.Add(Point2D{X: x, Y: y})
		}
		return n
	})

	var sum int
	for head := range heads {
		seen := Visit(grid, head, grid.GetPoint(head))
		sum += len(seen)
	}
	fmt.Println(sum)
}

func Visit(grid Grid2D[int], p Point2D, v int) Set[Point2D] {
	if v == 9 {
		return SetFrom(p)
	}

	var seen Set[Point2D]
	for _, n := range p.OrthogonalNeighbors() {
		if grid.InBoundsPoint(n) && grid.GetPoint(n) == v+1 {
			ns := Visit(grid, n, v+1)
			seen = seen.Union(ns)
		}
	}
	return seen
}
