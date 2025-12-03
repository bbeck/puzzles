package main

import (
	"fmt"

	. "github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	var sum int
	for in.HasNext() {
		var nums []int
		for _, c := range in.Line() {
			nums = append(nums, ParseInt(string(c)))
		}

		var best int
		for i := range nums {
			for j := i + 1; j < len(nums); j++ {
				best = Max(best, 10*nums[i]+nums[j])
			}
		}

		sum += best
	}

	fmt.Println(sum)
}
