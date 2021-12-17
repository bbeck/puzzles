package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	minX, maxX, minY, maxY := InputToTargetArea()

	var solutions []Solution
	for vx := -1000; vx <= 1000; vx++ {
		for vy := -1000; vy <= 1000; vy++ {
			ok, highest := Simulate(vx, vy, minX, maxX, minY, maxY)
			if ok {
				solutions = append(solutions, Solution{vx, vy, highest})
			}
		}
	}

	fmt.Println(len(solutions))
}

func Simulate(vx, vy, minX, maxX, minY, maxY int) (bool, int) {
	hit := func(p Probe) bool {
		return minX <= p.x && p.x <= maxX && minY <= p.y && p.y <= maxY
	}

	p := Probe{
		vx: vx,
		vy: vy,
	}

	var highest int
	for p.y >= minY {
		highest = aoc.MaxInt(highest, p.y)
		if hit(p) {
			return true, highest
		}

		p.Step()
	}

	return false, 0
}

type Solution struct {
	vx, vy, hy int
}

type Probe struct {
	x, y, vx, vy int
}

func (p *Probe) Step() {
	p.x += p.vx
	p.y += p.vy

	if p.vx > 0 {
		p.vx--
	} else if p.vx < 0 {
		p.vx++
	}
	p.vy--
}

func InputToTargetArea() (int, int, int, int) {
	return 25, 67, -260, -200
}
