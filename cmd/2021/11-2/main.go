package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	grid := InputToGrid()

	var tm int
	for tm = 1; ; tm++ {
		grid.ForEach(func(_ aoc.Point2D, o *Octopus) {
			o.Increase()
		})

		var count int
		grid.ForEach(func(_ aoc.Point2D, o *Octopus) {
			if o.Reset() {
				count++
			}
		})

		if count == 100 {
			break
		}
	}
	fmt.Println(tm)
}

type Octopus struct {
	Energy    int
	Neighbors []*Octopus
}

func (o *Octopus) Increase() {
	o.Energy++
	if o.Energy == 10 {
		for _, n := range o.Neighbors {
			n.Increase()
		}
	}
}

func (o *Octopus) Reset() bool {
	if o.Energy > 9 {
		o.Energy = 0
		return true
	}
	return false
}

func InputToGrid() aoc.Grid2D[*Octopus] {
	grid := aoc.NewGrid2D[*Octopus](10, 10)
	for y, line := range aoc.InputToLines(2021, 11) {
		for x, c := range line {
			grid.AddXY(x, y, &Octopus{Energy: int(c - '0')})
		}
	}

	grid.ForEach(func(p aoc.Point2D, o *Octopus) {
		grid.ForEachNeighbor(p, func(np aoc.Point2D, no *Octopus) {
			o.Neighbors = append(o.Neighbors, no)
		})
	})

	return grid
}
