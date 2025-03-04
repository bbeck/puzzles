package main

import (
	"fmt"
	. "github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	particles := InputToParticles()

	for tm := 0; tm < 1000; tm++ {
		positions := make(map[Point3D][]Particle)
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
	pos, vel, acc Point3D
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
	return in.LinesToS(func(in in.Scanner[Particle]) Particle {
		return Particle{
			pos: Point3D{X: in.Int(), Y: in.Int(), Z: in.Int()},
			vel: Point3D{X: in.Int(), Y: in.Int(), Z: in.Int()},
			acc: Point3D{X: in.Int(), Y: in.Int(), Z: in.Int()},
		}
	})
}
