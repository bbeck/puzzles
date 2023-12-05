package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	seeds, maps := InputToSeedsAndMaps()

	best := math.MaxInt
	for _, value := range seeds {
		for _, m := range maps {
			value = m.Lookup(value)
		}
		best = aoc.Min(best, value)
	}
	fmt.Println(best)
}

type Range struct {
	Destination int
	Source      int
	Len         int
}

type Map []Range

func (m Map) Lookup(n int) int {
	for _, r := range m {
		if r.Source <= n && n < r.Source+r.Len {
			return r.Destination + (n - r.Source)
		}
	}

	return n
}

func InputToSeedsAndMaps() ([]int, []Map) {
	var seeds []int
	var maps []Map

	var current Map
	for _, line := range aoc.InputToLines(2023, 5) {
		if strings.HasPrefix(line, "seeds:") {
			for _, seed := range strings.Fields(line[6:]) {
				seeds = append(seeds, aoc.ParseInt(seed))
			}
			continue
		}

		if line == "" {
			if current != nil {
				maps = append(maps, current)
				current = nil
			}
			continue
		}

		if strings.Contains(line, "-to-") {
			continue
		}

		nums := strings.Fields(line)
		current = append(current, Range{
			Destination: aoc.ParseInt(nums[0]),
			Source:      aoc.ParseInt(nums[1]),
			Len:         aoc.ParseInt(nums[2]),
		})
	}
	maps = append(maps, current)

	return seeds, maps
}
