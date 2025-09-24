package main

import (
	"fmt"

	. "github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	var xs, ys, zs []int
	for _, p := range InputToPositions() {
		xs = append(xs, p.X)
		ys = append(ys, p.Y)
		zs = append(zs, p.Z)
	}

	lenY := CycleLength(xs)
	lenZ := CycleLength(ys)
	lenX := CycleLength(zs)
	fmt.Println(LCM(lenX, lenY, lenZ))
}

func CycleLength(ps []int) int {
	vs := make([]int, len(ps)) // Velocities always start at 0

	// Since there's a cycle, we should eventually return to our original
	// positions and velocities.
	ops := append([]int{}, ps...)
	ovs := append([]int{}, vs...)

	var n int
	for {
		Update(ps, vs)
		n++

		if Equal(ps, ops) && Equal(vs, ovs) {
			break
		}
	}
	return n
}

func Update(p, v []int) {
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
			di, dj := deltas(p[i], p[j])
			v[i] += di
			v[j] += dj
		}
	}

	for i := range p {
		p[i] += v[i]
	}
}

func Equal[T comparable](a, b []T) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func InputToPositions() []Point3D {
	return in.LinesToS[Point3D](func(in in.Scanner[Point3D]) Point3D {
		return Point3D{X: in.Int(), Y: in.Int(), Z: in.Int()}
	})
}
