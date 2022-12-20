package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/aoc"
	"math"
	"regexp"
)

func main() {
	var sum int
	for _, bp := range InputToBlueprints() {
		sum += bp.ID * Run(bp)
	}
	fmt.Println(sum)
}

func Run(bp Blueprint) int {
	maxCosts := []int{
		aoc.Max(bp.Costs[0][0], bp.Costs[1][0], bp.Costs[2][0], bp.Costs[3][0]),
		aoc.Max(bp.Costs[0][1], bp.Costs[1][1], bp.Costs[2][1], bp.Costs[3][1]),
		aoc.Max(bp.Costs[0][2], bp.Costs[1][2], bp.Costs[2][2], bp.Costs[3][2]),
		math.MaxInt,
	}

	add := [][]int{
		{1, 0, 0, 0},
		{0, 1, 0, 0},
		{0, 0, 1, 0},
		{0, 0, 0, 1},
	}

	compress := func(s State) State {
		// Don't keep around more ore than we can spend for the rest of the time
		// TODO: Make this even tighter since robots are going to keep mining
		s.Ores[0] = aoc.Min(s.Ores[0], s.Time*maxCosts[0])
		s.Ores[1] = aoc.Min(s.Ores[1], s.Time*maxCosts[1])
		s.Ores[2] = aoc.Min(s.Ores[2], s.Time*maxCosts[2])
		return s
	}

	children := func(s State) []State {
		if s.Time <= 0 {
			return nil
		}

		var children []State

		// Build a robot
		for i := 3; i >= 0; i-- {
			if s.Robots[i] < maxCosts[i] &&
				s.Ores[0] >= bp.Costs[i][0] &&
				s.Ores[1] >= bp.Costs[i][1] &&
				s.Ores[2] >= bp.Costs[i][2] &&
				s.Ores[3] >= bp.Costs[i][3] {
				children = append(children, compress(State{
					Time: s.Time - 1,
					Ores: [4]int{
						s.Ores[0] + s.Robots[0] - bp.Costs[i][0],
						s.Ores[1] + s.Robots[1] - bp.Costs[i][1],
						s.Ores[2] + s.Robots[2] - bp.Costs[i][2],
						s.Ores[3] + s.Robots[3] - bp.Costs[i][3],
					},
					Robots: [4]int{
						s.Robots[0] + add[i][0],
						s.Robots[1] + add[i][1],
						s.Robots[2] + add[i][2],
						s.Robots[3] + add[i][3],
					},
				}))

				// If we were able to build a geode robot, then do it and don't consider
				// anything else.
				if i == 3 {
					return children
				}
			}
		}

		// Build nothing
		children = append(children, compress(State{
			Time: s.Time - 1,
			Ores: [4]int{
				s.Ores[0] + s.Robots[0],
				s.Ores[1] + s.Robots[1],
				s.Ores[2] + s.Robots[2],
				s.Ores[3] + s.Robots[3],
			},
			Robots: s.Robots,
		}))

		return children
	}

	start := State{Time: 24, Robots: [4]int{1, 0, 0, 0}}

	var best int
	goal := func(state State) bool {
		if state.Time == 0 {
			best = aoc.Max(best, state.Ores[3])
		}
		return false
	}

	aoc.BreadthFirstSearch(start, children, goal)
	return best
}

type State struct {
	Time   int
	Ores   [4]int
	Robots [4]int
}

func (s State) Copy() State { return s }

type Blueprint struct {
	ID    int
	Costs [][]int
}

func InputToBlueprints() []Blueprint {
	regex := regexp.MustCompile(`\d+`)

	return aoc.InputLinesTo(2022, 19, func(s string) (Blueprint, error) {
		ns := regex.FindAllString(s, -1)
		return Blueprint{
			ID: aoc.ParseInt(ns[0]),
			Costs: [][]int{
				{aoc.ParseInt(ns[1]), 0, 0, 0},
				{aoc.ParseInt(ns[2]), 0, 0, 0},
				{aoc.ParseInt(ns[3]), aoc.ParseInt(ns[4]), 0, 0},
				{aoc.ParseInt(ns[5]), 0, aoc.ParseInt(ns[6]), 0},
			},
		}, nil
	})
}
