package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/bbeck/advent-of-code/puz"
)

func main() {
	seeds, maps := InputToSeedsAndMaps()

	best := math.MaxInt
	for _, value := range seeds {
		for _, m := range maps {
			value = m.Lookup(value)
		}
		best = puz.Min(best, value)
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
	lines := puz.InputToLines(2023, 5)

	var seeds []int
	for _, field := range strings.Fields(lines[0][6:]) {
		seeds = append(seeds, puz.ParseInt(field))
	}

	groups := puz.Split(lines[1:], func(line string) bool {
		return line != "" && !strings.Contains(line, "-to-")
	})

	var maps []Map
	for i, group := range groups {
		maps = append(maps, make([]Range, 0))
		for _, line := range group {
			nums := strings.Fields(line)
			maps[i] = append(maps[i], Range{
				Destination: puz.ParseInt(nums[0]),
				Source:      puz.ParseInt(nums[1]),
				Length:      puz.ParseInt(nums[2]),
			})
		}
	}

	return seeds, maps
}
