package main

import (
	"fmt"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	location := aoc.Point2D{X: 0, Y: 0}
	direction := aoc.Point2D{X: 1, Y: 0}

	for _, instruction := range InputToInstructions(2020, 12) {
		switch {
		case instruction.action == "N":
			location.Y -= instruction.value

		case instruction.action == "S":
			location.Y += instruction.value

		case instruction.action == "E":
			location.X += instruction.value

		case instruction.action == "W":
			location.X -= instruction.value

		case instruction.action == "L":
			direction = RotateClockwise(direction, 360-instruction.value)

		case instruction.action == "R":
			direction = RotateClockwise(direction, instruction.value)

		case instruction.action == "F":
			location.X += direction.X * instruction.value
			location.Y += direction.Y * instruction.value
		}
	}

	fmt.Println(location.ManhattanDistance(aoc.Point2D{}))
}

func RotateClockwise(point aoc.Point2D, theta int) aoc.Point2D {
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
		X: point.X*cos[theta] - point.Y*sin[theta],
		Y: point.Y*cos[theta] + point.X*sin[theta],
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
