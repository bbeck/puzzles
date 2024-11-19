package main

import (
	"fmt"
	"github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/cpus"
)

func main() {
	panels := make(map[lib.Point2D]int)

	var robot lib.Turtle
	var isPaintInstruction bool

	cpu := cpus.IntcodeCPU{
		Memory: cpus.InputToIntcodeMemory(),
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
