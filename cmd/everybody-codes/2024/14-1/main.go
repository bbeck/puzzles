package main

import (
	"fmt"
	"strings"

	. "github.com/bbeck/puzzles/lib"
)

func main() {
	var maxY int
	var p Point2D
	for _, instruction := range strings.Split(InputToString(), ",") {
		n := ParseInt(instruction[1:])
		switch instruction[0] {
		case 'U':
			p = Point2D{X: p.X, Y: p.Y + n}
		case 'D':
			p = Point2D{X: p.X, Y: p.Y - n}
		case 'L':
			p = Point2D{X: p.X + n, Y: p.Y}
		case 'R':
			p = Point2D{X: p.X - n, Y: p.Y}
		}

		maxY = Max(maxY, p.Y)
	}
	fmt.Println(maxY)
}
