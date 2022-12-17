package main

import (
	"fmt"
	"strings"

	"github.com/bbeck/advent-of-code/aoc"
)

var valves map[string]Valve

func main() {
	valves = make(map[string]Valve)
	for _, v := range InputToValves() {
		valves[v.ID] = v
	}

	best := Run(valves["AA"], aoc.BitSetFrom(valves["AA"].Index), 30)
	fmt.Println(best)
}

type Key struct {
	current int
	opened  aoc.BitSet
	tm      int
}

var memo = make(map[Key]int)

func Run(current Valve, opened aoc.BitSet, tm int) int {
	if tm <= 0 {
		return 0
	}

	key := Key{current.Index, opened, tm}
	if best, ok := memo[key]; ok {
		return best
	}

	val := (tm - 1) * current.Rate

	var best int
	for nid, cost := range current.Neighbors {
		if !opened.Contains(current.Index) {
			// We can open the current valve and accrue its rate for the remaining time
			best = aoc.Max(best, Run(valves[nid], opened.Add(current.Index), tm-cost-1)+val)
		}

		// Or we can leave it shut
		best = aoc.Max(best, Run(valves[nid], opened, tm-cost))
	}

	memo[key] = best
	return best
}

type Valve struct {
	ID        string
	Index     int
	Rate      int
	Neighbors map[string]int
}

func InputToValves() []Valve {
	var ids []string
	var rates []int
	var neighbors [][]string
	indices := make(map[string]int)

	for i, line := range aoc.InputToLines(2022, 16) {
		line = strings.ReplaceAll(line, "Valve ", "")
		line = strings.ReplaceAll(line, "has flow rate=", "")
		line = strings.ReplaceAll(line, "; tunnels lead to valves", "")
		line = strings.ReplaceAll(line, "; tunnel leads to valve", "")
		line = strings.ReplaceAll(line, ",", "")
		fields := strings.Fields(line)

		ids = append(ids, fields[0])
		rates = append(rates, aoc.ParseInt(fields[1]))
		neighbors = append(neighbors, fields[2:])
		indices[fields[0]] = i
	}

	// Floyd-Warshall
	N := len(ids)
	cost := aoc.Make2D[int](N, N)
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			cost[i][j] = 10000
		}
	}
	for i := 0; i < N; i++ {
		cost[i][i] = 0
		for _, n := range neighbors[i] {
			cost[i][indices[n]] = 1
		}
	}

	for k := 0; k < N; k++ {
		for i := 0; i < N; i++ {
			for j := 0; j < N; j++ {
				if cost[i][k]+cost[k][j] < cost[i][j] {
					cost[i][j] = cost[i][k] + cost[k][j]
				}
			}
		}
	}

	var valves []Valve
	for i := 0; i < N; i++ {
		if rates[i] == 0 && ids[i] != "AA" {
			continue
		}

		ns := make(map[string]int)
		for j := 0; j < N; j++ {
			if i == j {
				continue
			}
			if rates[j] > 0 {
				ns[ids[j]] = cost[i][indices[ids[j]]]
			}
		}

		valves = append(valves, Valve{
			ID:        ids[i],
			Index:     len(valves),
			Rate:      rates[i],
			Neighbors: ns,
		})
	}
	return valves
}
