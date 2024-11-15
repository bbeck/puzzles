package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/puz"
	"strings"
)

func main() {
	var points int
	for _, winning := range InputToNumWinningNumbers() {
		if winning != 0 {
			points += puz.Pow(2, uint(winning-1))
		}
	}

	fmt.Println(points)
}

func InputToNumWinningNumbers() []int {
	return puz.InputLinesTo(2023, 4, func(line string) int {
		// A number is only winning if it appears more than once per line.
		line = strings.ReplaceAll(line, "|", "")
		_, rhs, _ := strings.Cut(line, ":")

		var seen, winning puz.Set[int]
		for _, field := range strings.Fields(rhs) {
			num := puz.ParseInt(field)
			if !seen.Add(num) {
				winning.Add(num)
			}
		}

		return len(winning)
	})
}
