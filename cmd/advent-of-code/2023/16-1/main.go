package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/lib"
)

func main() {
	turtle := lib.Turtle{Location: lib.Origin2D.Left(), Heading: lib.Right}
	grid := lib.InputToStringGrid2D()
	energized := lib.NewGrid2D[string](grid.Width, grid.Height)

	Walk(turtle, grid, energized)

	var sum int
	energized.ForEach(func(x int, y int, s string) {
		if s == "#" {
			sum++
		}
	})
	fmt.Println(sum)
}

func Walk(t lib.Turtle, grid, energized lib.Grid2D[string]) {
	var seen lib.Set[lib.Turtle]

	var step func(t lib.Turtle)
	step = func(t lib.Turtle) {
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
		case (cell == "|" && (heading == lib.Right || heading == lib.Left)) ||
			(cell == "-" && (heading == lib.Up || heading == lib.Down)):
			t.TurnLeft()
			step(t)
			t.TurnRight()
			t.TurnRight()
			step(t)

		case cell == "\\":
			if heading == lib.Up || heading == lib.Down {
				t.TurnLeft()
			} else {
				t.TurnRight()
			}
			step(t)

		case cell == "/":
			if heading == lib.Up || heading == lib.Down {
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
