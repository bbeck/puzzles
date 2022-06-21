package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/aoc"
	"strings"
)

func main() {
	components := InputToComponents()
	fmt.Println(StrongestBridge(components))
}

func StrongestBridge(components []Component) int {
	var helper func(needed int, used aoc.BitSet) int
	helper = func(needed int, used aoc.BitSet) int {
		var best int
		for i, c := range components {
			var next int
			if used.Contains(i) {
				continue
			} else if c.L == needed {
				next = c.R
			} else if c.R == needed {
				next = c.L
			} else {
				continue
			}

			best = aoc.MaxInt(best, c.L+c.R+helper(next, used.Add(i)))
		}

		return best
	}

	return helper(0, 0)
}

type Component struct {
	L, R int
}

func InputToComponents() []Component {
	return aoc.InputLinesTo(2017, 24, func(line string) (Component, error) {
		lhs, rhs, _ := strings.Cut(line, "/")
		return Component{
			L: aoc.ParseInt(lhs),
			R: aoc.ParseInt(rhs),
		}, nil
	})
}
