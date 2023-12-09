package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	points, velocities := InputToPoints(), InputToVelocities()

	var tl, br aoc.Point2D
	for {
		for i := 0; i < len(points); i++ {
			points[i] = aoc.Point2D{
				X: points[i].X + velocities[i].X,
				Y: points[i].Y + velocities[i].Y,
			}
		}

		tl, br = aoc.GetBounds(points)
		if br.Y-tl.Y <= 10 { // Characters are 10 pixels high
			break
		}
	}

	grid := aoc.NewGrid2D[bool](br.X+1, br.Y+1)
	for _, p := range points {
		grid.SetPoint(p, true)
	}

	for y := tl.Y; y <= br.Y; y++ {
		for x := tl.X; x <= br.X; x++ {
			if grid.Get(x, y) {
				fmt.Print("█")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func InputToPoints() []aoc.Point2D {
	var unused int
	return aoc.InputLinesTo(2018, 10, func(line string) aoc.Point2D {
		var p aoc.Point2D
		fmt.Sscanf(line, "position=<%d, %d> velocity=<%d, %d>", &p.X, &p.Y, &unused, &unused)
		return p
	})
}

func InputToVelocities() []aoc.Point2D {
	var unused int
	return aoc.InputLinesTo(2018, 10, func(line string) aoc.Point2D {
		var p aoc.Point2D
		fmt.Sscanf(line, "position=<%d, %d> velocity=<%d, %d>", &unused, &unused, &p.X, &p.Y)
		return p
	})
}
