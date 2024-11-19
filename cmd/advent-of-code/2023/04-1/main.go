package main

import (
	"fmt"
	"github.com/bbeck/puzzles/lib"
	"strings"
)

func main() {
	var points int
	for _, winning := range InputToNumWinningNumbers() {
		if winning != 0 {
			points += lib.Pow(2, uint(winning-1))
		}
	}

	fmt.Println(points)
}

func InputToNumWinningNumbers() []int {
	return lib.InputLinesTo(func(line string) int {
		// A number is only winning if it appears more than once per line.
		line = strings.ReplaceAll(line, "|", "")
		_, rhs, _ := strings.Cut(line, ":")

		var seen, winning lib.Set[int]
		for _, field := range strings.Fields(rhs) {
			num := lib.ParseInt(field)
			if !seen.Add(num) {
				winning.Add(num)
			}
		}

		return len(winning)
	})
}
