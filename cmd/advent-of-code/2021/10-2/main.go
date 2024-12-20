package main

import (
	"fmt"
	"github.com/bbeck/puzzles/lib"
	"sort"
)

var Closing = map[rune]rune{
	'(': ')',
	'[': ']',
	'{': '}',
	'<': '>',
}

var Points = map[rune]int{
	')': 1,
	']': 2,
	'}': 3,
	'>': 4,
}

func main() {
	var scores []int
	for _, line := range lib.InputToLines() {
		stack, isCorrupted := Check(line)
		if isCorrupted {
			continue
		}

		var score int
		for !stack.Empty() {
			score = 5*score + Points[stack.Pop()]
		}

		scores = append(scores, score)
	}
	sort.Ints(scores)
	fmt.Println(scores[len(scores)/2])
}

func Check(line string) (lib.Stack[rune], bool) {
	var stack lib.Stack[rune]
	for _, c := range line {
		if closing, isOpening := Closing[c]; isOpening {
			stack.Push(closing)
		} else if expected := stack.Pop(); c != expected {
			return stack, true
		}
	}

	return stack, false
}
