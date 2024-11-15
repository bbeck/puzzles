package main

import (
	"fmt"

	"github.com/bbeck/advent-of-code/puz"
)

func main() {
	grid := puz.InputToStringGrid2D(2023, 23)

	var start, end puz.Point2D
	grid.ForEachPoint(func(p puz.Point2D, s string) {
		if p.Y == 0 && s == "." {
			start = p
		}
		if p.Y == grid.Height-1 && s == "." {
			end = p
		}
	})

	children := func(s State) []State {
		var children []State
		grid.ForEachOrthogonalNeighbor(s.Point2D, func(q puz.Point2D, ch string) {
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
			longest = puz.Max(longest, s.N)
		}
		return false
	}

	puz.BreadthFirstSearch(State{start, puz.Origin2D, 0}, children, goal)
	fmt.Println(longest)
}

type State struct {
	puz.Point2D
	Parent puz.Point2D
	N      int
}
