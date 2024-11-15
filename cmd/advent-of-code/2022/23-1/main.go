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

	for round := 1; round <= 10; round++ {
		targets, counts := DetermineElfTargets(grove, moves, round)

		for f, t := range targets {
			if f != t && counts[t] == 1 {
				grove.Remove(f)
				grove.Add(t)
			}
		}
	}

	var count int
	tl, br := puz.GetBounds(grove.Entries())
	for y := tl.Y; y <= br.Y; y++ {
		for x := tl.X; x <= br.X; x++ {
			if !grove.Contains(puz.Point2D{X: x, Y: y}) {
				count++
			}
		}
	}
	fmt.Println(count)
}

type Move func(puz.Point2D) (puz.Point2D, bool)

func DetermineElfTargets(grove puz.Set[puz.Point2D], moves []Move, round int) (map[puz.Point2D]puz.Point2D, map[puz.Point2D]int) {
	targets := make(map[puz.Point2D]puz.Point2D)
	counts := make(map[puz.Point2D]int)

	for p := range grove {
		var numNeighbors int
		for _, n := range p.Neighbors() {
			if grove.Contains(n) {
				numNeighbors++
			}
		}

		var moved bool
		if numNeighbors > 0 {
			for i := 0; i < len(moves); i++ {
				if next, ok := moves[(round+i)%len(moves)](p); ok {
					targets[p] = next
					counts[next]++
					moved = true
					break
				}
			}
		}
		if !moved {
			targets[p] = p
			counts[p]++
		}
	}

	return targets, counts
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
