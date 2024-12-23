package main

import (
	"fmt"
	"github.com/bbeck/puzzles/lib"
)

func main() {
	positions := InputToPositions()
	velocities := make([]lib.Point3D, len(positions))

	for tm := 1; tm <= 1000; tm++ {
		UpdateVelocities(positions, velocities)

		for i := 0; i < len(positions); i++ {
			positions[i].X += velocities[i].X
			positions[i].Y += velocities[i].Y
			positions[i].Z += velocities[i].Z
		}
	}

	var sum int
	for i := 0; i < len(positions); i++ {
		sum += Energy(positions[i], velocities[i])
	}
	fmt.Println(sum)
}

func UpdateVelocities(p, v []lib.Point3D) {
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

	for i := 0; i < len(p); i++ {
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

func Energy(p, v lib.Point3D) int {
	pot := lib.Abs(p.X) + lib.Abs(p.Y) + lib.Abs(p.Z)
	kin := lib.Abs(v.X) + lib.Abs(v.Y) + lib.Abs(v.Z)
	return pot * kin
}

func InputToPositions() []lib.Point3D {
	return lib.InputLinesTo(func(line string) lib.Point3D {
		var p lib.Point3D
		fmt.Sscanf(line, "<x=%d, y=%d, z=%d>", &p.X, &p.Y, &p.Z)
		return p
	})
}
