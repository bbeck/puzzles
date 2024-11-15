package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/puz"
	"strings"
)

func main() {
	components := InputToComponents()
	fmt.Println(StrongestBridge(components))
}

func StrongestBridge(components []Component) int {
	var helper func(needed int, used puz.BitSet) int
	helper = func(needed int, used puz.BitSet) int {
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

			best = puz.Max(best, c.L+c.R+helper(next, used.Add(i)))
		}

		return best
	}

	return helper(0, 0)
}

type Component struct {
	L, R int
}

func InputToComponents() []Component {
	return puz.InputLinesTo(func(line string) Component {
		lhs, rhs, _ := strings.Cut(line, "/")
		return Component{
			L: puz.ParseInt(lhs),
			R: puz.ParseInt(rhs),
		}
	})
}
