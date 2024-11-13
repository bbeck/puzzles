package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/puz"
	"github.com/bbeck/advent-of-code/puz/cpus"
)

func main() {
	var robot puz.Turtle
	var isPaintInstruction bool

	// The robot starts on a white panel
	panels := map[puz.Point2D]int{
		robot.Location: 1,
	}

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

	tl, br := puz.GetBounds(puz.GetMapKeys(panels))
	for y := tl.Y; y <= br.Y; y++ {
		for x := tl.X; x <= br.X; x++ {
			if panels[puz.Point2D{X: x, Y: y}] == 1 {
				fmt.Print("█")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
