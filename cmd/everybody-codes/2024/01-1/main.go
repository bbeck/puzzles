package main

import (
	"fmt"
	"github.com/bbeck/puzzles/lib"
)

var Potions = map[byte]int{
	'B': 1,
	'C': 3,
}

func main() {
	var count int
	for _, enemy := range lib.InputToBytes() {
		count += Potions[enemy]
	}

	fmt.Println(count)
}
