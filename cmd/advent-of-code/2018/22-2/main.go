package main

import (
	"fmt"
	"strings"

	"github.com/bbeck/advent-of-code/lib"
)

func main() {
	cave := InputToCave()

	children := func(s State) []State {
		var children []State

		// We can move to an adjacent cell
		for _, p := range s.location.OrthogonalNeighbors() {
			if p.X < 0 || p.Y < 0 {
				continue
			}

			if terrain := cave.Get(p); ValidTerrainEquipment[terrain][s.equipment] {
				children = append(children, State{p, s.equipment})
			}
		}

		// We can change our equipment
		for equipment := 0; equipment <= 2; equipment++ {
			if terrain := cave.Get(s.location); equipment != s.equipment && ValidTerrainEquipment[terrain][equipment] {
				children = append(children, State{s.location, equipment})
			}
		}

		return children
	}

	goal := func(s State) bool {
		return s.location == cave.target && s.equipment == Torch
	}

	cost := func(from, to State) int {
		if from.location != to.location {
			return 1
		}
		if from.equipment != to.equipment {
			return 7
		}
		return 0
	}

	heuristic := func(s State) int {
		return cave.target.ManhattanDistance(s.location)
	}

	_, total, _ := lib.AStarSearch(State{lib.Origin2D, Torch}, children, goal, cost, heuristic)
	fmt.Println(total)
}

var (
	Rocky  = 0
	Wet    = 1
	Narrow = 2
)

var (
	Nothing      = 0
	Torch        = 1
	ClimbingGear = 2
)

var ValidTerrainEquipment = map[int]map[int]bool{
	Rocky:  {Nothing: false, Torch: true, ClimbingGear: true},
	Wet:    {Nothing: true, Torch: false, ClimbingGear: true},
	Narrow: {Nothing: true, Torch: true, ClimbingGear: false},
}

type State struct {
	location  lib.Point2D
	equipment int
}

type Cave struct {
	geologic map[lib.Point2D]int
	depth    int
	target   lib.Point2D
}

func (c Cave) Geologic(p lib.Point2D) int {
	if _, found := c.geologic[p]; !found {
		var geologic int
		switch {
		case (p.X == 0 && p.Y == 0) || p == c.target:
			geologic = 0
		case p.X == 0:
			geologic = p.Y * 48271
		case p.Y == 0:
			geologic = p.X * 16807
		default:
			geologic = c.Geologic(p.Left()) * c.Geologic(p.Up())
		}

		c.geologic[p] = (geologic + c.depth) % 20183
	}

	return c.geologic[p]
}

func (c Cave) Get(p lib.Point2D) int {
	return c.Geologic(p) % 3
}

func InputToCave() Cave {
	var depth int
	var target lib.Point2D

	for _, line := range lib.InputToLines() {
		k, v, _ := strings.Cut(line, ": ")
		if k == "depth" {
			depth = lib.ParseInt(v)
		} else if k == "target" {
			x, y, _ := strings.Cut(v, ",")
			target = lib.Point2D{X: lib.ParseInt(x), Y: lib.ParseInt(y)}
		}
	}

	return Cave{
		geologic: make(map[lib.Point2D]int),
		depth:    depth,
		target:   target,
	}
}
