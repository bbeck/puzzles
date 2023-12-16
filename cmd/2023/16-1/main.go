package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	turtle := aoc.Turtle{Location: aoc.Origin2D.Left(), Heading: aoc.Right}
	grid := aoc.InputToStringGrid2D(2023, 16)
	energized := aoc.NewGrid2D[string](grid.Width, grid.Height)

	Walk(turtle, grid, energized)

	var sum int
	energized.ForEach(func(x int, y int, s string) {
		if s == "#" {
			sum++
		}
	})
	fmt.Println(sum)
}

func Walk(t aoc.Turtle, grid, energized aoc.Grid2D[string]) {
	var seen aoc.Set[aoc.Turtle]

	var step func(t aoc.Turtle)
	step = func(t aoc.Turtle) {
		if !seen.Add(t) {
			return
		}

		t.Forward(1)

		location := t.Location
		if !grid.InBoundsPoint(location) {
			return
		}
		energized.SetPoint(location, "#")

		cell := grid.GetPoint(location)
		heading := t.Heading

		switch {
		case (cell == "|" && (heading == aoc.Right || heading == aoc.Left)) ||
			(cell == "-" && (heading == aoc.Up || heading == aoc.Down)):
			t.TurnLeft()
			step(t)
			t.TurnRight()
			t.TurnRight()
			step(t)

		case cell == "\\":
			if heading == aoc.Up || heading == aoc.Down {
				t.TurnLeft()
			} else {
				t.TurnRight()
			}
			step(t)

		case cell == "/":
			if heading == aoc.Up || heading == aoc.Down {
				t.TurnRight()
			} else {
				t.TurnLeft()
			}
			step(t)

		default:
			step(t)
		}
	}

	step(t)
}
