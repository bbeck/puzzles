package main

import (
	"fmt"

	"github.com/bbeck/advent-of-code/lib"
	"github.com/bbeck/advent-of-code/lib/cpus"
)

func main() {
	var ball, paddle lib.Point2D
	var score int

	input := func() int {
		// Move the paddle under the ball.  They're never more than one step away
		// from each other, so their difference is always one of +1/0/-1.
		return ball.X - paddle.X
	}

	var buffer []int
	output := func(value int) {
		buffer = append(buffer, value)
		if len(buffer) < 3 {
			return
		}

		x, y, id := buffer[0], buffer[1], buffer[2]
		buffer = nil

		switch {
		case id == 3:
			paddle = lib.Point2D{X: x, Y: y}
		case id == 4:
			ball = lib.Point2D{X: x, Y: y}
		case id > 4 && x == -1 && y == 0:
			score = id
		}
	}

	memory := cpus.InputToIntcodeMemory()
	memory[0] = 2

	cpu := cpus.IntcodeCPU{Memory: memory, Input: input, Output: output}
	cpu.Execute()
	fmt.Println(score)
}
