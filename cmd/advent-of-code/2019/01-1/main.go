package main

import (
	"fmt"

	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	var fuel int
	for _, mass := range in.Ints() {
		fuel += mass/3 - 2
	}

	fmt.Println(fuel)
}
