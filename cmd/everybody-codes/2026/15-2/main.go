package main

import (
	"fmt"
	"math"
	"slices"

	. "github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	grid := in.ToStringGrid2D()

	var start Point2D
	var stops = make(map[Point2D]int)
	grid.ForEachPoint(func(p Point2D, s string) {
		if s == "S" {
			start = p
		}
		if s == "@" {
			stops[p] = len(stops)
		}
	})

	// Start by building a map that holds the shortest path from the start (S)
	// to a stop (T).
	StoT := Costs(grid, Pose2D{Point: start, Heading: Right}, "@")

	// Next build a map that holds the shortest path from each stop to the ones
	// it can reach.  We start with the ones that are reachable from the start,
	// but grow it as we discover more.  Simultaneously we build the mapping from
	// a stop to the end as well.
	TtoT := make(map[Pose2D]map[Pose2D]int)
	TtoE := make(map[Pose2D]int)
	{
		var queue = Keys(StoT)
		for len(queue) > 0 {
			var current Pose2D
			current, queue = queue[0], queue[1:]

			if _, ok := TtoT[current]; ok {
				continue
			}
			TtoT[current] = make(map[Pose2D]int)

			costs := Costs(grid, current, "@", "E")
			for next, cost := range costs {
				if grid.GetPoint(next.Point) == "E" {
					if v, ok := TtoE[current]; !ok || cost < v {
						TtoE[current] = cost
					}
					continue
				}

				TtoT[current][next] = cost

				if _, ok := TtoT[next]; !ok {
					queue = append(queue, next)
				}
			}
		}
	}

	// Finally use dynamic programming to traverse from the start to 2026 stops.
	// We only need to loop 2025 times because we start the DP at one of the
	// stops.
	costs := StoT
	for range 2025 {
		next := make(map[Pose2D]int)
		for p, costSP := range costs {
			for q, costPQ := range TtoT[p] {
				var newCostSQ = costSP + costPQ
				if oldCostSQ, ok := next[q]; !ok || newCostSQ < oldCostSQ {
					next[q] = newCostSQ
				}
			}
		}
		costs = next
	}

	var best = math.MaxInt64
	for p, costSP := range costs {
		if costPE, ok := TtoE[p]; ok {
			best = Min(best, costSP+costPE)
		}
	}
	fmt.Println(best)
}

func Costs(grid Grid2D[string], start Pose2D, include ...string) map[Pose2D]int {
	var costs = make(map[Pose2D]int)

	var queue = []Pose2D{start}
	for len(queue) > 0 {
		var current Pose2D
		current, queue = queue[0], queue[1:]

		// A path ends on a @ so long as it's not our start.
		if current != start && grid.GetPoint(current.Point) == "@" {
			continue
		}

		candidates := []Pose2D{
			current.Forward(),
			current.TurnRight().Forward(),
			current.TurnLeft().Forward(),
		}
		for _, child := range candidates {
			if !grid.InBoundsPoint(child.Point) || grid.GetPoint(child.Point) == "#" {
				continue
			}
			if _, ok := costs[child]; ok {
				continue
			}
			costs[child] = costs[current] + 1
			queue = append(queue, child)
		}
	}

	for p := range costs {
		if v := grid.GetPoint(p.Point); !slices.Contains(include, v) {
			delete(costs, p)
		}
	}
	return costs
}
