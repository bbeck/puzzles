package main

import (
	"fmt"
	"strings"

	"github.com/bbeck/advent-of-code/puz"
)

var (
	Scores = map[int]int{0: 10, 1: 1, 2: 2, 3: 3, 4: 4, 5: 5, 6: 6, 7: 7, 8: 8, 9: 9}
)

func main() {
	positions := InputToStartingPositions()
	fmt.Println(puz.Max(CountWins(positions)))
}

func CountWins(positions []int) (int, int) {
	memo := make(map[[5]int][2]int)

	var helper func(turn, position0, position1, score0, score1 int) (int, int)
	helper = func(turn, position0, position1, score0, score1 int) (int, int) {
		if score0 >= 21 {
			return 1, 0
		}
		if score1 >= 21 {
			return 0, 1
		}

		key := [5]int{turn, position0, position1, score0, score1}
		if value, found := memo[key]; found {
			return value[0], value[1]
		}

		var wins0, wins1 int
		for die1 := 1; die1 <= 3; die1++ {
			for die2 := 1; die2 <= 3; die2++ {
				for die3 := 1; die3 <= 3; die3++ {
					var w0, w1 int
					if turn == 0 {
						position := (position0 + die1 + die2 + die3) % 10
						score := score0 + Scores[position]
						w0, w1 = helper(1, position, position1, score, score1)
					} else {
						position := (position1 + die1 + die2 + die3) % 10
						score := score1 + Scores[position]
						w0, w1 = helper(0, position0, position, score0, score)
					}

					wins0 += w0
					wins1 += w1
				}
			}
		}

		memo[key] = [2]int{wins0, wins1}
		return wins0, wins1
	}

	return helper(0, positions[0], positions[1], 0, 0)
}

func InputToStartingPositions() []int {
	return puz.InputLinesTo(func(line string) int {
		_, rhs, _ := strings.Cut(line, ": ")
		return puz.ParseInt(rhs)
	})
}
