package main

import (
	"fmt"
	"math"
	"slices"
	"strings"

	. "github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	grid, ghosts := InputToGridAndGhosts()

	// The approach we're going to take is to iterate through the rounds until we
	// either find a cycle, or a ghost is removed.  If we find a cycle then we'll
	// use the cycle to fast-forward until just before the next ghost is removed.

	step := func() {
		var seen = make(map[Point2D][]int)

		for idx := range ghosts {
			prev := ghosts[idx].Location
			ghosts[idx].Forward(1)
			if grid.GetPoint(ghosts[idx].Location) == "#" {
				ghosts[idx].Location = prev
				ghosts[idx].TurnRight()
				ghosts[idx].Collisions++
			}

			seen[ghosts[idx].Location] = append(seen[ghosts[idx].Location], idx)
		}

		for _, indices := range seen {
			if len(indices) > 1 {
				for _, idx := range indices {
					ghosts[idx].TurnAround()
					ghosts[idx].Collisions++
				}
			}
		}
	}

	prune := func(limit int) bool {
		n := len(ghosts)
		ghosts = slices.DeleteFunc(ghosts, func(ghost Ghost) bool {
			return ghost.Collisions >= limit
		})
		return len(ghosts) < n
	}

	id := func() string {
		var sb strings.Builder
		for _, ghost := range ghosts {
			fmt.Fprintf(&sb, "%d,%d;%s", ghost.Location.X, ghost.Location.Y, ghost.Heading)
		}
		return sb.String()
	}

	type State struct {
		Round      int
		Collisions []int
	}
	state := func(round int) State {
		var collisions = make([]int, len(ghosts))
		for idx := range ghosts {
			collisions[idx] = ghosts[idx].Collisions
		}
		return State{Round: round, Collisions: collisions}
	}

	const limit = 1000000000000
	var round = 0
	for len(ghosts) > 0 {
		// Mapping of key to the last round and state when we stop the key.
		var memo = make(map[string]State)

		// Keep going until we've pruned a ghost
		for {
			step()
			round++

			if prune(limit) {
				break
			}

			// Detect a cycle
			key := id()
			if prev, ok := memo[key]; ok {
				period := round - prev.Round

				// Figure out how many times we iterate around the cycle by seeing how
				// long it takes for a ghost to reach its collision limit.
				var numIterations = math.MaxInt
				for idx := range ghosts {
					var perCycle = ghosts[idx].Collisions - prev.Collisions[idx]
					if perCycle <= 0 {
						continue
					}

					// Largest N such that ghosts[idx].Collisions + N*perCycle < limit, so
					// no ghost reaches the limit during the fast-forward. The final
					// partial cycle is stepped normally so the removal happens on the
					// exact round.
					N := (limit - 1 - ghosts[idx].Collisions) / perCycle
					numIterations = Min(numIterations, N)
				}

				// Fast-forward
				round += numIterations * period
				for idx := range ghosts {
					ghosts[idx].Collisions += numIterations * (ghosts[idx].Collisions - prev.Collisions[idx])
				}

				continue
			}

			// This wasn't a cycle, remember this state for the future.
			memo[key] = state(round)
		}
	}

	fmt.Println(round)
}

type Ghost struct {
	Turtle
	Collisions int
}

func (g *Ghost) TurnAround() {
	g.TurnRight()
	g.TurnRight()
}

func InputToGridAndGhosts() (Grid2D[string], []Ghost) {
	var headings = map[string]Heading{"^": Up, "v": Down, ">": Right, "<": Left}

	var ghosts []Ghost
	grid := in.ToGrid2D(func(x, y int, s string) string {
		if s == "#" || s == "." {
			return s
		}

		ghosts = append(ghosts, Ghost{Location: Point2D{X: x, Y: y}, Heading: headings[s]})
		return "."
	})

	return grid, ghosts
}
