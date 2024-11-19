package main

import (
	"fmt"

	"github.com/bbeck/puzzles/lib"
)

func main() {
	grid := lib.InputToStringGrid2D()

	var start, end lib.Point2D
	grid.ForEachPoint(func(p lib.Point2D, s string) {
		if p.Y == 0 && s == "." {
			start = p
		}
		if p.Y == grid.Height-1 && s == "." {
			end = p
		}
	})

	children := func(s State) []State {
		var children []State
		grid.ForEachOrthogonalNeighbor(s.Point2D, func(q lib.Point2D, ch string) {
			if s.Parent == q {
				return
			}
			if ch == "." {
				children = append(children, State{q, s.Point2D, s.N + 1})
			}
			if ch == ">" && s.Point2D.Right() == q {
				children = append(children, State{q, s.Point2D, s.N + 1})
			}
			if ch == "v" && s.Point2D.Down() == q {
				children = append(children, State{q, s.Point2D, s.N + 1})
			}
			if ch == "<" && s.Point2D.Left() == q {
				children = append(children, State{q, s.Point2D, s.N + 1})
			}
			if ch == "^" && s.Point2D.Up() == q {
				children = append(children, State{q, s.Point2D, s.N + 1})
			}
		})
		return children
	}

	var longest int
	goal := func(s State) bool {
		if s.Point2D == end {
			longest = lib.Max(longest, s.N)
		}
		return false
	}

	lib.BreadthFirstSearch(State{start, lib.Origin2D, 0}, children, goal)
	fmt.Println(longest)
}

type State struct {
	lib.Point2D
	Parent lib.Point2D
	N      int
}
