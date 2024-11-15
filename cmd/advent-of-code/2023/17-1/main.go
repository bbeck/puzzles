package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/puz"
)

func main() {
	pool := puz.InputToIntGrid2D()

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

	end := puz.Point2D{X: pool.Width - 1, Y: pool.Height - 1}
	_, cost, _ := puz.AStarSearch(
		Crucible{Turtle: puz.Turtle{Heading: puz.Right}},
		children,
		func(c Crucible) bool { return c.Location == end },
		func(_, c Crucible) int { return pool.GetPoint(c.Location) },
		func(c Crucible) int { return end.ManhattanDistance(c.Location) },
	)

	fmt.Println(cost)
}

type Crucible struct {
	puz.Turtle
	Count int
}

func Moves(t puz.Turtle, g puz.Grid2D[int]) []puz.Turtle {
	steps := []func(puz.Turtle) puz.Turtle{
		// Left
		func(t puz.Turtle) puz.Turtle {
			t.TurnLeft()
			t.Forward(1)
			return t
		},

		// Right
		func(t puz.Turtle) puz.Turtle {
			t.TurnRight()
			t.Forward(1)
			return t
		},

		// Straight
		func(t puz.Turtle) puz.Turtle {
			t.Forward(1)
			return t
		},
	}

	var moves []puz.Turtle
	for _, step := range steps {
		if s := step(t); g.InBoundsPoint(s.Location) {
			moves = append(moves, s)
		}
	}
	return moves
}
