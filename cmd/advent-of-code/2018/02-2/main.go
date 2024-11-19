package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/lib"
)

func main() {
	ids := lib.InputToLines()

	var i, j int
outer:
	for i = 0; i < len(ids); i++ {
		for j = i + 1; j < len(ids); j++ {
			var count int
			for c := 0; count < 2 && c < len(ids[i]); c++ {
				if ids[i][c] != ids[j][c] {
					count++
				}
			}

			if count < 2 {
				break outer
			}
		}
	}

	for k := 0; k < len(ids[i]); k++ {
		if ids[i][k] == ids[j][k] {
			fmt.Printf("%c", ids[i][k])
		}
	}
	fmt.Println()
}
