package main

import (
	"fmt"
	"strings"

	"github.com/bbeck/puzzles/lib"
)

func main() {
	valves, start := InputToValves()
	partitions := GetPartitions(valves, start)

	var best int
	for _, partition := range partitions {
		state1 := NewState(26, start, uint16(partition.First|(1<<start)&0xFFFF))
		state2 := NewState(26, start, uint16(partition.Second|(1<<start)&0xFFFF))
		best = lib.Max(best, Run(state1, valves)+Run(state2, valves))
	}
	fmt.Println(best)
}

// GetPartitions returns all unique partitions of the valves among the two
// users.  This function ignores the order of the partitions, having player 1
// open valves [1, 2, 3] and player 2 open valves [4, 5] is identical to having
// player 1 open valves [4, 5] and player 2 open valves [1, 2, 3].
func GetPartitions(valves []Valve, start int) []Partition {
	MASK := uint64(0)
	for n := 0; n < len(valves); n++ {
		MASK |= 1 << n
	}
	MASK &= ^(1 << start)

	N := lib.Pow(uint64(2), uint64(len(valves)))

	var partitions lib.Set[Partition]
	for n := uint64(0); n < N; n++ {
		a, b := n&MASK, (^n)&MASK
		a, b = lib.Min(a, b), lib.Max(a, b)
		partitions.Add(Partition{First: a, Second: b})
	}

	return partitions.Entries()
}

type Partition struct {
	First, Second uint64
}

var memo = make(map[State]int)

func Run(state State, valves []Valve) int {
	if best, ok := memo[state]; ok {
		return best
	}

	tm := state.Time()
	valve := valves[state.Current()]
	opened := state.Opened()

	var best int
	for nid, cost := range valve.Neighbors {
		if cost >= tm || state.IsOpened(nid) {
			continue
		}

		next := NewState(tm-cost-1, nid, opened|(1<<nid))
		rate := (tm - cost - 1) * valves[nid].Rate
		best = lib.Max(best, Run(next, valves)+rate)
	}

	memo[state] = best
	return best
}

// State represents the current state packed into a single uint64.  From MSB to
// LSB we have:
//   - Remaining time (5-bits)
//   - Current location (4-bits)
//   - Opened valves (16-bits)
type State uint64

func NewState(tm int, current int, opened uint16) State {
	return State((tm&0x1F)<<20 | (current&0xF)<<16 | int(opened))
}
func (s State) Time() int           { return int((s >> 20) & 0x1F) }
func (s State) Current() int        { return int(s>>16) & 0xF }
func (s State) Opened() uint16      { return uint16(s & 0xFFFF) }
func (s State) IsOpened(n int) bool { return s.Opened()&(1<<n) > 0 }

type Valve struct {
	ID        string
	Rate      int
	Neighbors map[int]int
}

func InputToValves() ([]Valve, int) {
	ids := make([]string, 0)
	rates := make(map[string]int)
	neighbors := make(map[string][]string)

	for _, line := range lib.InputToLines() {
		line = strings.ReplaceAll(line, "Valve ", "")
		line = strings.ReplaceAll(line, "has flow rate=", "")
		line = strings.ReplaceAll(line, "; tunnels lead to valves", "")
		line = strings.ReplaceAll(line, "; tunnel leads to valve", "")
		line = strings.ReplaceAll(line, ",", "")
		fields := strings.Fields(line)

		id := fields[0]
		ids = append(ids, id)
		rates[id] = lib.ParseInt(fields[1])
		neighbors[id] = fields[2:]
	}

	// Floyd-Warshall
	cost := make(map[string]map[string]int)
	for _, i := range ids {
		cost[i] = make(map[string]int)
		for _, j := range ids {
			cost[i][j] = 10000
		}

		cost[i][i] = 0
		for _, n := range neighbors[i] {
			cost[i][n] = 1
		}
	}
	for _, k := range ids {
		for _, i := range ids {
			for _, j := range ids {
				if ck := cost[i][k] + cost[k][j]; ck < cost[i][j] {
					cost[i][j] = ck
				}
			}
		}
	}

	// We're going to keep the valves that have a non-zero rate along with valve
	// AA because it's our starting location.  Assign each of these an index.
	indices := make(map[string]int)

	var N int
	for _, id := range ids {
		if rates[id] > 0 || id == "AA" {
			indices[id] = N
			N++
		}
	}

	// Finally, build the valves.
	valves := make([]Valve, N)
	for _, id := range ids {
		index, ok := indices[id]
		if !ok {
			continue
		}

		ns := make(map[int]int)
		for nid, nindex := range indices {
			if c := cost[id][nid]; c > 0 && rates[nid] > 0 {
				ns[nindex] = c
			}
		}

		valves[index] = Valve{
			ID:        id,
			Rate:      rates[id],
			Neighbors: ns,
		}
	}
	return valves, indices["AA"]
}
