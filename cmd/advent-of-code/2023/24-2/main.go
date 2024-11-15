package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/puz"
)

func main() {
	stones := InputToHailstones()
	dx, dy, dz, tm := FindRockVelocities(stones)

	// Now that we know the rock velocity components and intersection time solve
	// the simultaneous equation to determine the rock position components.
	//   rpx + rvx*t = spx + svx*t ==> rpx = spx + (svx-rvx)*t
	x := stones[0].px + (stones[0].vx-dx)*tm
	y := stones[0].py + (stones[0].vy-dy)*tm
	z := stones[0].pz + (stones[0].vz-dz)*tm
	fmt.Println(x + y + z)
}

func FindRockVelocities(stones []Hailstone) (dx, dy, dz, tm int) {
	pxs := []int{stones[0].px, stones[1].px, stones[2].px}
	pys := []int{stones[0].py, stones[1].py, stones[2].py}
	pzs := []int{stones[0].pz, stones[1].pz, stones[2].pz}
	vxs := []int{stones[0].vx, stones[1].vx, stones[2].vx}
	vys := []int{stones[0].vy, stones[1].vy, stones[2].vy}
	vzs := []int{stones[0].vz, stones[1].vz, stones[2].vz}

	// Search to discover the x and y velocity components that the rock should
	// have.  This will enumerate "reasonable" velocity values and identify which
	// cause the first 3 stones to intersect in the same point.
outer:
	for dx = -1000; dx <= 1000; dx++ {
		for dy = -1000; dy <= 1000; dy++ {
			var x1, y1, x2, y2 int
			var ok1, ok2 bool
			x1, y1, tm, ok1 = Intersect(pxs[0], pys[0], vxs[0]-dx, vys[0]-dy, pxs[1], pys[1], vxs[1]-dx, vys[1]-dy)
			x2, y2, _, ok2 = Intersect(pxs[1], pys[1], vxs[1]-dx, vys[1]-dy, pxs[2], pys[2], vxs[2]-dx, vys[2]-dy)

			if ok1 && ok2 && x1 == x2 && y1 == y2 {
				break outer
			}
		}
	}

	// Now that we know the x and y components of the velocity, repeat the search
	// to uncover the x and z velocity components -- but this time only vary the
	// z component since we already know the x component.
	for dz = -1000; dz <= 1000; dz++ {
		var x1, z1, x2, z2 int
		var ok1, ok2 bool
		x1, z1, tm, ok1 = Intersect(pxs[0], pzs[0], vxs[0]-dx, vzs[0]-dz, pxs[1], pzs[1], vxs[1]-dx, vzs[1]-dz)
		x2, z2, _, ok2 = Intersect(pxs[1], pzs[1], vxs[1]-dx, vzs[1]-dz, pxs[2], pzs[2], vxs[2]-dx, vzs[2]-dz)

		if ok1 && ok2 && x1 == x2 && z1 == z2 {
			break
		}
	}

	return dx, dy, dz, tm
}

func Intersect(ax, ay, avx, avy, bx, by, bvx, bvy int) (int, int, int, bool) {
	d := avx*bvy - avy*bvx
	if d == 0 {
		return 0, 0, 0, false
	}

	t := ((bx-ax)*bvy - (by-ay)*bvx) / d
	x := ax + t*avx
	y := ay + t*avy
	return x, y, t, true
}

type Hailstone struct {
	px, py, pz, vx, vy, vz int
}

func InputToHailstones() []Hailstone {
	return puz.InputLinesTo(2023, 24, func(line string) Hailstone {
		var h Hailstone
		fmt.Sscanf(line, "%d, %d, %d @ %d, %d, %d", &h.px, &h.py, &h.pz, &h.vx, &h.vy, &h.vz)
		return h
	})
}
