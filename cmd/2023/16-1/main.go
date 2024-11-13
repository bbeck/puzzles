package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/puz"
)

func main() {
	turtle := puz.Turtle{Location: puz.Origin2D.Left(), Heading: puz.Right}
	grid := puz.InputToStringGrid2D(2023, 16)
	energized := puz.NewGrid2D[string](grid.Width, grid.Height)

	Walk(turtle, grid, energized)

	var sum int
	energized.ForEach(func(x int, y int, s string) {
		if s == "#" {
			sum++
		}
	})
	fmt.Println(sum)
}

func Walk(t puz.Turtle, grid, energized puz.Grid2D[string]) {
	var seen puz.Set[puz.Turtle]

	var step func(t puz.Turtle)
	step = func(t puz.Turtle) {
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
		case (cell == "|" && (heading == puz.Right || heading == puz.Left)) ||
			(cell == "-" && (heading == puz.Up || heading == puz.Down)):
			t.TurnLeft()
			step(t)
			t.TurnRight()
			t.TurnRight()
			step(t)

		case cell == "\\":
			if heading == puz.Up || heading == puz.Down {
				t.TurnLeft()
			} else {
				t.TurnRight()
			}
			step(t)

		case cell == "/":
			if heading == puz.Up || heading == puz.Down {
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
