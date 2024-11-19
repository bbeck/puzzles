package main

import (
	"fmt"
	"github.com/bbeck/puzzles/lib"
	"strings"
)

func main() {
	var sum int
	for _, row := range InputToRows() {
		sum += lib.Max(row...) - lib.Min(row...)
	}
	fmt.Println(sum)
}

func InputToRows() [][]int {
	return lib.InputLinesTo(func(line string) []int {
		var row []int
		for _, s := range strings.Fields(line) {
			row = append(row, lib.ParseInt(s))
		}

		return row
	})
}
