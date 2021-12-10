package main

import (
	"fmt"

	"github.com/bbeck/advent-of-code/aoc"
)

var SCORES = map[int32]int{
	')': 3,
	']': 57,
	'}': 1197,
	'>': 25137,
}

func main() {
	var score int
	for _, line := range aoc.InputToLines(2021, 10) {
		ok, c := Validate(line)
		if !ok {
			score += SCORES[c]
		}
	}

	fmt.Println(score)
}

func Validate(s string) (bool, int32) {
	stack := aoc.NewStack()
	for _, c := range s {
		if c == '(' || c == '[' || c == '{' || c == '<' {
			stack.Push(c)
			continue
		}

		top := stack.Pop().(int32)
		if c == ')' && top != '(' {
			return false, c
		}
		if c == ']' && top != '[' {
			return false, c
		}
		if c == '}' && top != '{' {
			return false, c
		}
		if c == '>' && top != '<' {
			return false, c
		}
	}

	return true, 0
}
