package main

import (
	"fmt"
	"strings"

	"github.com/bbeck/advent-of-code/puz"
)

func main() {
	valves, start := InputToValves()

	best := Run(NewState(30, start, 1<<start), valves)
	fmt.Println(best)
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
		best = puz.Max(best, Run(next, valves)+rate)
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

	for _, line := range puz.InputToLines(2022, 16) {
		line = strings.ReplaceAll(line, "Valve ", "")
		line = strings.ReplaceAll(line, "has flow rate=", "")
		line = strings.ReplaceAll(line, "; tunnels lead to valves", "")
		line = strings.ReplaceAll(line, "; tunnel leads to valve", "")
		line = strings.ReplaceAll(line, ",", "")
		fields := strings.Fields(line)

		id := fields[0]
		ids = append(ids, id)
		rates[id] = puz.ParseInt(fields[1])
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
