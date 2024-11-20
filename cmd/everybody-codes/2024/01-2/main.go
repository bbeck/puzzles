package main

import (
	"fmt"
	"github.com/bbeck/puzzles/lib"
)

var Potions = map[byte]int{
	'B': 1,
	'C': 3,
	'D': 5,
}

var Empties = map[byte]int{
	'x': 1,
}

func main() {
	var count int
	for _, group := range lib.Chunk(lib.InputToBytes(), 2) {
		count += Potions[group[0]] + Potions[group[1]]

		if Empties[group[0]]+Empties[group[1]] == 0 {
			count += 2
		}
	}

	fmt.Println(count)
}
