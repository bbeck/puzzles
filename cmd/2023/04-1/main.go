package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/aoc"
	"strings"
)

func main() {
	var points int
	for _, winning := range InputToNumWinningNumbers() {
		if winning != 0 {
			points += aoc.Pow(2, uint(winning-1))
		}
	}

	fmt.Println(points)
}

func InputToNumWinningNumbers() []int {
	return aoc.InputLinesTo(2023, 4, func(line string) (int, error) {
		// A number is only winning if it appears more than once per line.
		line = strings.ReplaceAll(line, "|", "")
		_, rhs, _ := strings.Cut(line, ":")

		var seen, winning aoc.Set[int]
		for _, field := range strings.Fields(rhs) {
			num := aoc.ParseInt(field)
			if !seen.Add(num) {
				winning.Add(num)
			}
		}

		return len(winning), nil
	})
}
