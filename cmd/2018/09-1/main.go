package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/puz"
)

func main() {
	numPlayers, numMarbles := InputToParameters()

	var ring puz.Ring[int]
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

	fmt.Println(puz.Max(scores...))
}

func InputToParameters() (int, int) {
	var numPlayers, numMarbles int
	fmt.Sscanf(puz.InputToString(2018, 9), "%d players; last marble is worth %d points", &numPlayers, &numMarbles)
	return numPlayers, numMarbles
}
