package main

import (
	"fmt"
	"strings"

	"github.com/bbeck/advent-of-code/puz"
)

var (
	Scores = map[int]int{1: 1, 2: 2, 3: 3, 4: 4, 5: 5, 6: 6, 7: 7, 8: 8, 9: 9, 0: 10}
)

func main() {
	positions := InputToStartingPositions()

	var die Die
	var rolls int
	var scores [2]int

	for turn := 0; scores[0] < 1000 && scores[1] < 1000; turn = (turn + 1) % 2 {
		sum := die.Roll() + die.Roll() + die.Roll()
		rolls += 3

		positions[turn] = (positions[turn] + sum) % 10
		scores[turn] += Scores[positions[turn]]
	}

	fmt.Println(rolls * puz.Min(scores[:]...))
}

type Die int

func (d *Die) Roll() int {
	*d += 1
	if *d > 100 {
		*d = 1
	}
	return int(*d)
}

func InputToStartingPositions() []int {
	return puz.InputLinesTo(func(line string) int {
		_, rhs, _ := strings.Cut(line, ": ")
		return puz.ParseInt(rhs)
	})
}
