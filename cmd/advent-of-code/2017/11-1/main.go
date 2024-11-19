package main

import (
	"fmt"
	"strings"

	"github.com/bbeck/advent-of-code/lib"
)

func main() {
	deltas := map[string]lib.Point2D{
		"n":  {X: 0, Y: -1},
		"ne": {X: 1, Y: -1},
		"se": {X: 1, Y: 0},
		"s":  {X: 0, Y: 1},
		"sw": {X: -1, Y: 1},
		"nw": {X: -1, Y: 0},
	}

	var location lib.Point2D
	for _, step := range InputToSteps() {
		delta := deltas[step]
		location = lib.Point2D{X: location.X + delta.X, Y: location.Y + delta.Y}
	}

	fmt.Println(HexDistance(lib.Origin2D, location))
}

func HexDistance(a, b lib.Point2D) int {
	// https://www.redblobgames.com/grids/hexagons/#distances
	aq, ar, as := a.X, a.Y, -a.X-a.Y
	bq, br, bs := b.X, b.Y, -b.X-b.Y
	return (lib.Abs(aq-bq) + lib.Abs(ar-br) + lib.Abs(as-bs)) / 2
}

func InputToSteps() []string {
	s := lib.InputToString()
	return strings.Split(s, ",")
}
