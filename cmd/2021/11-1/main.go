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
			o.Flashed = false
		}
	}

	fmt.Println(flashes)
}

type Octopus struct {
	Location  aoc.Point2D
	Energy    int
	Neighbors []*Octopus
	Flashed   bool
}

func (o *Octopus) Add() {
	if o.Flashed {
		return
	}

	o.Energy++
	if o.Energy >= 10 {
		o.Energy = 0

		if !o.Flashed {
			o.Flashed = true
			for _, n := range o.Neighbors {
				n.Add()
			}
		}
	}
}

func InputToOctopus() []*Octopus {
	os := make(map[aoc.Point2D]*Octopus)
	for y, line := range aoc.InputToLines(2021, 11) {
		for x, c := range line {
			os[aoc.Point2D{X: x, Y: y}] = &Octopus{Location: aoc.Point2D{x, y}, Energy: aoc.ParseInt(string(c))}
		}
	}

	var octs []*Octopus
	for p, o := range os {
		var neighbors []*Octopus
		if _, ok := os[p.Up()]; ok {
			neighbors = append(neighbors, os[p.Up()])
		}
		if _, ok := os[p.Down()]; ok {
			neighbors = append(neighbors, os[p.Down()])
		}
		if _, ok := os[p.Left()]; ok {
			neighbors = append(neighbors, os[p.Left()])
		}
		if _, ok := os[p.Right()]; ok {
			neighbors = append(neighbors, os[p.Right()])
		}
		if _, ok := os[p.Up().Right()]; ok {
			neighbors = append(neighbors, os[p.Up().Right()])
		}
		if _, ok := os[p.Up().Left()]; ok {
			neighbors = append(neighbors, os[p.Up().Left()])
		}
		if _, ok := os[p.Down().Right()]; ok {
			neighbors = append(neighbors, os[p.Down().Right()])
		}
		if _, ok := os[p.Down().Left()]; ok {
			neighbors = append(neighbors, os[p.Down().Left()])
		}
		o.Neighbors = neighbors

		octs = append(octs, o)
	}

	return octs
}
