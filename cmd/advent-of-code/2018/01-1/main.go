package main

import (
	"fmt"

	"github.com/bbeck/advent-of-code/puz"
)

func main() {
	var frequency int
	for _, i := range puz.InputToInts() {
		frequency += i
	}

	fmt.Println(frequency)
}
