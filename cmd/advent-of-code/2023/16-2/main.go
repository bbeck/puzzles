package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/puz"
)

func main() {
	grid := puz.InputToStringGrid2D()

	var best int
	for x := 0; x < grid.Width; x++ {
		t := puz.Turtle{Location: puz.Point2D{X: x, Y: -1}, Heading: puz.Down}
		best = puz.Max(best, TryConfiguration(t, grid))

		t = puz.Turtle{Location: puz.Point2D{X: x, Y: grid.Height}, Heading: puz.Up}
		best = puz.Max(best, TryConfiguration(t, grid))
	}

	for y := 0; y < grid.Height; y++ {
		t := puz.Turtle{Location: puz.Point2D{X: -1, Y: y}, Heading: puz.Right}
		best = puz.Max(best, TryConfiguration(t, grid))

		t = puz.Turtle{Location: puz.Point2D{X: grid.Width, Y: y}, Heading: puz.Left}
		best = puz.Max(best, TryConfiguration(t, grid))
	}

	fmt.Println(best)
}

func TryConfiguration(t puz.Turtle, grid puz.Grid2D[string]) int {
	energized := puz.NewGrid2D[string](grid.Width, grid.Height)
	Walk(t, grid, energized)

	var sum int
	energized.ForEach(func(x int, y int, s string) {
		if s == "#" {
			sum++
		}
	})

	return sum
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
