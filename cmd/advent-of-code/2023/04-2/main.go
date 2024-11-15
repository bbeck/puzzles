package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/puz"
	"strings"
)

func main() {
	counts := make(map[int]int)
	for id, winning := range InputToNumWinningNumbers() {
		counts[id]++

		for i := 1; i <= winning; i++ {
			counts[id+i] += counts[id]
		}
	}

	var total int
	for _, count := range counts {
		total += count
	}
	fmt.Println(total)
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
