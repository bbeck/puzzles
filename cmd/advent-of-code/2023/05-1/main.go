package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/bbeck/puzzles/lib"
)

func main() {
	seeds, maps := InputToSeedsAndMaps()

	best := math.MaxInt
	for _, value := range seeds {
		for _, m := range maps {
			value = m.Lookup(value)
		}
		best = lib.Min(best, value)
	}
	fmt.Println(best)
}

type Map []Range
type Range struct {
	Destination int
	Source      int
	Length      int
}

func (m Map) Lookup(target int) int {
	for _, r := range m {
		if r.Source <= target && target < r.Source+r.Length {
			return r.Destination + target - r.Source
		}
	}

	return target
}

func InputToSeedsAndMaps() ([]int, []Map) {
	lines := lib.InputToLines()

	var seeds []int
	for _, field := range strings.Fields(lines[0][6:]) {
		seeds = append(seeds, lib.ParseInt(field))
	}

	groups := lib.Split(lines[1:], func(line string) bool {
		return line != "" && !strings.Contains(line, "-to-")
	})

	var maps []Map
	for i, group := range groups {
		maps = append(maps, make([]Range, 0))
		for _, line := range group {
			nums := strings.Fields(line)
			maps[i] = append(maps[i], Range{
				Destination: lib.ParseInt(nums[0]),
				Source:      lib.ParseInt(nums[1]),
				Length:      lib.ParseInt(nums[2]),
			})
		}
	}

	return seeds, maps
}
