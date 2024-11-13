package main

import (
	"fmt"
	"strings"

	"github.com/bbeck/advent-of-code/puz"
)

func main() {
	deltas := map[string]puz.Point2D{
		"n":  {X: 0, Y: -1},
		"ne": {X: 1, Y: -1},
		"se": {X: 1, Y: 0},
		"s":  {X: 0, Y: 1},
		"sw": {X: -1, Y: 1},
		"nw": {X: -1, Y: 0},
	}

	var location puz.Point2D
	for _, step := range InputToSteps() {
		delta := deltas[step]
		location = puz.Point2D{X: location.X + delta.X, Y: location.Y + delta.Y}
	}

	fmt.Println(HexDistance(puz.Origin2D, location))
}

func HexDistance(a, b puz.Point2D) int {
	// https://www.redblobgames.com/grids/hexagons/#distances
	aq, ar, as := a.X, a.Y, -a.X-a.Y
	bq, br, bs := b.X, b.Y, -b.X-b.Y
	return (puz.Abs(aq-bq) + puz.Abs(ar-br) + puz.Abs(as-bs)) / 2
}

func InputToSteps() []string {
	s := puz.InputToString(2017, 11)
	return strings.Split(s, ",")
}
