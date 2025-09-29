package main

import (
	"fmt"

	. "github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	grid := in.ToGrid2D(func(_, _ int, s string) string { return s })

	// The grid is organized into three columns where the left and right columns
	// have a single entrance point.  Thus, for them a simple traveling salesman
	// algorithm will suffice.
	//
	// For the middle column we'll need to traverse to the bottom and then both
	// left and right.  We can trick the traveling salesman algorithm into doing
	// this by making the herbs at the entrances to the left and right columns
	// different.
	rLeft := grid.SubGrid(0, 0, grid.Width/3, grid.Height)
	dLeft := RoundTripLength(rLeft, Point2D{X: 84, Y: 75})

	rRight := grid.SubGrid(2*grid.Width/3, 0, grid.Width/3, grid.Height)
	dRight := RoundTripLength(rRight, Point2D{X: 0, Y: 75})

	rCenter := grid.SubGrid(grid.Width/3, 0, grid.Width/3, grid.Height)
	rCenter.Set(83, 75, "Z")
	dCenter := RoundTripLength(rCenter, Point2D{X: 42, Y: 0})

	// The edge between the center and left columns looks like
	//
	//   #.. | #.#
	//   #E. | .K#
	//   ### | ###
	//
	// In order to get from the K into the left column we need to take 2 steps on
	// the way in, and another 2 steps on the way out.  We also need to repeat
	// this on the other side for a total of 8 extra steps that weren't accounted
	// for.
	fmt.Println(dLeft + dRight + dCenter + 8)
}

func RoundTripLength(grid Grid2D[string], start Point2D) int {
	type State struct {
		P     Point2D
		Herbs BitSet
	}

	var herbs = make(map[string]int)
	grid.ForEach(func(_ int, _ int, s string) {
		if s != "." && s != "#" && s != "~" {
			if _, found := herbs[s]; !found {
				herbs[s] = len(herbs)
			}
		}
	})

	children := func(s State) []State {
		var children []State
		grid.ForEachOrthogonalNeighborPoint(s.P, func(q Point2D, v string) {
			if v == "#" || v == "~" {
				return
			}

			hs := s.Herbs
			if v, ok := herbs[v]; ok {
				hs = hs.Add(v)
			}

			children = append(children, State{P: q, Herbs: hs})
		})
		return children
	}

	goal := func(s State) bool {
		return s.P == start && s.Herbs.Size() == len(herbs)
	}

	path, _ := BreadthFirstSearch(State{P: start}, children, goal)
	return len(path) - 1
}
