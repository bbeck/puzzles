package main

import (
	"fmt"

	"github.com/bbeck/puzzles/lib"
)

func main() {
	var fuel int
	for _, mass := range lib.InputToInts() {
		fuel += mass/3 - 2
	}

	fmt.Println(fuel)
}
