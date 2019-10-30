package main

import (
	"fmt"
	"log"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	particles := InputToParticles(2017, 20)

	for n := 0; n < 1000; n++ {
		for _, particle := range particles {
			if particle == nil {
				continue
			}

			particle.vx += particle.ax
			particle.vy += particle.ay
			particle.vz += particle.az
			particle.x += particle.vx
			particle.y += particle.vy
			particle.z += particle.vz
		}

		// Now that all of the particles have been moved, check for collisions and
		// remove any particles that have collided.
		seen := make(map[Point3D]int)
		for id, particle := range particles {
			if particle == nil {
				continue
			}

			p := Point3D{x: particle.x, y: particle.y, z: particle.z}
			if oid, ok := seen[p]; ok {
				particles[oid] = nil
				particles[id] = nil
			}

			seen[p] = id
		}

		var count int
		for _, particle := range particles {
			if particle != nil {
				count++
			}
		}

		fmt.Printf("n: %d, num left: %d\n", n, count)
	}
}

type Point3D struct {
	x, y, z int
}

type Particle struct {
	x, y, z    int
	vx, vy, vz int
	ax, ay, az int
}

func (p *Particle) Distance() int {
	abs := func(n int) int {
		if n < 0 {
			n = -n
		}
		return n
	}

	return abs(p.x) + abs(p.y) + abs(p.z)
}

func InputToParticles(year, day int) []*Particle {
	var particles []*Particle
	for _, line := range aoc.InputToLines(year, day) {
		var x, y, z int
		var vx, vy, vz int
		var ax, ay, az int

		_, err := fmt.Sscanf(line, "p=<%d,%d,%d>, v=<%d,%d,%d>, a=<%d,%d,%d>", &x, &y, &z, &vx, &vy, &vz, &ax, &ay, &az)
		if err != nil {
			log.Fatalf("unable to parse line: %s", line)
		}

		particles = append(particles, &Particle{x, y, z, vx, vy, vz, ax, ay, az})
	}

	return particles
}
