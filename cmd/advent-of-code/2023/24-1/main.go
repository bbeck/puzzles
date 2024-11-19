package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/lib"
)

const (
	Min = 200000000000000
	Max = 400000000000000
)

func main() {
	stones := InputToHailstones()

	var count int
	for i := 0; i < len(stones); i++ {
		for j := i + 1; j < len(stones); j++ {
			a, b := stones[i], stones[j]
			x, y, tma, tmb, ok := Intersect(a.px, a.py, a.vx, a.vy, b.px, b.py, b.vx, b.vy)
			if ok && tma > 0 && tmb > 0 && Min < x && x < Max && Min < y && y < Max {
				count++
			}
		}
	}
	fmt.Println(count)
}

func Intersect(apx, apy, avx, avy, bpx, bpy, bvx, bvy int) (float64, float64, float64, float64, bool) {
	d := avx*bvy - avy*bvx
	if d == 0 {
		return 0, 0, 0, 0, false
	}

	t := float64((bpx-apx)*bvy-(bpy-apy)*bvx) / float64(d)
	u := float64((bpx-apx)*avy-(bpy-apy)*avx) / float64(d)
	x := float64(apx) + t*float64(avx)
	y := float64(apy) + t*float64(avy)
	return x, y, t, u, true
}

type Hailstone struct {
	px, py, pz, vx, vy, vz int
}

func InputToHailstones() []Hailstone {
	return lib.InputLinesTo(func(line string) Hailstone {
		var h Hailstone
		fmt.Sscanf(line, "%d, %d, %d @ %d, %d, %d", &h.px, &h.py, &h.pz, &h.vx, &h.vy, &h.vz)
		return h
	})
}
