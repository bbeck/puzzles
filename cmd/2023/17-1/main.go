package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	pool := aoc.InputToIntGrid2D(2023, 17)

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

	end := aoc.Point2D{X: pool.Width - 1, Y: pool.Height - 1}
	_, cost, _ := aoc.AStarSearch(
		Crucible{Turtle: aoc.Turtle{Heading: aoc.Right}},
		children,
		func(c Crucible) bool { return c.Location == end },
		func(_, c Crucible) int { return pool.GetPoint(c.Location) },
		func(c Crucible) int { return end.ManhattanDistance(c.Location) },
	)

	fmt.Println(cost)
}

type Crucible struct {
	aoc.Turtle
	Count int
}

func Moves(t aoc.Turtle, g aoc.Grid2D[int]) []aoc.Turtle {
	steps := []func(aoc.Turtle) aoc.Turtle{
		// Left
		func(t aoc.Turtle) aoc.Turtle {
			t.TurnLeft()
			t.Forward(1)
			return t
		},

		// Right
		func(t aoc.Turtle) aoc.Turtle {
			t.TurnRight()
			t.Forward(1)
			return t
		},

		// Straight
		func(t aoc.Turtle) aoc.Turtle {
			t.Forward(1)
			return t
		},
	}

	var moves []aoc.Turtle
	for _, step := range steps {
		if s := step(t); g.InBoundsPoint(s.Location) {
			moves = append(moves, s)
		}
	}
	return moves
}
