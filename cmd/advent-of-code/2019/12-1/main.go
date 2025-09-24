package main

import (
	"fmt"

	. "github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	positions := InputToPositions()
	velocities := make([]Point3D, len(positions))

	for tm := 1; tm <= 1000; tm++ {
		UpdateVelocities(positions, velocities)

		for i := range positions {
			positions[i].X += velocities[i].X
			positions[i].Y += velocities[i].Y
			positions[i].Z += velocities[i].Z
		}
	}

	var sum int
	for i := range positions {
		sum += Energy(positions[i], velocities[i])
	}
	fmt.Println(sum)
}

func UpdateVelocities(p, v []Point3D) {
	deltas := func(a, b int) (int, int) {
		switch {
		case a < b:
			return 1, -1
		case a > b:
			return -1, 1
		default:
			return 0, 0
		}
	}

	for i := range p {
		for j := i + 1; j < len(p); j++ {
			dxi, dxj := deltas(p[i].X, p[j].X)
			dyi, dyj := deltas(p[i].Y, p[j].Y)
			dzi, dzj := deltas(p[i].Z, p[j].Z)

			v[i].X += dxi
			v[i].Y += dyi
			v[i].Z += dzi

			v[j].X += dxj
			v[j].Y += dyj
			v[j].Z += dzj
		}
	}
}

func Energy(p, v Point3D) int {
	pot := Abs(p.X) + Abs(p.Y) + Abs(p.Z)
	kin := Abs(v.X) + Abs(v.Y) + Abs(v.Z)
	return pot * kin
}

func InputToPositions() []Point3D {
	return in.LinesToS[Point3D](func(in in.Scanner[Point3D]) Point3D {
		return Point3D{X: in.Int(), Y: in.Int(), Z: in.Int()}
	})
}
