package main

import (
	"fmt"
	"math"

	. "github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	sequences := InputToSequences()

	var sum int
	for _, s := range sequences {
		sum += Eval(s)
	}
	fmt.Println(sum)
}

func Eval(ns []int) int {
	var intervals = make([][]Interval, 2)
	var parity int
	var position int

	for _, n := range ns {
		// Determine the range of possible endpoints for this interval.
		var lo, hi = 0, math.MaxInt64
		for _, interval := range intervals[parity] {
			if Contains(interval, position) {
				lo = Max(lo, interval.Start)
				hi = Min(hi, interval.End)
			}
		}

		var next = -1

		if position-n > lo && !Crosses(position, position-n, intervals[parity]) {
			// We're able to move backwards
			next = position - n
		} else {
			// Try to find a move forwards
			for p := position + n; p < hi; p++ {
				if !Crosses(position, p, intervals[parity]) {
					next = p
					break
				}
			}
		}

		if next != -1 {
			intervals[parity] = append(intervals[parity], Interval{Start: Min(position, next), End: Max(position, next)})
			parity = (parity + 1) % 2
			position = next
		}
	}

	return position
}

func Contains(interval Interval, point int) bool {
	return interval.Start <= point && point <= interval.End
}

func Crosses(position, next int, intervals []Interval) bool {
	for _, interval := range intervals {
		if Contains(interval, next) && !Contains(interval, position) {
			return true
		}
	}
	return false
}

type Interval struct {
	Start, End int
}

func InputToSequences() [][]int {
	return in.LinesToS(func(in in.Scanner[[]int]) []int {
		return in.Ints()
	})
}
