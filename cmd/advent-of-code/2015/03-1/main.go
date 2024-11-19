package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/lib"
)

func main() {
	var seen lib.Set[lib.Point2D]
	seen.Add(lib.Origin2D)

	var location lib.Point2D
	for _, dir := range InputToDirections() {
		location = lib.Point2D{X: location.X + dir.X, Y: location.Y + dir.Y}
		seen.Add(location)
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
