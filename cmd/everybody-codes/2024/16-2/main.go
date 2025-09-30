package main

import (
	"fmt"
	"strings"

	. "github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/in"
)

const N = 2024_2024_2024

func main() {
	turns, wheels := InputToTurnsAndWheels()

	next := func(indices []int) []int {
		var next []int
		for i := range indices {
			next = append(next, (indices[i]+turns[i])%len(wheels[i]))
		}
		return next
	}

	id := func(indices []int) string {
		var values []string
		for _, index := range indices {
			values = append(values, fmt.Sprintf("%d", index))
		}

		return strings.Join(values, " ")
	}

	score := func(indices []int) int {
		var fc FrequencyCounter[string]
		for i := range indices {
			s := wheels[i][indices[i]]
			for j, ch := range s {
				if j%2 == 0 {
					fc.Add(string(ch))
				}
			}
		}

		var score int
		for _, entry := range fc.Entries() {
			if entry.Count >= 3 {
				score += entry.Count - 2
			}
		}
		return score
	}

	prefix, cycle := FindCycleWithIdentity(make([]int, len(wheels)), next, id)
	iters := (N - len(prefix)) / len(cycle)
	suffix := cycle[:(N - len(prefix) - iters*len(cycle))]

	var pScore int
	for i := range prefix {
		pScore += score(prefix[i])
	}

	var cScore int
	for i := range cycle {
		cScore += iters * score(cycle[i])
	}

	var sScore int
	for i := range suffix {
		sScore += score(suffix[i])
	}

	fmt.Println(pScore + cScore + sScore)
}

func InputToTurnsAndWheels() ([]int, [][]string) {
	c1 := in.ChunkS()
	turns := c1.Ints()

	var wheels = make([][]string, len(turns))

	c2 := in.ChunkS()
	for _, line := range c2.Lines() {
		for i := range wheels {
			if len(line) >= 4*i+3 {
				symbol := line[4*i : 4*i+3]
				if symbol != "   " {
					wheels[i] = append(wheels[i], symbol)
				}
			}
		}
	}

	return turns, wheels
}
