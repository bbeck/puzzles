package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	var seen aoc.Set[aoc.Point2D]
	seen.Add(aoc.Origin2D)

	var location aoc.Point2D
	for _, dir := range InputToDirections() {
		location = aoc.Point2D{X: location.X + dir.X, Y: location.Y + dir.Y}
		seen.Add(location)
	}

	fmt.Println(len(seen))
}

func InputToDirections() []aoc.Point2D {
	origin := aoc.Origin2D

	var directions []aoc.Point2D
	for _, b := range aoc.InputToBytes(2015, 3) {
		switch b {
		case '^':
			directions = append(directions, origin.Up())
		case '<':
			directions = append(directions, origin.Left())
		case '>':
			directions = append(directions, origin.Right())
		case 'v':
			directions = append(directions, origin.Down())
		}
	}

	return directions
}
