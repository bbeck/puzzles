package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/lib"
)

func main() {
	var black lib.Set[HexPoint]
	for _, steps := range InputToSteps() {
		var current HexPoint
		for _, step := range steps {
			current = current.Step(step)
		}

		if !black.Add(current) {
			black.Remove(current)
		}
	}

	fmt.Println(len(black))
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

// https://www.redblobgames.com/grids/hexagons/
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
