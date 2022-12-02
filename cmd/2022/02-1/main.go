package main

import (
	"fmt"
	"strings"

	"github.com/bbeck/advent-of-code/aoc"
)

var Scores = map[string]map[string]int{
	"A": { // rock
		"X": 1 + 3, // rock     (draw)
		"Y": 2 + 6, // paper    (win)
		"Z": 3 + 0, // scissors (lose)
	},
	"B": { // paper
		"X": 1 + 0, // rock     (lose)
		"Y": 2 + 3, // paper    (draw)
		"Z": 3 + 6, // scissors (win)
	},
	"C": { // scissors
		"X": 1 + 6, // rock     (win)
		"Y": 2 + 0, // paper    (lose)
		"Z": 3 + 3, // scissors (draw)
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
