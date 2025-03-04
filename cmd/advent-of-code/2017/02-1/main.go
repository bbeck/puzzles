package main

import (
	"fmt"
	. "github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	var sum int
	for _, row := range InputToRows() {
		sum += Max(row...) - Min(row...)
	}
	fmt.Println(sum)
}

func InputToRows() [][]int {
	return in.LinesToS(func(in in.Scanner[[]int]) []int {
		return in.Ints()
	})
}
