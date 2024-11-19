package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/lib"
)

var Scores = map[string]int{
	"A X": 3 + 0, // rock     lose (scissors)
	"A Y": 1 + 3, // rock     draw (rock)
	"A Z": 2 + 6, // rock     win  (paper)
	"B X": 1 + 0, // paper    lose (rock)
	"B Y": 2 + 3, // paper    draw (paper)
	"B Z": 3 + 6, // paper    win  (scissors)
	"C X": 2 + 0, // scissors lose (paper)
	"C Y": 3 + 3, // scissors draw (scissors)
	"C Z": 1 + 6, // scissors win  (rock)
}

func main() {
	var score int
	for _, line := range lib.InputToLines() {
		score += Scores[line]
	}
	fmt.Println(score)
}
