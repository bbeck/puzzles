package main

import (
	"fmt"

	. "github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/in"
)

var Deltas = map[byte]int{'L': -1, 'R': +1}

func main() {
	var count int

	var dial = 50
	for in.HasNext() {
		line := in.Line()
		step, num := Deltas[line[0]], ParseInt(line[1:])

		for range num {
			dial = Modulo(dial+step, 100)
			if dial == 0 {
				count++
			}
		}
	}

	fmt.Println(count)
}
