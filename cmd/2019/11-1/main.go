package main

import (
	"fmt"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	// The robot's location and heading
	robot := &Turtle{
		location:  aoc.Point2D{},
		direction: "N",
	}
	color := -1

	hull := make(map[aoc.Point2D]int)
	painted := make(map[aoc.Point2D]bool)

	cpu := &CPU{
		memory: InputToMemory(2019, 11),
		input: func(addr int) int {
			return hull[robot.location]
		},
		output: func(value int) {
			// Check if we already have the color to paint, if not then this is the
			// color and we need to wait until the next output call for the direction.
			if color == -1 {
				color = value
				return
			}

			// We can act now that we have all of the information
			hull[robot.location] = color
			painted[robot.location] = true

			color = -1
			if value == 0 {
				robot.Left()
			} else {
				robot.Right()
			}
			robot.Forward()
		},
	}
	cpu.Execute()

	fmt.Printf("number of panels painted: %d\n", len(painted))
}

type Turtle struct {
	location  aoc.Point2D
	direction string
}

func (t *Turtle) Forward() {
	switch t.direction {
	case "N":
		t.location = t.location.Up()
	case "E":
		t.location = t.location.Right()
	case "S":
		t.location = t.location.Down()
	case "W":
		t.location = t.location.Left()
	}
}

func (t *Turtle) Left() {
	switch t.direction {
	case "N":
		t.direction = "W"
	case "E":
		t.direction = "N"
	case "S":
		t.direction = "E"
	case "W":
		t.direction = "S"
	}
}

func (t *Turtle) Right() {
	switch t.direction {
	case "N":
		t.direction = "E"
	case "E":
		t.direction = "S"
	case "S":
		t.direction = "W"
	case "W":
		t.direction = "N"
	}
}
