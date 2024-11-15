package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/puz"
)

func main() {
	ship := puz.Turtle{Heading: puz.Right}
	for _, instruction := range InputToInstructions() {
		switch instruction.Action {
		case "N":
			ship.Location.Y -= instruction.Value
		case "S":
			ship.Location.Y += instruction.Value
		case "E":
			ship.Location.X += instruction.Value
		case "W":
			ship.Location.X -= instruction.Value
		case "L":
			for n := 0; n < instruction.Value/90; n++ {
				ship.TurnLeft()
			}
		case "R":
			for n := 0; n < instruction.Value/90; n++ {
				ship.TurnRight()
			}
		case "F":
			ship.Forward(instruction.Value)
		}
	}

	fmt.Println(puz.Origin2D.ManhattanDistance(ship.Location))
}

type Instruction struct {
	Action string
	Value  int
}

func InputToInstructions() []Instruction {
	return puz.InputLinesTo(func(line string) Instruction {
		return Instruction{
			Action: string(line[0]),
			Value:  puz.ParseInt(line[1:]),
		}
	})
}
