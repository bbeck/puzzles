package main

import (
	"fmt"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	ship := aoc.Point2D{X: 0, Y: 0}
	waypoint := aoc.Point2D{X: ship.X + 10, Y: ship.Y - 1}

	for _, instruction := range InputToInstructions(2020, 12) {
		switch instruction.action {
		case "N":
			waypoint.Y -= instruction.value

		case "S":
			waypoint.Y += instruction.value

		case "E":
			waypoint.X += instruction.value

		case "W":
			waypoint.X -= instruction.value

		case "L":
			waypoint = RotateClockwise(waypoint, ship, 360-instruction.value)

		case "R":
			waypoint = RotateClockwise(waypoint, ship, instruction.value)

		case "F":
			dx := waypoint.X - ship.X
			dy := waypoint.Y - ship.Y
			ship.X += dx * instruction.value
			ship.Y += dy * instruction.value
			waypoint.X = ship.X + dx
			waypoint.Y = ship.Y + dy
		}
	}

	fmt.Println(ship.ManhattanDistance(aoc.Point2D{}))
}

func RotateClockwise(point aoc.Point2D, origin aoc.Point2D, theta int) aoc.Point2D {
	// Clockwise rotation about the origin
	//   x' = x*cos(theta) - y*sin(theta)
	//   y' = y*cos(theta) + x*sin(theta)
	//
	//   theta  sin(theta)  cos(theta)
	//       0          0           1
	//      90          1           0
	//     180          0          -1
	//     270         -1           0
	sin := map[int]int{90: 1, 180: 0, 270: -1}
	cos := map[int]int{90: 0, 180: -1, 270: 0}

	return aoc.Point2D{
		X: (point.X-origin.X)*cos[theta] - (point.Y-origin.Y)*sin[theta] + origin.X,
		Y: (point.Y-origin.Y)*cos[theta] + (point.X-origin.X)*sin[theta] + origin.Y,
	}
}

type Instruction struct {
	action string
	value  int
}

func InputToInstructions(year, day int) []Instruction {
	var instructions []Instruction
	for _, line := range aoc.InputToLines(year, day) {
		instructions = append(instructions, Instruction{
			action: string(line[0]),
			value:  aoc.ParseInt(line[1:]),
		})
	}

	return instructions
}
