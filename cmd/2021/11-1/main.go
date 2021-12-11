package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	os := InputToOctopus()

	var flashes int
	for step := 1; step <= 100; step++ {
		for _, o := range os {
			o.Add()
		}

		for _, o := range os {
			if o.Flashed {
				flashes++
			}
			o.Reset()
		}
	}

	fmt.Println(flashes)
}

type Octopus struct {
	Energy    int
	Neighbors []*Octopus
	Flashed   bool
}

func (o *Octopus) Add() {
	// If this octopus has already flashed this step then it can't flash again
	if o.Flashed {
		return
	}

	o.Energy++
	if o.Energy > 9 {
		if !o.Flashed {
			o.Flashed = true

			// Propagate to neighbors
			for _, n := range o.Neighbors {
				n.Add()
			}
		}
	}
}

func (o *Octopus) Reset() {
	if o.Flashed {
		o.Energy = 0
		o.Flashed = false
	}
}

func InputToOctopus() map[aoc.Point2D]*Octopus {
	os := make(map[aoc.Point2D]*Octopus)
	for y, line := range aoc.InputToLines(2021, 11) {
		for x, c := range line {
			p := aoc.Point2D{X: x, Y: y}
			os[p] = &Octopus{Energy: aoc.ParseInt(string(c))}
		}
	}

	for p, o := range os {
		o.Neighbors = Neighbors(p, os)
	}

	return os
}

func Neighbors(p aoc.Point2D, os map[aoc.Point2D]*Octopus) []*Octopus {
	candidates := []aoc.Point2D{
		p.Up(),
		p.Up().Right(),
		p.Right(),
		p.Right().Down(),
		p.Down(),
		p.Down().Left(),
		p.Left(),
		p.Left().Up(),
	}

	var neighbors []*Octopus
	for _, c := range candidates {
		if o, ok := os[c]; ok {
			neighbors = append(neighbors, o)
		}
	}

	return neighbors
}
