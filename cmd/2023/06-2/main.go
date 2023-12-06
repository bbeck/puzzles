package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	time, distance := InputToTimeAndDistance()

	win := func(hold int) bool { return distance-hold*(time-hold) < 0 }
	first := sort.Search(time, win)

	// To binary search for losses we need to make sure we're "to the right" of
	// the wins.  To do this we'll keep doubling our hold period until we start
	// holding for too long and lose.
	var start int
	for start = first; ; start *= 2 {
		if !win(start) {
			break
		}
	}
	lose := func(hold int) bool { return distance-hold*(time-hold) >= 0 }
	last := sort.Search(start, lose)

	fmt.Println(last - first)
}

func InputToTimeAndDistance() (int, int) {
	var time, distance int
	for _, line := range aoc.InputToLines(2023, 6) {
		var sb strings.Builder
		for _, field := range strings.Fields(line)[1:] {
			sb.WriteString(field)
		}

		num := aoc.ParseInt(sb.String())
		if strings.HasPrefix(line, "Time") {
			time = num
		} else {
			distance = num
		}
	}

	return time, distance
}
