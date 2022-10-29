package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	ship := aoc.Point2D{}
	waypoint := aoc.Point2D{X: 10, Y: 1}

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
				waypoint = aoc.Point2D{X: -waypoint.Y, Y: waypoint.X}
			}
		case "R":
			for n := 0; n < instruction.Value/90; n++ {
				waypoint = aoc.Point2D{X: waypoint.Y, Y: -waypoint.X}
			}
		case "F":
			ship.X += instruction.Value * waypoint.X
			ship.Y += instruction.Value * waypoint.Y
		}
	}

	fmt.Println(aoc.Origin2D.ManhattanDistance(ship))
}

type Instruction struct {
	Action string
	Value  int
}

func InputToInstructions() []Instruction {
	return aoc.InputLinesTo(2020, 12, func(line string) (Instruction, error) {
		return Instruction{
			Action: string(line[0]),
			Value:  aoc.ParseInt(line[1:]),
		}, nil
	})
}
