package main

import (
	"fmt"
	"github.com/bbeck/puzzles/lib"
)

func main() {
	grid := InputToGrid()

	var tm int
	for tm = 1; ; tm++ {
		grid.ForEachPoint(func(_ lib.Point2D, o *Octopus) {
			o.Increase()
		})

		var count int
		grid.ForEachPoint(func(_ lib.Point2D, o *Octopus) {
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

func InputToGrid() lib.Grid2D[*Octopus] {
	grid := lib.InputToGrid2D(func(x int, y int, s string) *Octopus {
		return &Octopus{Energy: lib.ParseInt(s)}
	})

	grid.ForEachPoint(func(p lib.Point2D, o *Octopus) {
		grid.ForEachNeighborPoint(p, func(np lib.Point2D, no *Octopus) {
			o.Neighbors = append(o.Neighbors, no)
		})
	})

	return grid
}
