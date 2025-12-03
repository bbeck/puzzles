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

		sum += Value(nums, 0, 12, make(map[string]int))
	}

	fmt.Println(sum)
}

func Value(nums []int, idx int, remaining uint, memo map[string]int) int {
	// We're done
	if remaining == 0 {
		return 0
	}

	// We've run out of numbers to use
	if idx == len(nums) {
		return -1
	}

	key := fmt.Sprintf("%d|%d", idx, remaining)
	if value, ok := memo[key]; ok && value >= 0 {
		return value
	}

	// First consider not using the ith digit
	value := Value(nums, idx+1, remaining, memo)

	// Next consider using the ith digit (if we can)
	if v := Value(nums, idx+1, remaining-1, memo); v >= 0 {
		value = Max(
			value,
			nums[idx]*Pow(10, remaining-1)+v,
		)
	}

	memo[key] = value
	return value
}
