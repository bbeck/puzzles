package main

import (
	"fmt"
	"github.com/bbeck/puzzles/lib"
)

func main() {
	numPlayers, numMarbles := InputToParameters()

	var ring lib.Ring[int]
	ring.InsertAfter(0)

	scores := make([]int, numPlayers)
	for marble := 1; marble <= numMarbles; marble++ {
		if marble%23 != 0 {
			// Normal play, skip 1 clockwise then insert marble
			ring.Next()
			ring.InsertAfter(marble)
			continue
		}

		// Score this marble and the marble 7 counter-clockwise for the elf
		elf := marble % numPlayers
		ring.PrevN(7)
		scores[elf] += marble + ring.Remove()
	}

	fmt.Println(lib.Max(scores...))
}

func InputToParameters() (int, int) {
	var numPlayers, numMarbles int
	fmt.Sscanf(lib.InputToString(), "%d players; last marble is worth %d points", &numPlayers, &numMarbles)
	return numPlayers, numMarbles
}
