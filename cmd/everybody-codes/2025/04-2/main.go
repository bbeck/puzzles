package main

import (
	"fmt"
	"sort"

	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	gears := in.Ints()
	fmt.Println(
		sort.Search(10000000000000, func(i int) bool {
			return i*gears[0]/gears[(len(gears)-1)] >= 10000000000000
		}),
	)
}
