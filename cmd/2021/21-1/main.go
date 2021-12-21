package main

import (
	"fmt"
	"strings"

	"github.com/bbeck/advent-of-code/aoc"
)

var (
	LENGTH = 10
	SIDES  = 1000
	SCORES = map[int]int{0: 10, 1: 1, 2: 2, 3: 3, 4: 4, 5: 5, 6: 6, 7: 7, 8: 8, 9: 9}
)

func main() {
	positions := InputToStartingPositions()
	scores := []int{0, 0}
	die := 0 // 1000 sided die, from 0 to 999

	var turn int
	for turn = 0; scores[0] < 1000 && scores[1] < 1000; turn++ {
		player := turn % 2

		// Determine how far this player is going to travel by rolling the die 3 times.  Keep
		// in mind that the die ranges from 0 to 999 instead of 1 to 1000, so we have to add
		// 3 to compensate.
		distance := 3 + die + (die + 1) + (die + 2)
		die = (die + 3) % SIDES

		// Calculate the player's new position on the track.
		positions[player] = (positions[player] + distance) % LENGTH

		// The player's score is equal to their position on the track, but keep in mind that
		// our track ranges from 0 to 9 whereas the real track ranges from 1 to 10.
		scores[player] += SCORES[positions[player]]
	}

	fmt.Println(aoc.MinInt(scores[0], scores[1]) * turn * 3)
}

func InputToStartingPositions() []int {
	lines := aoc.InputToLines(2021, 21)

	return []int{
		aoc.ParseInt(strings.Split(lines[0], ": ")[1]),
		aoc.ParseInt(strings.Split(lines[1], ": ")[1]),
	}
}
