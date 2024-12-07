package main

import (
	"fmt"
	"strings"

	. "github.com/bbeck/puzzles/lib"
)

func main() {
	var sum int
	for _, e := range InputToEquations() {
		if IsSolvable(e) {
			sum += e.Answer
		}
	}
	fmt.Println(sum)
}

func IsSolvable(e Equation) bool {
	N := len(e.Nums) - 1

	var found bool
	EnumerateChoices(2, N, func(choice []int) bool {
		value := e.Nums[0]
		for b := 0; b < N; b++ {
			switch choice[b] {
			case 0: // +
				value += e.Nums[b+1]
			case 1: // *
				value *= e.Nums[b+1]
			}
		}

		found = value == e.Answer
		return found
	})

	return found
}

type Equation struct {
	Answer int
	Nums   []int
}

func InputToEquations() []Equation {
	return InputLinesTo(func(s string) Equation {
		lhs, rhs, _ := strings.Cut(s, ":")

		var nums []int
		for _, n := range strings.Fields(rhs) {
			nums = append(nums, ParseInt(n))
		}

		return Equation{
			Answer: ParseInt(lhs),
			Nums:   nums,
		}
	})
}
