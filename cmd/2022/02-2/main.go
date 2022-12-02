package main

import (
	"fmt"
	"strings"

	"github.com/bbeck/advent-of-code/aoc"
)

var Scores = map[string]map[string]int{
	"A": { // rock
		"X": 3 + 0, // lose (scissors)
		"Y": 1 + 3, // draw (rock)
		"Z": 2 + 6, // win  (paper)
	},
	"B": { // paper
		"X": 1 + 0, // lose (rock)
		"Y": 2 + 3, // draw (paper)
		"Z": 3 + 6, // win  (scissors)
	},
	"C": { // scissors
		"X": 2 + 0, // lose (paper)
		"Y": 3 + 3, // draw (scissors)
		"Z": 1 + 6, // win  (rock)
	},
}

func main() {
	var score int
	for _, line := range aoc.InputToLines(2022, 2) {
		a, b, _ := strings.Cut(line, " ")
		score += Scores[a][b]
	}
	fmt.Println(score)
}
