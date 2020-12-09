package main

import (
	"fmt"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	ns := aoc.InputToInts(2020, 9)
	preamble := 25

	for i := preamble; i < len(ns); i++ {
		if !SumExists(ns[i-preamble:i], ns[i]) {
			fmt.Println(ns[i])
			break
		}
	}
}

func SumExists(ns []int, target int) bool {
	seen := make(map[int]bool)
	for _, n := range ns {
		if seen[target-n] {
			return true
		}
		seen[n] = true
	}
	return false
}
