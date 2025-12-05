package main

import (
	"fmt"

	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	intervals := InputToIntervals()
	ns := in.Ints()

	var count int
	for _, n := range ns {
		for _, interval := range intervals {
			if interval.Start <= n && n < interval.End {
				count++
				break
			}
		}
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
