package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/puz"
)

func main() {
	var seen puz.Set[puz.Point2D]
	seen.Add(puz.Origin2D)

	var santa, robot puz.Point2D
	for i, dir := range InputToDirections() {
		if i%2 == 0 {
			santa = puz.Point2D{X: santa.X + dir.X, Y: santa.Y + dir.Y}
		} else {
			robot = puz.Point2D{X: robot.X + dir.X, Y: robot.Y + dir.Y}
		}
		seen.Add(santa, robot)
	}

	fmt.Println(len(seen))
}

func InputToDirections() []puz.Point2D {
	origin := puz.Origin2D

	var directions []puz.Point2D
	for _, b := range puz.InputToBytes(2015, 3) {
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
