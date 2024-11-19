package main

import (
	"fmt"
	"github.com/bbeck/puzzles/lib"
)

func main() {
	var seen lib.Set[lib.Point2D]
	seen.Add(lib.Origin2D)

	var santa, robot lib.Point2D
	for i, dir := range InputToDirections() {
		if i%2 == 0 {
			santa = lib.Point2D{X: santa.X + dir.X, Y: santa.Y + dir.Y}
		} else {
			robot = lib.Point2D{X: robot.X + dir.X, Y: robot.Y + dir.Y}
		}
		seen.Add(santa, robot)
	}

	fmt.Println(len(seen))
}

func InputToDirections() []lib.Point2D {
	origin := lib.Origin2D

	var directions []lib.Point2D
	for _, b := range lib.InputToBytes() {
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
