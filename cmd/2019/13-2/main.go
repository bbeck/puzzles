package main

import (
	"fmt"

	"github.com/bbeck/advent-of-code/aoc"
	"github.com/bbeck/advent-of-code/aoc/cpus"
)

func main() {
	memory := cpus.InputToIntcodeMemory(2019, 13)
	memory[0] = 2

	// All we care about is the ball, the paddle, and the score.  To keep the
	// game going we just have to keep moving the paddle to be under the ball.
	var ball, paddle aoc.Point2D
	var score int

	var buffer []int
	output := func(value int) {
		buffer = append(buffer, value)
		if len(buffer) < 3 {
			return
		}

		x := buffer[0]
		y := buffer[1]
		id := buffer[2]
		buffer = make([]int, 0)

		switch {
		case x == -1 && y == 0:
			score = id
		case id == 3:
			paddle = aoc.Point2D{x, y}
		case id == 4:
			ball = aoc.Point2D{x, y}
		}
	}

	input := func() int {
		var move int
		if paddle.X > ball.X {
			move = -1
		}
		if paddle.X < ball.X {
			move = 1
		}
		return move
	}

	cpu := cpus.IntcodeCPU{
		Memory: memory,
		Input:  func(addr int) int { return input() },
		Output: func(value int) { output(value) },
	}
	cpu.Execute()
	fmt.Printf("score: %d\n", score)
}
