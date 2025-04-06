package main

import (
	"fmt"
	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	var frequency int
	for _, i := range in.Ints() {
		frequency += i
	}

	fmt.Println(frequency)
}
