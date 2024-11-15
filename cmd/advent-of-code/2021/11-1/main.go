package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/puz"
)

func main() {
	grid := InputToGrid()

	var count int
	for tm := 1; tm <= 100; tm++ {
		grid.ForEach(func(_, _ int, o *Octopus) {
			o.Increase()
		})

		grid.ForEach(func(_, _ int, o *Octopus) {
			if o.Reset() {
				count++
			}
		})
	}
	fmt.Println(count)
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

func InputToGrid() puz.Grid2D[*Octopus] {
	grid := puz.InputToGrid2D(func(x int, y int, s string) *Octopus {
		return &Octopus{Energy: puz.ParseInt(s)}
	})

	grid.ForEach(func(x, y int, o *Octopus) {
		grid.ForEachNeighbor(x, y, func(_, _ int, no *Octopus) {
			o.Neighbors = append(o.Neighbors, no)
		})
	})

	return grid
}
