package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/puz"
)

func main() {
	particles := InputToParticles()

	for tm := 0; tm < 1000; tm++ {
		positions := make(map[puz.Point3D][]Particle)
		for i := 0; i < len(particles); i++ {
			particles[i].Step()
			positions[particles[i].pos] = append(positions[particles[i].pos], particles[i])
		}

		var next []Particle
		for _, ps := range positions {
			if len(ps) > 1 {
				continue
			}
			next = append(next, ps...)
		}

		particles = next
	}

	fmt.Println(len(particles))
}

type Particle struct {
	pos, vel, acc puz.Point3D
}

func (p *Particle) Step() {
	p.vel.X += p.acc.X
	p.vel.Y += p.acc.Y
	p.vel.Z += p.acc.Z
	p.pos.X += p.vel.X
	p.pos.Y += p.vel.Y
	p.pos.Z += p.vel.Z
}

func InputToParticles() []Particle {
	return puz.InputLinesTo(func(line string) Particle {
		var particle Particle
		fmt.Sscanf(line, "p=<%d,%d,%d>, v=<%d,%d,%d>, a=<%d,%d,%d>",
			&particle.pos.X, &particle.pos.Y, &particle.pos.Z,
			&particle.vel.X, &particle.vel.Y, &particle.vel.Z,
			&particle.acc.X, &particle.acc.Y, &particle.acc.Z,
		)
		return particle
	})
}
