package main

import (
	"fmt"

	"github.com/bbeck/advent-of-code/puz"
)

func main() {
	var fuel int
	for _, mass := range puz.InputToInts(2019, 1) {
		fuel += mass/3 - 2
	}

	fmt.Println(fuel)
}
