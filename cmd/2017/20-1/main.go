package main

import (
	"fmt"
	"log"
	"math"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	particles := InputToParticles(2017, 20)

	for n := 0; n < 1000; n++ {
		var closestID int
		var closestDistance = math.MaxInt64

		for id, particle := range particles {
			particle.vx += particle.ax
			particle.vy += particle.ay
			particle.vz += particle.az
			particle.x += particle.vx
			particle.y += particle.vy
			particle.z += particle.vz

			// Keep track of whether or not this is the closest particle to the origin
			distance := particle.Distance()
			if distance < closestDistance {
				closestID = id
				closestDistance = distance
			}
		}

		fmt.Printf("n:%d, id:%d, particle:%+v\n", n, closestID, particles[closestID])
	}
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
