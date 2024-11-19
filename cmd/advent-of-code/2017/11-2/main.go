package main

import (
	"fmt"
	"strings"

	"github.com/bbeck/puzzles/lib"
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
	var furthest int
	for _, step := range InputToSteps() {
		delta := deltas[step]
		location = lib.Point2D{X: location.X + delta.X, Y: location.Y + delta.Y}
		furthest = lib.Max(furthest, HexDistance(lib.Origin2D, location))
	}

	fmt.Println(furthest)
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
