package main

import (
	"fmt"
	"strings"

	. "github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	points, moves := InputToPointsAndMoves()

	var current = points["START"]
	var visited = SetFrom(current)
	for _, move := range moves {
		current = Point2D{
			X: (current.X + points[move].X) / 2,
			Y: (current.Y + points[move].Y) / 2,
		}
		visited.Add(current)
	}
	fmt.Println(len(visited))
}

func InputToPointsAndMoves() (map[string]Point2D, []string) {
	var points = make(map[string]Point2D)
	var moves []string
	for in.HasNext() {
		lhs, rhs := in.CutS[any]("=")
		if name := lhs.String(); name != "MOVES" {
			points[name] = Point2D{X: rhs.Int(), Y: rhs.Int()}
		} else {
			moves = strings.Split(rhs.String(), "")
		}
	}

	return points, moves
}
