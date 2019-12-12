package main

import (
	"fmt"
	"log"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	// This part asks us to determine when a repeat state will occur among the
	// moon states and implies that a direct computation won't be computationally
	// feasible.
	//
	// We're going to decompose the problem into three parts, one for each axis.
	// This is possible because the adjustment of a moon's position or velocity
	// along an axis is dependent only on the data from that axis.  We will
	// compute the length of the cycle along each axis for position and velocity
	// pairs.  We will also make sure that the cycle repeats at the first pair so
	// that there's no initial offsetting that needs to happen.
	//
	// Once we've determined the cycle lengths, we can compute the least common
	// multiple of them to determine how many steps it takes to get back to
	// (0, 0, 0).
	cycle := func(pv func(Moon) (int, int)) (int, int) {
		var ps, vs []int
		for _, moon := range InputToMoons(2019, 12) {
			p, v := pv(moon)
			ps = append(ps, p)
			vs = append(vs, v)
		}

		seen := make(map[string]int)

		var step int
		for step = 0; ; step++ {
			// update the velocities
			for i := 0; i < len(ps); i++ {
				for j := i + 1; j < len(ps); j++ {
					if ps[i] < ps[j] {
						vs[i]++
						vs[j]--
					} else if ps[j] < ps[i] {
						vs[i]--
						vs[j]++
					}
				}
			}

			// update the positions
			for i := 0; i < len(ps); i++ {
				ps[i] += vs[i]
			}

			// check if this new state has been seen before
			key := fmt.Sprintf("p:%+v,v:%+v", ps, vs)
			if s, ok := seen[key]; ok {
				return s, step - s
			}
			seen[key] = step
		}
	}

	tailX, lengthX := cycle(func(m Moon) (int, int) { return m.position.X, m.velocity.X })
	tailY, lengthY := cycle(func(m Moon) (int, int) { return m.position.Y, m.velocity.Y })
	tailZ, lengthZ := cycle(func(m Moon) (int, int) { return m.position.Z, m.velocity.Z })
	if tailX != 0 || tailY != 0 || tailZ != 0 {
		log.Fatalf("tail is not zero, x: %d, y: %d, z: %d", tailX, tailY, tailZ)
	}

	fmt.Println(LCM(LCM(lengthX, lengthY), lengthZ))
}

func GCD(a, b int) int {
	if a == 0 {
		return b
	}
	return GCD(b%a, a)
}

func LCM(a, b int) int {
	return (a * b) / GCD(a, b)
}

type Moon struct {
	position Point
	velocity Point
}

func (m Moon) String() string {
	return fmt.Sprintf("pos=<x=%d, y=%d, z=%d>, vel=<x=%d, y=%d, z=%d>",
		m.position.X, m.position.Y, m.position.Z, m.velocity.X, m.velocity.Y, m.velocity.Z)
}

type Point struct {
	X, Y, Z int
}

func InputToMoons(year, day int) []Moon {
	var moons []Moon
	for _, line := range aoc.InputToLines(year, day) {
		var x, y, z int
		if _, err := fmt.Sscanf(line, "<x=%d, y=%d, z=%d>", &x, &y, &z); err != nil {
			log.Fatalf("unable to parse input line: %s\n", line)
		}

		moons = append(moons, Moon{position: Point{x, y, z}})
	}

	return moons
}
