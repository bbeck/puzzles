package main

import (
	"fmt"
	"strings"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	positions := InputToStartingPositions()
	a, b := CountWins(positions[0], positions[1], 0, 0, 0)
	fmt.Println(aoc.MaxInt(a, b))
}

type Key struct {
	position1, position2, score1, score2, turn int
}
type Pair struct {
	w, l int
}

var seen = make(map[Key]Pair)

func CountWins(position1, position2 int, score1, score2 int, turn int) (int, int) {
	key := Key{position1, position2, score1, score2, turn}
	if p, ok := seen[key]; ok {
		return p.w, p.l
	}

	if score1 >= 21 {
		return 1, 0
	}
	if score2 >= 21 {
		return 0, 1
	}

	var w, l int
	if turn == 0 {
		for d1 := 1; d1 <= 3; d1++ {
			for d2 := 1; d2 <= 3; d2++ {
				for d3 := 1; d3 <= 3; d3++ {
					position := Normalize(position1 + d1 + d2 + d3)
					a, b := CountWins(position, position2, score1+position, score2, 1)
					w += a
					l += b
				}
			}
		}
	} else {
		for d1 := 1; d1 <= 3; d1++ {
			for d2 := 1; d2 <= 3; d2++ {
				for d3 := 1; d3 <= 3; d3++ {
					position := Normalize(position2 + d1 + d2 + d3)
					a, b := CountWins(position1, position, score1, score2+position, 0)
					w += a
					l += b
				}
			}
		}
	}

	seen[key] = Pair{w, l}
	return w, l
}

func Normalize(p int) int {
	for p > 10 {
		p -= 10
	}
	return p
}

func InputToStartingPositions() []int {
	lines := aoc.InputToLines(2021, 21)

	return []int{
		aoc.ParseInt(strings.Split(lines[0], ": ")[1]),
		aoc.ParseInt(strings.Split(lines[1], ": ")[1]),
	}
}
