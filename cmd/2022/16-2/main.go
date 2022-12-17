package main

import (
	"fmt"
	"strings"

	"github.com/bbeck/advent-of-code/aoc"
)

var valves map[string]Valve

func main() {
	var zero string
	valves = make(map[string]Valve)
	for _, v := range InputToValves() {
		valves[v.ID] = v
		if v.Index == 0 {
			zero = v.ID
		}
	}

	valves[zero] = Valve{
		ID:        zero,
		Index:     valves["AA"].Index,
		Rate:      valves[zero].Rate,
		Neighbors: valves[zero].Neighbors,
	}
	valves["AA"] = Valve{
		ID:        "AA",
		Index:     0,
		Rate:      valves["AA"].Rate,
		Neighbors: valves["AA"].Neighbors,
	}

	for _, v := range valves {
		fmt.Printf("%+v\n", v)
	}

	N := uint64(len(valves))
	MASK := aoc.Pow(uint64(2), N) - 2
	L := MASK + 2

	var best int
	for n := uint64(0); n < L; n++ {
		meTargets := aoc.BitSet(n & MASK)
		elTargets := aoc.BitSet(^uint64(meTargets) & MASK)

		best = aoc.Max(best,
			Run(valves["AA"], 0, meTargets, 26)+
				Run(valves["AA"], 0, elTargets, 26),
		)

		fmt.Println(n, "/", L, best)
		memo = make(map[Key]int)
	}
	fmt.Println("best:", best)
}

type Key struct {
	current string
	opened  aoc.BitSet
	targets aoc.BitSet
	tm      int
}

var memo = make(map[Key]int)

func Run(current Valve, opened aoc.BitSet, targets aoc.BitSet, tm int) int {
	if tm <= 0 || opened.Contains(current.Index) {
		return 0
	}

	key := Key{current.ID, opened, targets, tm}
	if best, ok := memo[key]; ok {
		return best
	}

	var best int
	for nid, cost := range current.Neighbors {
		// We can open the current valve and accrue its rate for the remaining time
		if current.Rate > 0 && targets.Contains(current.Index) {
			val := (tm - 1) * current.Rate
			best = aoc.Max(best, Run(valves[nid], opened.Add(current.Index), targets, tm-cost-1)+val)
		}

		// Or we can leave it shut
		best = aoc.Max(best, Run(valves[nid], opened, targets, tm-cost))
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
