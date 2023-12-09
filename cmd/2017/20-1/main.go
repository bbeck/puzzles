package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/aoc"
	"math"
)

func main() {
	// Determine the position of each particle in the distant future and see
	// which is closest.
	const T = 1000
	closest := -1
	distance := math.MaxInt

	for id, particle := range InputToParticles() {
		d := aoc.Origin3D.ManhattanDistance(particle.Position(T))
		if d < distance {
			distance = d
			closest = id
		}
	}

	fmt.Println(closest)
}

type Particle struct {
	p, v, a aoc.Point3D
}

func (p Particle) Position(t int) aoc.Point3D {
	// The position of a particle undergoing constant acceleration at a time t in
	// the future is given by: p(t) = p0 + v0*t + 1/2*a*t^2
	t2 := t * t / 2
	return aoc.Point3D{
		X: p.p.X + p.v.X*t + p.a.X*t2,
		Y: p.p.X + p.v.Y*t + p.a.Y*t2,
		Z: p.p.X + p.v.Z*t + p.a.Z*t2,
	}
}

func InputToParticles() []Particle {
	return aoc.InputLinesTo(2017, 20, func(line string) Particle {
		var particle Particle
		fmt.Sscanf(line, "p=<%d,%d,%d>, v=<%d,%d,%d>, a=<%d,%d,%d>",
			&particle.p.X, &particle.p.Y, &particle.p.Z,
			&particle.v.X, &particle.v.Y, &particle.v.Z,
			&particle.a.X, &particle.a.Y, &particle.a.Z,
		)
		return particle
	})
}
