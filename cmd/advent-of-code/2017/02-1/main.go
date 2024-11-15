package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/puz"
	"strings"
)

func main() {
	var sum int
	for _, row := range InputToRows() {
		sum += puz.Max(row...) - puz.Min(row...)
	}
	fmt.Println(sum)
}

func InputToRows() [][]int {
	return puz.InputLinesTo(func(line string) []int {
		var row []int
		for _, s := range strings.Fields(line) {
			row = append(row, puz.ParseInt(s))
		}

		return row
	})
}
