package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/puz"
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
	fmt.Println(puz.LCM(lenX, lenY, lenZ))
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

	for i := 0; i < len(p); i++ {
		for j := i + 1; j < len(p); j++ {
			di, dj := deltas(p[i], p[j])
			v[i] += di
			v[j] += dj
		}
	}

	for i := 0; i < len(p); i++ {
		p[i] += v[i]
	}
}

func Equal[T comparable](a, b []T) bool {
	if len(a) != len(b) {
		return false
	}

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func InputToPositions() []puz.Point3D {
	return puz.InputLinesTo(func(line string) puz.Point3D {
		var p puz.Point3D
		fmt.Sscanf(line, "<x=%d, y=%d, z=%d>", &p.X, &p.Y, &p.Z)
		return p
	})
}
