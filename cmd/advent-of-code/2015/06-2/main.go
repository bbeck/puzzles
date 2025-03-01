package main

import (
	"fmt"

	. "github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	lights := NewGrid2D[int](1000, 1000)
	for _, instruction := range InputToInstructions() {
		for x := instruction.TopLeft.X; x <= instruction.BottomRight.X; x++ {
			for y := instruction.TopLeft.Y; y <= instruction.BottomRight.Y; y++ {
				lights.Set(x, y, instruction.Op(lights.Get(x, y)))
			}
		}
	}

	var brightness int
	for y := range lights.Height {
		for x := range lights.Width {
			brightness += lights.Get(x, y)
		}
	}
	fmt.Println(brightness)
}

type Instruction struct {
	Op                   func(int) int
	TopLeft, BottomRight Point2D
}

var Ops = map[string]func(int) int{
	"turn on":  func(n int) int { return n + 1 },
	"turn off": func(n int) int { return Max(n-1, 0) },
	"toggle":   func(n int) int { return n + 2 },
}

func InputToInstructions() []Instruction {
	var instructions []Instruction
	for in.HasNext() {
		op := in.OneOf("turn on", "turn off", "toggle")
		p1x, p1y, p2x, p2y := in.Int(), in.Int(), in.Int(), in.Int()

		instructions = append(instructions, Instruction{
			Op:          Ops[op],
			TopLeft:     Point2D{X: Min(p1x, p2x), Y: Min(p1y, p2y)},
			BottomRight: Point2D{X: Max(p1x, p2x), Y: Max(p1y, p2y)},
		})
	}
	return instructions
}
