package main

import (
	"fmt"
	"github.com/bbeck/puzzles/lib"
)

func main() {
	pool := lib.InputToIntGrid2D()
	end := lib.Point2D{X: pool.Width - 1, Y: pool.Height - 1}

	children := func(c Crucible) []Crucible {
		var children []Crucible
		for _, t := range Moves(c.Turtle, pool) {
			isEnd := t.Location == end
			isTurn := t.Heading != c.Heading

			switch {
			// We can never take more than 10 steps in any direction.
			case c.Count > 10:
				continue

			// Must always end going straight and have taken at least 4 steps.
			case isEnd && (isTurn || c.Count < 3):
				continue

			// Must always go at least 4 steps before turning.
			case isTurn && c.Count < 4:
				continue
			}

			if isTurn {
				children = append(children, Crucible{Turtle: t, Count: 1})
			} else {
				children = append(children, Crucible{Turtle: t, Count: c.Count + 1})
			}
		}
		return children
	}

	_, cost, _ := lib.AStarSearch(
		Crucible{Turtle: lib.Turtle{Heading: lib.Right}},
		children,
		func(c Crucible) bool { return c.Location == end },
		func(_, c Crucible) int { return pool.GetPoint(c.Location) },
		func(c Crucible) int { return end.ManhattanDistance(c.Location) },
	)

	fmt.Println(cost)
}

type Crucible struct {
	lib.Turtle
	Count int
}

func Moves(t lib.Turtle, g lib.Grid2D[int]) []lib.Turtle {
	steps := []func(lib.Turtle) lib.Turtle{
		// Left
		func(t lib.Turtle) lib.Turtle {
			t.TurnLeft()
			t.Forward(1)
			return t
		},

		// Right
		func(t lib.Turtle) lib.Turtle {
			t.TurnRight()
			t.Forward(1)
			return t
		},

		// Straight
		func(t lib.Turtle) lib.Turtle {
			t.Forward(1)
			return t
		},
	}

	var moves []lib.Turtle
	for _, step := range steps {
		if s := step(t); g.InBoundsPoint(s.Location) {
			moves = append(moves, s)
		}
	}
	return moves
}
