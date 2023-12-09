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
	// We shouldn't ever build more robot of a given type than the most we can
	// spend in a round.  That's because even if we spent everything in the round
	// the very next round we'll already be at max.  More than that is just a
	// waste.
	maxNeeded := [4]int{
		aoc.Max(bp.Costs[0][0], bp.Costs[1][0], bp.Costs[2][0], bp.Costs[3][0]),
		aoc.Max(bp.Costs[0][1], bp.Costs[1][1], bp.Costs[2][1], bp.Costs[3][1]),
		aoc.Max(bp.Costs[0][2], bp.Costs[1][2], bp.Costs[2][2], bp.Costs[3][2]),
		math.MaxInt,
	}

	add := [4][4]int{
		{1, 0, 0, 0},
		{0, 1, 0, 0},
		{0, 0, 1, 0},
		{0, 0, 0, 1},
	}

	var helper func(State) int
	helper = func(s State) int {
		if s.Time == 0 {
			return s.Ores[3]
		}

		// If we do nothing then we end up with as much geode as we already have
		// plus whatever our geode robots can mine in the remaining time.
		best := s.Ores[3] + s.Robots[3]*s.Time

		// Alternatively we can build a robot.  We may not yet have the resources
		// needed so will have to wait for them to be mined.
		for b := 0; b < 4; b++ {
			// Don't consider building a robot if we already have the maximum.
			if s.Robots[b] >= maxNeeded[b] {
				continue
			}

			// Figure out how long we have to wait to get enough of each ore to build
			// this robot.
			var wait int
			for ore := 0; ore < 4; ore++ {
				needed := bp.Costs[b][ore] - s.Ores[ore]
				if needed <= 0 {
					continue
				}

				if s.Robots[ore] == 0 {
					// We can't make this robot because we can't make an ingredient.
					wait = math.MaxInt
					break
				}

				// Integer division might round down, so take the integer ceiling.
				dt := needed / s.Robots[ore]
				if needed%s.Robots[ore] != 0 {
					dt++
				}
				wait = aoc.Max(wait, dt)
			}

			// Make sure we don't have to wait longer than we have.
			remaining := s.Time - wait - 1
			if remaining <= 0 {
				continue
			}

			// Jump to the time when we can build the robot.
			best = aoc.Max(best, helper(State{
				Time: remaining,
				Ores: [4]int{
					// For the non-geode ores, don't allow them to accumulate beyond the
					// amount we can spend in the remaining time.
					aoc.Min(s.Ores[0]+s.Robots[0]*(wait+1)-bp.Costs[b][0], remaining*maxNeeded[0]),
					aoc.Min(s.Ores[1]+s.Robots[1]*(wait+1)-bp.Costs[b][1], remaining*maxNeeded[1]),
					aoc.Min(s.Ores[2]+s.Robots[2]*(wait+1)-bp.Costs[b][2], remaining*maxNeeded[2]),
					s.Ores[3] + s.Robots[3]*(wait+1) - bp.Costs[b][3],
				},
				Robots: [4]int{
					s.Robots[0] + add[b][0],
					s.Robots[1] + add[b][1],
					s.Robots[2] + add[b][2],
					s.Robots[3] + add[b][3],
				},
			}))
		}

		return best
	}

	return helper(State{Time: 24, Robots: [4]int{1, 0, 0, 0}})
}

type State struct {
	Time   int
	Ores   [4]int
	Robots [4]int
}

type Blueprint struct {
	ID    int
	Costs [][]int
}

func InputToBlueprints() []Blueprint {
	regex := regexp.MustCompile(`\d+`)

	return aoc.InputLinesTo(2022, 19, func(s string) Blueprint {
		ns := regex.FindAllString(s, -1)
		return Blueprint{
			ID: aoc.ParseInt(ns[0]),
			Costs: [][]int{
				{aoc.ParseInt(ns[1]), 0, 0, 0},
				{aoc.ParseInt(ns[2]), 0, 0, 0},
				{aoc.ParseInt(ns[3]), aoc.ParseInt(ns[4]), 0, 0},
				{aoc.ParseInt(ns[5]), 0, aoc.ParseInt(ns[6]), 0},
			},
		}
	})
}
