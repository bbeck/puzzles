package main

import (
	"fmt"
	"strings"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	nums := InputToInts(2020, 15)
	goal := 30000000

	// Mapping of a number to the last time it was spoken, initialized with our
	// input.
	spoken := make(map[int]int)
	for turn, num := range nums {
		spoken[num] = turn
	}

	for turn := len(nums); turn <= goal; turn++ {
		last := nums[turn-1]

		var num int
		if lastTurn, found := spoken[last]; found {
			// This number has been spoken before, so speak the difference between
			// this turn and the turn it was last spoken
			num = turn - lastTurn - 1
		}

		spoken[last] = turn - 1
		nums = append(nums, num)
	}

	fmt.Println(nums[goal])
}

func InputToInts(year, day int) []int {
	ns := []int{-1} // padding to make 1-based
	for _, line := range aoc.InputToLines(year, day) {
		for _, s := range strings.Split(line, ",") {
			ns = append(ns, aoc.ParseInt(s))
		}
	}

	return ns
}
