package main

import (
	"fmt"

	"github.com/bbeck/puzzles/lib/in"
)

var Potions = map[byte]int{'A': 0, 'B': 1, 'C': 3, 'D': 5}
var Empties = map[byte]int{'x': 1}

func main() {
	var count int
	for in.HasNext() {
		b1, b2 := in.Byte(), in.Byte()
		count += Potions[b1] + Potions[b2]
		if Empties[b1]+Empties[b2] == 0 {
			count += 2
		}
	}

	fmt.Println(count)
}
