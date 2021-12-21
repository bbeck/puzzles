package main

import (
	"fmt"
	"strings"

	"github.com/bbeck/advent-of-code/aoc"
)

var (
	LENGTH = 10
	SCORES = map[int]int{0: 10, 1: 1, 2: 2, 3: 3, 4: 4, 5: 5, 6: 6, 7: 7, 8: 8, 9: 9}
)

func main() {
	positions := InputToStartingPositions()

	wins1, wins2 := CountWins(positions)
	fmt.Println(aoc.MaxInt(wins1, wins2))
}

func CountWins(positions [2]int) (int, int) {
	type State struct {
		positions [2]int
		scores    [2]int
		player    int
	}
	type Wins [2]int

	memo := make(map[State]Wins)

	var helper func(position0, position1, score0, score1, player int) Wins
	helper = func(position0, position1, score0, score1, player int) Wins {
		state := State{
			positions: [2]int{position0, position1},
			scores:    [2]int{score0, score1},
			player:    player,
		}

		if wins, ok := memo[state]; ok {
			return wins
		}

		if state.scores[0] >= 21 {
			return Wins{1, 0}
		}
		if state.scores[1] >= 21 {
			return Wins{0, 1}
		}

		var wins0, wins1 int
		for die1 := 1; die1 <= 3; die1++ {
			for die2 := 1; die2 <= 3; die2++ {
				for die3 := 1; die3 <= 3; die3++ {
					var wins Wins
					if player == 0 {
						position := (position0 + die1 + die2 + die3) % LENGTH
						score := SCORES[position]
						wins = helper(position, position1, score0+score, score1, 1)
					} else {
						position := (position1 + die1 + die2 + die3) % LENGTH
						score := SCORES[position]
						wins = helper(position0, position, score0, score1+score, 0)
					}

					wins0 += wins[0]
					wins1 += wins[1]
				}
			}
		}

		memo[state] = Wins{wins0, wins1}
		return memo[state]
	}

	wins := helper(positions[0], positions[1], 0, 0, 0)
	return wins[0], wins[1]
}

func InputToStartingPositions() [2]int {
	lines := aoc.InputToLines(2021, 21)

	return [2]int{
		aoc.ParseInt(strings.Split(lines[0], ": ")[1]),
		aoc.ParseInt(strings.Split(lines[1], ": ")[1]),
	}
}
