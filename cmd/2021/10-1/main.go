package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/aoc"
)

var Closing = map[rune]rune{
	'(': ')',
	'[': ']',
	'{': '}',
	'<': '>',
}

var Points = map[rune]int{
	')': 3,
	']': 57,
	'}': 1197,
	'>': 25137,
}

func main() {
	var score int
	for _, line := range aoc.InputToLines(2021, 10) {
		var stack aoc.Stack[rune]
		for _, c := range line {
			if closing, isOpening := Closing[c]; isOpening {
				stack.Push(closing)
			} else if expected := stack.Pop(); c != expected {
				score += Points[c]
				break
			}
		}
	}

	fmt.Println(score)
}
