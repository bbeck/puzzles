package main

import (
	"fmt"

	. "github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	lights := NewGrid2D[bool](1000, 1000)
	for _, instruction := range InputToInstructions() {
		for x := instruction.TopLeft.X; x <= instruction.BottomRight.X; x++ {
			for y := instruction.TopLeft.Y; y <= instruction.BottomRight.Y; y++ {
				lights.Set(x, y, instruction.Op(lights.Get(x, y)))
			}
		}
	}

	var on int
	for y := range lights.Height {
		for x := range lights.Width {
			if lights.Get(x, y) {
				on++
			}
		}
	}
	fmt.Println(on)
}

type Instruction struct {
	Op                   func(bool) bool
	TopLeft, BottomRight Point2D
}

var Ops = map[string]func(bool) bool{
	"turn on":  func(b bool) bool { return true },
	"turn off": func(b bool) bool { return false },
	"toggle":   func(b bool) bool { return !b },
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
