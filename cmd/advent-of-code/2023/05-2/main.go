package main

import (
	"fmt"
	"strings"

	"github.com/bbeck/puzzles/lib"
)

func main() {
	seeds, maps := InputToSeedsAndMaps()

	for _, m := range maps {
		var next []Interval
		for _, s := range seeds {
			next = append(next, m.FindOutput(s)...)
		}

		seeds = next
	}

	var starts []int
	for _, s := range seeds {
		starts = append(starts, s.Start)
	}

	fmt.Println(lib.Min(starts...))
}

type Interval struct {
	Start int
	End   int
}

func (i Interval) Intersect(j Interval) Interval {
	return Interval{
		Start: lib.Max(i.Start, j.Start),
		End:   lib.Min(i.End, j.End),
	}
}

type Map []Range
type Range struct {
	Source      Interval
	Destination Interval
}

func (m Map) FindOutput(interval Interval) []Interval {
	var out []Interval
	for _, r := range m {
		i := interval.Intersect(r.Source)
		if i.Start >= i.End {
			continue
		}

		start := r.Destination.Start + i.Start - r.Source.Start
		end := start + i.End - i.Start
		out = append(out, Interval{Start: start, End: end})
	}

	if len(out) == 0 {
		return []Interval{interval}
	}

	return out
}

func InputToSeedsAndMaps() ([]Interval, []Map) {
	lines := lib.InputToLines()

	var seeds []Interval
	fields := strings.Fields(lines[0][6:])
	for i := 0; i < len(fields); i += 2 {
		start := lib.ParseInt(fields[i])
		length := lib.ParseInt(fields[i+1])
		seeds = append(seeds, Interval{start, start + length - 1})
	}

	groups := lib.Split(lines[1:], func(line string) bool {
		return line != "" && !strings.Contains(line, "-to-")
	})

	var maps []Map
	for i, group := range groups {
		maps = append(maps, make([]Range, 0))
		for _, line := range group {
			nums := strings.Fields(line)
			destination := lib.ParseInt(nums[0])
			source := lib.ParseInt(nums[1])
			length := lib.ParseInt(nums[2])

			maps[i] = append(maps[i], Range{
				Source:      Interval{source, source + length - 1},
				Destination: Interval{destination, destination + length - 1},
			})
		}
	}

	return seeds, maps
}
