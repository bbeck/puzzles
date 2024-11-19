package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/lib"
)

func main() {
	grove := InputToGrove()

	moves := []Move{
		func(p lib.Point2D) (lib.Point2D, bool) {
			return p.Up(), !grove.Contains(p.Up().Left()) && !grove.Contains(p.Up()) && !grove.Contains(p.Up().Right())
		},
		func(p lib.Point2D) (lib.Point2D, bool) {
			return p.Down(), !grove.Contains(p.Down().Left()) && !grove.Contains(p.Down()) && !grove.Contains(p.Down().Right())
		},
		func(p lib.Point2D) (lib.Point2D, bool) {
			return p.Left(), !grove.Contains(p.Left()) && !grove.Contains(p.Left().Up()) && !grove.Contains(p.Left().Down())
		},
		func(p lib.Point2D) (lib.Point2D, bool) {
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
	tl, br := lib.GetBounds(grove.Entries())
	for y := tl.Y; y <= br.Y; y++ {
		for x := tl.X; x <= br.X; x++ {
			if !grove.Contains(lib.Point2D{X: x, Y: y}) {
				count++
			}
		}
	}
	fmt.Println(count)
}

type Move func(lib.Point2D) (lib.Point2D, bool)

func DetermineElfTargets(grove lib.Set[lib.Point2D], moves []Move, round int) (map[lib.Point2D]lib.Point2D, map[lib.Point2D]int) {
	targets := make(map[lib.Point2D]lib.Point2D)
	counts := make(map[lib.Point2D]int)

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

func InputToGrove() lib.Set[lib.Point2D] {
	var grove lib.Set[lib.Point2D]
	for y, line := range lib.InputToLines() {
		for x, c := range line {
			if c == '#' {
				grove.Add(lib.Point2D{X: x, Y: y})
			}
		}
	}
	return grove
}
