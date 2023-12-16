package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	grid := aoc.InputToStringGrid2D(2023, 16)

	var best int
	for x := 0; x < grid.Width; x++ {
		t := aoc.Turtle{Location: aoc.Point2D{X: x, Y: -1}, Heading: aoc.Down}
		best = aoc.Max(best, TryConfiguration(t, grid))

		t = aoc.Turtle{Location: aoc.Point2D{X: x, Y: grid.Height}, Heading: aoc.Up}
		best = aoc.Max(best, TryConfiguration(t, grid))
	}

	for y := 0; y < grid.Height; y++ {
		t := aoc.Turtle{Location: aoc.Point2D{X: -1, Y: y}, Heading: aoc.Right}
		best = aoc.Max(best, TryConfiguration(t, grid))

		t = aoc.Turtle{Location: aoc.Point2D{X: grid.Width, Y: y}, Heading: aoc.Left}
		best = aoc.Max(best, TryConfiguration(t, grid))
	}

	fmt.Println(best)
}

func TryConfiguration(t aoc.Turtle, grid aoc.Grid2D[string]) int {
	energized := aoc.NewGrid2D[string](grid.Width, grid.Height)
	Walk(t, grid, energized)

	var sum int
	energized.ForEach(func(x int, y int, s string) {
		if s == "#" {
			sum++
		}
	})

	return sum
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
