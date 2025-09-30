package main

import (
	"fmt"

	. "github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/in"
)

var turns, wheels = InputToTurnsAndWheels()
var W = len(wheels)

func main() {
	minimum, maximum := ScoreRange(make([]int, W), 256)
	fmt.Println(maximum, minimum)
}

var memo = make(map[string][]int)

func ScoreRange(indices []int, n int) (int, int) {
	if n == 0 {
		return 0, 0
	}

	id := fmt.Sprintf("%v | %d", indices, n)
	if value, found := memo[id]; found {
		return value[0], value[1]
	}

	var iPush, iStay, iPull []int
	for i := range W {
		iPush = append(iPush, (indices[i]+turns[i]+1)%len(wheels[i]))
		iStay = append(iStay, (indices[i]+turns[i]+0)%len(wheels[i]))
		iPull = append(iPull, (indices[i]+turns[i]-1)%len(wheels[i]))
	}

	minPush, maxPush := ScoreRange(iPush, n-1)
	minStay, maxStay := ScoreRange(iStay, n-1)
	minPull, maxPull := ScoreRange(iPull, n-1)

	sPush := Score(iPush)
	sStay := Score(iStay)
	sPull := Score(iPull)

	minimum := min(minPush+sPush, minStay+sStay, minPull+sPull)
	maximum := max(maxPush+sPush, maxStay+sStay, maxPull+sPull)

	memo[id] = []int{minimum, maximum}
	return minimum, maximum
}

func Score(indices []int) int {
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
