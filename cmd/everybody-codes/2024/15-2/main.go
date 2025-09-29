package main

import (
	"fmt"

	. "github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/in"
)

var Letters = SetFrom(
	"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N",
	"O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z",
)

func main() {
	grid := in.ToGrid2D(func(_, _ int, s string) string {
		return s
	})

	var start Point2D
	for x := 0; x < grid.Width; x++ {
		if grid.Get(x, 0) == "." {
			start = Point2D{X: x}
			break
		}
	}

	l := RoundTripLength(grid, start)
	fmt.Println(l)
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

	path, ok := BreadthFirstSearch(State{P: start}, children, goal)
	if !ok {
		return -1
	}
	return len(path) - 1
}
