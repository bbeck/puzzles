package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/puz"
)

func main() {
	calories := puz.InputLinesTo(func(line string) int {
		if line == "" {
			return 0
		}
		return puz.ParseInt(line)
	})

	var best int
	for _, group := range puz.Split(calories, func(n int) bool { return n != 0 }) {
		best = puz.Max(best, puz.Sum(group...))
	}
	fmt.Println(best)
}
