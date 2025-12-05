package main

import (
	"fmt"
	"sort"

	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	intervals := InputToIntervals()
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].Start < intervals[j].Start
	})

	merged := []Interval{intervals[0]}
	for i, m := 0, 0; i < len(intervals); i++ {
		_, mEnd := merged[m].Start, merged[m].End
		iStart, iEnd := intervals[i].Start, intervals[i].End

		switch {
		case mEnd <= iStart:
			// This interval is completely after the previous one
			merged = append(merged, intervals[i])
			m++

		case mEnd < iEnd:
			// This interval overlaps with the previous one
			merged[m].End = iEnd
		}
	}

	var count int
	for _, interval := range merged {
		count += interval.End - interval.Start
	}
	fmt.Println(count)
}

func InputToIntervals() []Interval {
	s := in.ChunkS()

	var intervals []Interval
	for s.HasNext() {
		var start, end int
		s.Scanf("%d-%d", &start, &end)

		intervals = append(intervals, Interval{Start: start, End: end + 1})
	}

	return intervals
}

type Interval struct {
	Start, End int
}
