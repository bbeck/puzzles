package main

import (
	"cmp"
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"

	. "github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	grid := in.ToStringGrid2D()

	var start Point2D
	grid.ForEachPoint(func(p Point2D, s string) {
		if s == "S" {
			start = p
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

	// Since the iteration count is so high we know we must have a cycle in the
	// path.  Identify it keeping in mind that costs are strictly increasing and
	// to be able to see the cycle we'll need to normalize them.
	step := func(costs map[Pose2D]int) map[Pose2D]int {
		next := make(map[Pose2D]int, len(costs))
		for p, costSP := range costs {
			for q, costPQ := range TtoT[p] {
				var newCostSQ = costSP + costPQ
				if oldCostSQ, ok := next[q]; !ok || newCostSQ < oldCostSQ {
					next[q] = newCostSQ
				}
			}
		}
		return next
	}

	// In order to help the id function operate more quickly we'll order the
	// keys ahead of time.
	var poses = Keys(TtoT)
	slices.SortFunc(poses, func(a, b Pose2D) int {
		return cmp.Or(
			cmp.Compare(a.Point.Y, b.Point.Y),
			cmp.Compare(a.Point.X, b.Point.X),
			cmp.Compare(a.Heading, b.Heading),
		)
	})

	id := func(costs map[Pose2D]int) string {
		var minimum = Min(Values(costs)...)
		var sb strings.Builder
		for _, pose := range poses {
			if cost, ok := costs[pose]; ok {
				sb.WriteString(strconv.Itoa(cost - minimum))
				sb.WriteString(":")
			} else {
				sb.WriteString("-1:")
			}
		}
		return sb.String()
	}

	prefix, cycle := FindCycleWithIdentity(StoT, step, id)

	// Calculate how much cost is added per cycle.
	after := step(cycle[len(cycle)-1])
	delta := Min(Values(after)...) - Min(Values(cycle[0])...)

	// Calculate how many iterations of the cycle we need to take, and how far
	// into the last cycle we'll be.
	q := (202620262025 - len(prefix)) / len(cycle)
	r := (202620262025 - len(prefix)) % len(cycle)

	// When we're finished with the cycle we need to exit at the pose that has
	// the cheapest cost to the end.
	var best = math.MaxInt64
	for key, cost := range cycle[r] {
		if costToE, ok := TtoE[key]; ok && cost+costToE < best {
			best = cost + costToE
		}
	}

	// Now that we know the full path, compute the total cost.  The total cost is
	// the cost from the start to the end (baked into best) plus the cost of
	// repetitively going through the cycle.
	var cost = best + delta*q
	fmt.Println(cost)
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
