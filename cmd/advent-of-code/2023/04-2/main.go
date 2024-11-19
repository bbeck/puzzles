package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/lib"
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
