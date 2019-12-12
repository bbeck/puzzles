package main

import (
	"fmt"
	"log"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	moons := InputToMoons(2019, 12)
	for step := 0; step < 1000; step++ {
		// First update the velocity of every moon applying gravity
		for i := 0; i < len(moons); i++ {
			for j := i + 1; j < len(moons); j++ {
				if moons[i].position.X < moons[j].position.X {
					moons[i].velocity.X++
					moons[j].velocity.X--
				} else if moons[i].position.X > moons[j].position.X {
					moons[i].velocity.X--
					moons[j].velocity.X++
				}
				if moons[i].position.Y < moons[j].position.Y {
					moons[i].velocity.Y++
					moons[j].velocity.Y--
				} else if moons[i].position.Y > moons[j].position.Y {
					moons[i].velocity.Y--
					moons[j].velocity.Y++
				}
				if moons[i].position.Z < moons[j].position.Z {
					moons[i].velocity.Z++
					moons[j].velocity.Z--
				} else if moons[i].position.Z > moons[j].position.Z {
					moons[i].velocity.Z--
					moons[j].velocity.Z++
				}
			}
		}

		// Next update the position of every moon applying velocity
		for i := 0; i < len(moons); i++ {
			moons[i].position.X += moons[i].velocity.X
			moons[i].position.Y += moons[i].velocity.Y
			moons[i].position.Z += moons[i].velocity.Z
		}
	}

	fmt.Printf("total energy: %d\n", Energy(moons))
}

func Energy(moons []Moon) int {
	abs := func(n int) int {
		if n < 0 {
			return -n
		}
		return n
	}

	var energy int
	for _, moon := range moons {
		potential := abs(moon.position.X) + abs(moon.position.Y) + abs(moon.position.Z)
		kinetic := abs(moon.velocity.X) + abs(moon.velocity.Y) + abs(moon.velocity.Z)

		energy += potential * kinetic
	}

	return energy
}

type Moon struct {
	position Point
	velocity Point
}

func (m Moon) String() string {
	return fmt.Sprintf("pos=<x=%2d, y=%3d, z=%2d>, vel=<x=%2d, y=%2d, z=%2d>",
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
