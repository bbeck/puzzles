package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/aoc"
	"sort"
)

var Points = map[int32]int{
	')': 1,
	']': 2,
	'}': 3,
	'>': 4,
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
			score = 5*score + Points[c]
		}

		scores = append(scores, score)
	}

	sort.Ints(scores)
	fmt.Println(scores[len(scores)/2])
}

var Closing = map[int32]int32{
	'(': ')',
	'[': ']',
	'{': '}',
	'<': '>',
}

func Validate(s string) (bool, []int32) {
	stack := aoc.NewStack()
	for _, c := range s {
		if closing, ok := Closing[c]; ok {
			stack.Push(closing)
			continue
		}

		if top := stack.Pop().(int32); c != top {
			return false, nil
		}
	}

	var missing []int32
	for !stack.Empty() {
		missing = append(missing, stack.Pop().(int32))
	}
	return true, missing
}
