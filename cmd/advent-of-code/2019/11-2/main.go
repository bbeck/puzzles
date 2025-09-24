package main

import (
	"fmt"

	. "github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/cpus"
)

func main() {
	var robot Turtle
	var isPaintInstruction bool

	// The robot starts on a white panel
	panels := map[Point2D]int{
		robot.Location: 1,
	}

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

	tl, br := GetBounds(GetMapKeys(panels))
	for y := tl.Y; y <= br.Y; y++ {
		for x := tl.X; x <= br.X; x++ {
			if panels[Point2D{X: x, Y: y}] == 1 {
				fmt.Print("█")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
