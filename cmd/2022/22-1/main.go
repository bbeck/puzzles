package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/aoc"
	"strings"
)

func main() {
	board, directions := InputToBoard()
	var t aoc.Turtle

	// Take one step forward
	step := func(t aoc.Turtle) aoc.Turtle {
		before := t

		t.Forward(1)

		if !board.InBoundsPoint(t.Location) || board.GetPoint(t.Location) == 0 {
			// We're out of bounds, wrap around to the other side of this row/column.
			if t.Heading == aoc.Right {
				t.Location.X = 0
			} else if t.Heading == aoc.Left {
				t.Location.X = board.Width - 1
			} else if t.Heading == aoc.Up {
				t.Location.Y = board.Height - 1
			} else if t.Heading == aoc.Down {
				t.Location.Y = 0
			}

			// Walk forward until we're back on the board.
			for board.GetPoint(t.Location) == 0 {
				t.Forward(1)
			}

		}

		if board.GetPoint(t.Location) == '#' {
			// We hit a wall, go back to our previous location.
			return before
		}

		return t
	}

	// Position the turtle at the starting location.
	t.Heading = aoc.Right
	for {
		if board.GetPoint(t.Location) == '.' {
			break
		}
		t.Forward(1)
	}

	for _, dir := range directions {
		switch dir {
		case "L":
			t.TurnLeft()

		case "R":
			t.TurnRight()

		default:
			for i := 0; i < aoc.ParseInt(dir); i++ {
				t = step(t)
			}
		}
	}

	x, y := t.Location.X+1, t.Location.Y+1
	var dir int
	switch t.Heading {
	case aoc.Right:
		dir = 0
	case aoc.Down:
		dir = 1
	case aoc.Left:
		dir = 2
	case aoc.Up:
		dir = 3
	}
	fmt.Println(1000*y + 4*x + dir)
}

func InputToBoard() (aoc.Grid2D[rune], []string) {
	lines := aoc.InputToLines(2022, 22)

	var W, H int
	for y := 0; y < len(lines)-2; y++ {
		W = aoc.Max(W, len(lines[y]))
		H = y + 1
	}

	board := aoc.NewGrid2D[rune](W, H)
	for y := 0; y < len(lines)-2; y++ {
		for x, c := range lines[y] {
			if c != ' ' {
				board.Set(x, y, c)
			}
		}
	}

	directions := lines[len(lines)-1]
	directions = strings.ReplaceAll(directions, "L", " L ")
	directions = strings.ReplaceAll(directions, "R", " R ")

	return board, strings.Fields(directions)
}
