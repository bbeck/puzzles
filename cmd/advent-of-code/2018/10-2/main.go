package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/lib"
)

func main() {
	points, velocities := InputToPoints(), InputToVelocities()

	var tl, br lib.Point2D
	var tm int
	for tm = 1; ; tm++ {
		for i := 0; i < len(points); i++ {
			points[i] = lib.Point2D{
				X: points[i].X + velocities[i].X,
				Y: points[i].Y + velocities[i].Y,
			}
		}

		tl, br = lib.GetBounds(points)
		if br.Y-tl.Y <= 10 { // Characters are 10 pixels high
			break
		}
	}

	fmt.Println(tm)
}

func InputToPoints() []lib.Point2D {
	var unused int
	return lib.InputLinesTo(func(line string) lib.Point2D {
		var p lib.Point2D
		fmt.Sscanf(line, "position=<%d, %d> velocity=<%d, %d>", &p.X, &p.Y, &unused, &unused)
		return p
	})
}

func InputToVelocities() []lib.Point2D {
	var unused int
	return lib.InputLinesTo(func(line string) lib.Point2D {
		var p lib.Point2D
		fmt.Sscanf(line, "position=<%d, %d> velocity=<%d, %d>", &unused, &unused, &p.X, &p.Y)
		return p
	})
}
