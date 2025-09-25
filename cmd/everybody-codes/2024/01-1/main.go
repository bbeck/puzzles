package main

import (
	"fmt"

	"github.com/bbeck/puzzles/lib/in"
)

var Potions = map[byte]int{'A': 0, 'B': 1, 'C': 3}

func main() {
	var count int
	for in.HasNext() {
		count += Potions[in.Byte()]
	}

	fmt.Println(count)
}
