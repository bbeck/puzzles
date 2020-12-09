package main

import (
	"fmt"

	"github.com/bbeck/advent-of-code/aoc"
	"github.com/bbeck/advent-of-code/aoc/cpus"
)

func main() {
	// The robot's location and heading
	robot := &Turtle{
		location:  aoc.Point2D{},
		direction: "N",
	}
	color := -1

	hull := make(map[aoc.Point2D]int)
	hull[robot.location] = 1

	cpu := cpus.IntcodeCPU{
		Memory: cpus.InputToIntcodeMemory(2019, 11),
		Input: func(addr int) int {
			return hull[robot.location]
		},
		Output: func(value int) {
			// Check if we already have the color to paint, if not then this is the
			// color and we need to wait until the next output call for the direction.
			if color == -1 {
				color = value
				return
			}

			// We can act now that we have all of the information
			hull[robot.location] = color

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

	Show(hull)
}

func Show(m map[aoc.Point2D]int) {
	var ps []aoc.Point2D
	for p := range m {
		ps = append(ps, p)
	}

	minX, minY, maxX, maxY := aoc.GetBounds(ps)
	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			if m[aoc.Point2D{X: x, Y: y}] == 1 {
				fmt.Print("\u2588")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
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
