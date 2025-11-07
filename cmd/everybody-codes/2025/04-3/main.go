package main

import (
	"fmt"

	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	teeth := in.Ints()

	var ratio = 1.
	for i := 0; i < len(teeth); i += 2 {
		ratio *= float64(teeth[i]) / float64(teeth[i+1])
	}
	fmt.Println(int(100 * ratio))
}
