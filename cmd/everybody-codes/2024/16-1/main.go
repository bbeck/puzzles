package main

import (
	"fmt"
	"strings"

	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	turns, wheels := InputToTurnsAndWheels()

	next := func(indices []int) []int {
		var next []int
		for i := range indices {
			next = append(next, (indices[i]+turns[i])%len(wheels[i]))
		}
		return next
	}

	indices := make([]int, len(wheels))
	for range 100 {
		indices = next(indices)
	}

	var symbols []string
	for i := range wheels {
		symbols = append(symbols, wheels[i][indices[i]])
	}
	fmt.Println(strings.Join(symbols, " "))
}

func InputToTurnsAndWheels() ([]int, [][]string) {
	c1 := in.ChunkS()
	turns := c1.Ints()

	var wheels = make([][]string, len(turns))

	c2 := in.ChunkS()
	for _, line := range c2.Lines() {
		for i := range wheels {
			if len(line) >= 4*i+3 {
				symbol := line[4*i : 4*i+3]
				if symbol != "   " {
					wheels[i] = append(wheels[i], symbol)
				}
			}
		}
	}

	return turns, wheels
}
