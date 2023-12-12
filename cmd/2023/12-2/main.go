package main

import (
	"fmt"
	"strings"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	var ways int
	for _, row := range InputToRows() {
		ways += Count(row.Springs, row.Groups)
	}
	fmt.Println(ways)
}

func Count(springs string, groups []int) int {
	S, G := len(springs), len(groups)

	type Key struct{ s, g, run int }
	memo := make(map[Key]int)

	var helper func(int, int, int) int
	helper = func(s, g, run int) int {
		// Base case: We've solved this case before.
		key := Key{s, g, run}
		if ways, found := memo[key]; found {
			return ways
		}

		// Base case: We've reached the end of the input.
		if s == S {
			// If we're on the last group and have seen enough # then we found a way
			if g == G-1 && run == groups[g] {
				memo[key] = 1
				return 1
			}
			// If we're beyond the last group and haven't seen a # then we found a way
			if g == G && run == 0 {
				memo[key] = 1
				return 1
			}
			// We haven't completed a group or have seen extra #, so this is not a way
			memo[key] = 0
			return 0
		}

		var ways int
		if springs[s] == '.' || springs[s] == '?' {
			if run == 0 {
				// We're not in the middle of a group, so we can continue with a .
				ways += helper(s+1, g, 0)
			}

			if g < G && run == groups[g] {
				// We've finished the prior group, so we can continue with a .
				ways += helper(s+1, g+1, 0)
			}
		}

		if springs[s] == '#' || springs[s] == '?' {
			// Continue adding this # to the current group
			ways += helper(s+1, g, run+1)
		}

		memo[key] = ways
		return ways
	}

	return helper(0, 0, 0)
}

type Row struct {
	Springs string
	Groups  []int
}

func InputToRows() []Row {
	return aoc.InputLinesTo(2023, 12, func(line string) Row {
		springs, rhs, _ := strings.Cut(line, " ")

		var groups []int
		for _, g := range strings.Split(rhs, ",") {
			groups = append(groups, aoc.ParseInt(g))
		}

		// Increase everything 5x
		springs = strings.Join(aoc.Repeat([]string{springs}, 5), "?")
		groups = aoc.Repeat(groups, 5)

		return Row{Springs: springs, Groups: groups}
	})
}
