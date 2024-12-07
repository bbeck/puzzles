package main

import (
	"fmt"
	"strings"

	. "github.com/bbeck/puzzles/lib"
)

func main() {
	var sum int
	for _, e := range InputToEquations() {
		if IsSolvable(e.Answer, e.Nums) {
			sum += e.Answer
		}
	}
	fmt.Println(sum)
}

func IsSolvable(answer int, nums []int) bool {
	if len(nums) == 1 {
		return nums[0] == answer
	}

	rest := nums[:len(nums)-1]
	last := nums[len(nums)-1]

	// We're able to multiply if the last number divides the answer evenly.
	if answer%last == 0 {
		if IsSolvable(answer/last, rest) {
			return true
		}
	}

	// Otherwise try adding.
	return IsSolvable(answer-last, rest)
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
