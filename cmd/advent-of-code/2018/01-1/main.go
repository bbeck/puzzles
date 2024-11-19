package main

import (
	"fmt"

	"github.com/bbeck/puzzles/lib"
)

func main() {
	var frequency int
	for _, i := range lib.InputToInts() {
		frequency += i
	}

	fmt.Println(frequency)
}
