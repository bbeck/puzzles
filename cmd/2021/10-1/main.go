package main

import (
	"fmt"

	"github.com/bbeck/advent-of-code/aoc"
)

var Points = map[int32]int{
	')': 3,
	']': 57,
	'}': 1197,
	'>': 25137,
}

func main() {
	var score int
	for _, line := range aoc.InputToLines(2021, 10) {
		if ok, c := Validate(line); !ok {
			score += Points[c]
		}
	}

	fmt.Println(score)
}

var Closing = map[int32]int32{
	'(': ')',
	'[': ']',
	'{': '}',
	'<': '>',
}

func Validate(s string) (bool, int32) {
	stack := aoc.NewStack()
	for _, c := range s {
		if closing, ok := Closing[c]; ok {
			stack.Push(closing)
			continue
		}

		if top := stack.Pop().(int32); c != top {
			return false, c
		}
	}

	return true, 0
}
