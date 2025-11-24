package main

import (
	"fmt"

	"github.com/bbeck/puzzles/lib/in"
)

const N = 90

func main() {
	var sum int
	for _, n := range in.Ints() {
		sum += N / n
	}
	fmt.Println(sum)
}
