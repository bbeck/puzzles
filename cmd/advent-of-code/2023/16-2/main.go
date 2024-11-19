package main

import (
	"fmt"
	"github.com/bbeck/puzzles/lib"
)

func main() {
	grid := lib.InputToStringGrid2D()

	var best int
	for x := 0; x < grid.Width; x++ {
		t := lib.Turtle{Location: lib.Point2D{X: x, Y: -1}, Heading: lib.Down}
		best = lib.Max(best, TryConfiguration(t, grid))

		t = lib.Turtle{Location: lib.Point2D{X: x, Y: grid.Height}, Heading: lib.Up}
		best = lib.Max(best, TryConfiguration(t, grid))
	}

	for y := 0; y < grid.Height; y++ {
		t := lib.Turtle{Location: lib.Point2D{X: -1, Y: y}, Heading: lib.Right}
		best = lib.Max(best, TryConfiguration(t, grid))

		t = lib.Turtle{Location: lib.Point2D{X: grid.Width, Y: y}, Heading: lib.Left}
		best = lib.Max(best, TryConfiguration(t, grid))
	}

	fmt.Println(best)
}

func TryConfiguration(t lib.Turtle, grid lib.Grid2D[string]) int {
	energized := lib.NewGrid2D[string](grid.Width, grid.Height)
	Walk(t, grid, energized)

	var sum int
	energized.ForEach(func(x int, y int, s string) {
		if s == "#" {
			sum++
		}
	})

	return sum
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
