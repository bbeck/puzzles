package main

import (
	"fmt"
	"log"
	"math"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	numPlayers, numMarbles := InputToParameters(2018, 9)

	ring := aoc.NewRing()
	ring.InsertAfter(0)

	scores := make(map[int]int)
	for marble := 1; marble <= numMarbles; marble++ {
		player := ((marble - 1) % numPlayers) + 1

		if marble%23 != 0 {
			// This is a normal insertion, no points scored
			ring.Next()
			ring.InsertAfter(marble)
			continue
		}

		// The player keeps the marble they would have placed, adding it to their
		// score.  Also the marble 7 marbles counter-clockwise from the current
		// marble is removed and also added to their score.
		scores[player] += marble
		ring.PrevN(7)
		scores[player] += ring.Remove().(int)
	}

	// Determine the high score
	highScore := math.MinInt64
	for _, score := range scores {
		if score > highScore {
			highScore = score
		}
	}

	fmt.Printf("high score: %d\n", highScore)
}

func InputToParameters(year, day int) (int, int) {
	line := aoc.InputToString(year, day)
	var numPlayers, numMarbles int
	if _, err := fmt.Sscanf(line, "%d players; last marble is worth %d points", &numPlayers, &numMarbles); err != nil {
		log.Fatalf("unable to parse line: %s", line)
	}

	return numPlayers, numMarbles
}
