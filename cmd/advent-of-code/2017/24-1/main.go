package main

import (
	"fmt"
	. "github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	components := InputToComponents()
	fmt.Println(StrongestBridge(components))
}

func StrongestBridge(components []Component) int {
	var helper func(needed int, used BitSet) int
	helper = func(needed int, used BitSet) int {
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

			best = Max(best, c.L+c.R+helper(next, used.Add(i)))
		}

		return best
	}

	return helper(0, 0)
}

type Component struct {
	L, R int
}

func InputToComponents() []Component {
	return in.LinesToS(func(in in.Scanner[Component]) Component {
		return Component{L: in.Int(), R: in.Int()}
	})
}
