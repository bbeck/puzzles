package main

import (
	"fmt"
	"log"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	ns := aoc.InputToInts(2020, 9)
	target := DetermineInvalid(ns)

	start, end := FindRange(ns, target)
	min := aoc.MinInt(ns[start], ns[start+1:end]...)
	max := aoc.MaxInt(ns[start], ns[start+1:end]...)

	fmt.Println(min + max)
}

func FindRange(ns []int, target int) (int, int) {
	for start := 0; start < len(ns)-1; start++ {
		sum := ns[start]
		for end := start + 1; end < len(ns); end++ {
			sum += ns[end]
			if sum == target {
				return start, end
			}
		}
	}

	log.Fatalf("unable to find a range that summed to: %d", target)
	return 0, 0
}

func DetermineInvalid(ns []int) int {
	preamble := 25
	for i := preamble; i < len(ns); i++ {
		if !SumExists(ns[i-preamble:i], ns[i]) {
			return ns[i]
		}
	}

	log.Fatalf("could not find invalid number in: %v", ns)
	return 0
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
