package main

import (
	"fmt"
	"log"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	goal := InputToCoordinate(2015, 25)
	indices := GenerateIndices(goal.Y, goal.X)

	code := 20151125
	for n := 1; n < indices[goal]; n++ {
		code = (code * 252533) % 33554393
	}

	fmt.Printf("code: %d\n", code)
}

func InputToCoordinate(year, day int) aoc.Point2D {
	s := aoc.InputToString(year, day)

	var row, col int
	_, err := fmt.Sscanf(s, "To continue, please consult the code grid in the manual.  Enter the code at row %d, column %d.", &row, &col)
	if err != nil {
		log.Fatalf("unable to parse input: %+v", err)
	}

	return aoc.Point2D{X: col, Y: row}
}

func GenerateIndices(row, col int) map[aoc.Point2D]int {
	indices := make(map[aoc.Point2D]int)
	current := aoc.Point2D{X: 1, Y: 1}
	maxY := 1

	for index := 1; ; index++ {
		indices[current] = index
		if current.X == col && current.Y == row {
			break
		}

		if current.Y <= 1 {
			maxY++

			current.X = 1
			current.Y = maxY
			continue
		}

		current.X++
		current.Y--
	}

	return indices
}
