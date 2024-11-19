package main

import (
	"fmt"
	"github.com/bbeck/puzzles/lib"
)

func main() {
	ship := lib.Point2D{}
	waypoint := lib.Point2D{X: 10, Y: 1}

	for _, instruction := range InputToInstructions() {
		switch instruction.Action {
		case "N":
			waypoint.Y += instruction.Value
		case "S":
			waypoint.Y -= instruction.Value
		case "E":
			waypoint.X += instruction.Value
		case "W":
			waypoint.X -= instruction.Value
		case "L":
			for n := 0; n < instruction.Value/90; n++ {
				waypoint = lib.Point2D{X: -waypoint.Y, Y: waypoint.X}
			}
		case "R":
			for n := 0; n < instruction.Value/90; n++ {
				waypoint = lib.Point2D{X: waypoint.Y, Y: -waypoint.X}
			}
		case "F":
			ship.X += instruction.Value * waypoint.X
			ship.Y += instruction.Value * waypoint.Y
		}
	}

	fmt.Println(lib.Origin2D.ManhattanDistance(ship))
}

type Instruction struct {
	Action string
	Value  int
}

func InputToInstructions() []Instruction {
	return lib.InputLinesTo(func(line string) Instruction {
		return Instruction{
			Action: string(line[0]),
			Value:  lib.ParseInt(line[1:]),
		}
	})
}
