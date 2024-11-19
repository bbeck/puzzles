package main

import (
	"fmt"
	"github.com/bbeck/puzzles/lib"
)

var Scores = map[string]int{
	"A X": 1 + 3, // rock     rock     (draw)
	"A Y": 2 + 6, // rock     paper    (win)
	"A Z": 3 + 0, // rock     scissors (lose)
	"B X": 1 + 0, // paper    rock     (lose)
	"B Y": 2 + 3, // paper    paper    (draw)
	"B Z": 3 + 6, // paper    scissors (win)
	"C X": 1 + 6, // scissors rock     (win)
	"C Y": 2 + 0, // scissors paper    (lose)
	"C Z": 3 + 3, // scissors scissors (draw)
}

func main() {
	var score int
	for _, line := range lib.InputToLines() {
		score += Scores[line]
	}
	fmt.Println(score)
}
