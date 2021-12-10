package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/aoc"
	"sort"
)

var SCORES = map[int32]int{
	'(': 1,
	'[': 2,
	'{': 3,
	'<': 4,
}

func main() {
	var scores []int
	for _, line := range aoc.InputToLines(2021, 10) {
		ok, missing := Validate(line)
		if !ok {
			continue
		}

		var score int
		for _, c := range missing {
			score = 5*score + SCORES[c]
		}

		scores = append(scores, score)
	}

	sort.Ints(scores)
	middle := scores[len(scores)/2]
	fmt.Println(middle)
}

func Validate(s string) (bool, []int32) {
	stack := aoc.NewStack()
	for _, c := range s {
		if c == '(' || c == '[' || c == '{' || c == '<' {
			stack.Push(c)
			continue
		}

		top := stack.Pop().(int32)
		if c == ')' && top != '(' {
			return false, nil
		}
		if c == ']' && top != '[' {
			return false, nil
		}
		if c == '}' && top != '{' {
			return false, nil
		}
		if c == '>' && top != '<' {
			return false, nil
		}
	}

	var missing []int32
	for !stack.Empty() {
		c := stack.Pop().(int32)
		missing = append(missing, c)
	}
	return true, missing
}
