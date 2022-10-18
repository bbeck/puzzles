package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/aoc"
	"github.com/bbeck/advent-of-code/aoc/cpus"
)

func main() {
	panels := make(map[aoc.Point2D]int)

	var robot aoc.Turtle
	var isPaintInstruction bool

	cpu := cpus.IntcodeCPU{
		Memory: cpus.InputToIntcodeMemory(2019, 11),
		Input: func() int {
			return panels[robot.Location]
		},
		Output: func(value int) {
			isPaintInstruction = !isPaintInstruction
			if isPaintInstruction {
				panels[robot.Location] = value
				return
			}

			if value == 0 {
				robot.TurnLeft()
			} else {
				robot.TurnRight()
			}
			robot.Forward(1)
		},
	}
	cpu.Execute()

	fmt.Println(len(panels))
}
