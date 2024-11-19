package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/lib"
)

func main() {
	pool := lib.InputToIntGrid2D()

	children := func(c Crucible) []Crucible {
		var children []Crucible
		for _, t := range Moves(c.Turtle, pool) {
			if t.Heading != c.Heading {
				children = append(children, Crucible{Turtle: t, Count: 1})
				continue
			}

			if c.Count < 3 {
				children = append(children, Crucible{Turtle: t, Count: c.Count + 1})
			}
		}
		return children
	}

	end := lib.Point2D{X: pool.Width - 1, Y: pool.Height - 1}
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
