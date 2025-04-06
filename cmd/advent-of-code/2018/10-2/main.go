package main

import (
	"fmt"
	. "github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	points, velocities := InputToPointsAndVelocities()

	var tl, br Point2D
	var tm int
	for tm = 1; ; tm++ {
		for i := 0; i < len(points); i++ {
			points[i] = Point2D{
				X: points[i].X + velocities[i].X,
				Y: points[i].Y + velocities[i].Y,
			}
		}

		tl, br = GetBounds(points)
		if br.Y-tl.Y <= 10 { // Characters are 10 pixels high
			break
		}
	}

	fmt.Println(tm)
}

func InputToPointsAndVelocities() ([]Point2D, []Point2D) {
	var ps, vs []Point2D
	for in.HasNextLine() {
		ps = append(ps, Point2D{X: in.Int(), Y: in.Int()})
		vs = append(vs, Point2D{X: in.Int(), Y: in.Int()})
	}

	return ps, vs
}
