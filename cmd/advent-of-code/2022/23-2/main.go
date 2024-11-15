package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/puz"
)

func main() {
	grove := InputToGrove()

	moves := []Move{
		func(p puz.Point2D) (puz.Point2D, bool) {
			return p.Up(), !grove.Contains(p.Up().Left()) && !grove.Contains(p.Up()) && !grove.Contains(p.Up().Right())
		},
		func(p puz.Point2D) (puz.Point2D, bool) {
			return p.Down(), !grove.Contains(p.Down().Left()) && !grove.Contains(p.Down()) && !grove.Contains(p.Down().Right())
		},
		func(p puz.Point2D) (puz.Point2D, bool) {
			return p.Left(), !grove.Contains(p.Left()) && !grove.Contains(p.Left().Up()) && !grove.Contains(p.Left().Down())
		},
		func(p puz.Point2D) (puz.Point2D, bool) {
			return p.Right(), !grove.Contains(p.Right()) && !grove.Contains(p.Right().Up()) && !grove.Contains(p.Right().Down())
		},
	}

	var round int
	for round = 1; ; round++ {
		targets := DetermineElfTargets(grove, moves, round)

		var moved int
		for f, t := range targets {
			if f == t {
				continue
			}

			if grove.Add(t) {
				grove.Remove(f)
				moved++
				continue
			}

			// If there was an elf already in our new position that means they moved
			// earlier this round.  We know based on the rules that they came from the
			// opposite direction as this elf.  Move them back.
			grove.Remove(t)
			moved--

			dx, dy := t.X-f.X, t.Y-f.Y
			grove.Add(puz.Point2D{X: t.X + dx, Y: t.Y + dy})
		}

		if moved == 0 {
			break
		}
	}
	fmt.Println(round)
}

type Move func(puz.Point2D) (puz.Point2D, bool)

func DetermineElfTargets(grove puz.Set[puz.Point2D], moves []Move, round int) map[puz.Point2D]puz.Point2D {
	targets := make(map[puz.Point2D]puz.Point2D)

	for p := range grove {
		targets[p] = p

		var hasNeighbors bool
		for _, n := range p.Neighbors() {
			if grove.Contains(n) {
				hasNeighbors = true
				break
			}
		}

		if !hasNeighbors {
			continue
		}

		for i := 0; i < len(moves); i++ {
			if next, ok := moves[(round+i-1)%len(moves)](p); ok {
				targets[p] = next
				break
			}
		}
	}

	return targets
}

func InputToGrove() puz.Set[puz.Point2D] {
	var grove puz.Set[puz.Point2D]
	for y, line := range puz.InputToLines() {
		for x, c := range line {
			if c == '#' {
				grove.Add(puz.Point2D{X: x, Y: y})
			}
		}
	}
	return grove
}
