package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	var seen aoc.Set[aoc.Point2D]
	seen.Add(aoc.Origin2D)

	var santa, robot aoc.Point2D
	for i, dir := range InputToDirections() {
		if i%2 == 0 {
			santa = aoc.Point2D{X: santa.X + dir.X, Y: santa.Y + dir.Y}
		} else {
			robot = aoc.Point2D{X: robot.X + dir.X, Y: robot.Y + dir.Y}
		}
		seen.Add(santa, robot)
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
