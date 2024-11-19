package main

import (
	"fmt"

	"github.com/bbeck/advent-of-code/lib"
)

func main() {
	var frequency int
	for _, i := range lib.InputToInts() {
		frequency += i
	}

	fmt.Println(frequency)
}
