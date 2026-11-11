package main

import (
	"fmt"

	. "github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	edges, nh, nv := InputToGraph()

	//
	// Use Kuhn's algorithm to compute the maximum matching.
	//

	// Given a v, which h has it been matched to
	var v2h = make(map[int]int)

	// Try to match a h to a v, not using any edges that have already been used
	var try func(h int, used []bool) bool
	try = func(h int, used []bool) bool {
		for _, v := range edges[h] {
			if used[v] {
				continue
			}
			used[v] = true

			if otherh, matched := v2h[v]; !matched || try(otherh, used) {
				v2h[v] = h
				return true
			}
		}

		return false
	}

	// Count how many matches we can make
	var matching int
	for h := range nh {
		if try(h, make([]bool, nv)) {
			matching++
		}
	}
	fmt.Println(matching)
}

func InputToGraph() (edges [][]int, nh int, nv int) {
	grid := in.ToStringGrid2D()

	var hOf = make(map[Point2D]int)
	for y := 0; y < grid.Height; y++ {
		for x := 0; x < grid.Width; x++ {
			if grid.Get(x, y) != "#" {
				continue
			}
			if x == 0 || grid.Get(x-1, y) != "#" {
				nh++
			}
			hOf[Point2D{X: x, Y: y}] = nh - 1
		}
	}

	var vOf = make(map[Point2D]int)
	for x := 0; x < grid.Width; x++ {
		for y := 0; y < grid.Height; y++ {
			if grid.Get(x, y) != "#" {
				continue
			}
			if y == 0 || grid.Get(x, y-1) != "#" {
				nv++
			}
			vOf[Point2D{X: x, Y: y}] = nv - 1
		}
	}

	edges = make([][]int, nh)
	for p, h := range hOf {
		edges[h] = append(edges[h], vOf[p])
	}

	return
}
