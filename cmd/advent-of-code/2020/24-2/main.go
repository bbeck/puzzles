package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/lib"
)

func main() {
	var floor lib.Set[HexPoint]
	for _, steps := range InputToSteps() {
		var current HexPoint
		for _, step := range steps {
			current = current.Step(step)
		}

		if !floor.Add(current) {
			floor.Remove(current)
		}
	}

	for n := 0; n < 100; n++ {
		floor = Next(floor)
	}

	fmt.Println(len(floor))
}

func Next(floor lib.Set[HexPoint]) lib.Set[HexPoint] {
	// We want to consider any black tile and any tiles adjacent to it.
	var tiles lib.Set[HexPoint]
	for p := range floor {
		tiles.Add(p.Neighbors()...)
	}

	var next lib.Set[HexPoint]
	for p := range tiles {
		var count int
		for _, q := range p.Neighbors() {
			if floor.Contains(q) {
				count++
			}
		}

		if floor.Contains(p) && 0 < count && count <= 2 {
			next.Add(p)
		} else if !floor.Contains(p) && count == 2 {
			next.Add(p)
		}
	}
	return next
}

func InputToSteps() [][]string {
	return lib.InputLinesTo(func(line string) []string {
		var steps []string
		for len(line) > 0 {
			if line[0] == 'n' || line[0] == 's' {
				steps = append(steps, line[:2])
				line = line[2:]
			} else {
				steps = append(steps, line[:1])
				line = line[1:]
			}
		}

		return steps
	})
}

// Deltas discussed in the axial coordinates section from here:
//
//	https://www.redblobgames.com/grids/hexagons/#neighbors-axial
var Deltas = map[string]HexPoint{
	"ne": {1, -1},
	"e":  {1, 0},
	"se": {0, 1},
	"sw": {-1, 1},
	"w":  {-1, 0},
	"nw": {0, -1},
}

type HexPoint struct{ Q, R int }

func (p HexPoint) Step(dir string) HexPoint {
	delta := Deltas[dir]
	return HexPoint{Q: p.Q + delta.Q, R: p.R + delta.R}
}

func (p HexPoint) Neighbors() []HexPoint {
	neighbors := make([]HexPoint, 0, 6)
	for _, delta := range Deltas {
		neighbors = append(neighbors, HexPoint{Q: p.Q + delta.Q, R: p.R + delta.R})
	}
	return neighbors
}
