package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/puz"
)

func main() {
	var seen puz.Set[puz.Point2D]
	seen.Add(puz.Origin2D)

	var location puz.Point2D
	for _, dir := range InputToDirections() {
		location = puz.Point2D{X: location.X + dir.X, Y: location.Y + dir.Y}
		seen.Add(location)
	}

	fmt.Println(len(seen))
}

func InputToDirections() []puz.Point2D {
	origin := puz.Origin2D

	var directions []puz.Point2D
	for _, b := range puz.InputToBytes() {
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
