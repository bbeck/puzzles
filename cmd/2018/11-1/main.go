package main

import (
	"fmt"
	"math"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	serial := aoc.InputToInt(2018, 11)
	N := 300

	cells := make(map[aoc.Point2D]int)
	for x := 1; x <= N; x++ {
		for y := 1; y <= N; y++ {
			rack := x + 10
			power := (((rack*y)+serial)*rack)/100%10 - 5
			cells[aoc.Point2D{X: x, Y: y}] = power
		}
	}

	var best = math.MinInt64
	var bestP aoc.Point2D
	for x := 1; x <= N-3; x++ {
		for y := 1; y <= N-3; y++ {
			p := aoc.Point2D{X: x, Y: y}
			total := cells[p] + cells[p.Right()] + cells[p.Right().Right()] +
				cells[p.Down()] + cells[p.Down().Right()] + cells[p.Down().Right().Right()] +
				cells[p.Down().Down()] + cells[p.Down().Down().Right()] + cells[p.Down().Down().Right().Right()]

			if total > best {
				best = total
				bestP = p
			}
		}
	}

	fmt.Printf("square with top left %s has best power level of %d\n", bestP, best)
}
