package main

import (
	"fmt"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	grid := aoc.InputToStringGrid2D(2023, 23)

	var start, end aoc.Point2D
	grid.ForEachPoint(func(p aoc.Point2D, s string) {
		if p.Y == 0 && s == "." {
			start = p
		}
		if p.Y == grid.Height-1 && s == "." {
			end = p
		}
	})

	children := func(s State) []State {
		var children []State
		grid.ForEachOrthogonalNeighbor(s.Point2D, func(q aoc.Point2D, ch string) {
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
			longest = aoc.Max(longest, s.N)
		}
		return false
	}

	aoc.BreadthFirstSearch(State{start, aoc.Origin2D, 0}, children, goal)
	fmt.Println(longest)
}

type State struct {
	aoc.Point2D
	Parent aoc.Point2D
	N      int
}
