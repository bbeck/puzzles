package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/puz"
	"strings"
)

func main() {
	components := InputToComponents()
	fmt.Println(LongestBridge(components))
}

func LongestBridge(components []Component) int {
	var helper func(needed int, used puz.BitSet) (int, int)
	helper = func(needed int, used puz.BitSet) (int, int) {
		var bestLength, bestStrength int
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

			length, strength := helper(next, used.Add(i))
			bestLength, bestStrength = Best(length+1, strength+c.L+c.R, bestLength, bestStrength)
		}

		return bestLength, bestStrength
	}

	_, strength := helper(0, 0)
	return strength
}

func Best(length1, strength1, length2, strength2 int) (int, int) {
	if length1 > length2 || (length1 == length2 && strength1 > strength2) {
		return length1, strength1
	}
	return length2, strength2
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
