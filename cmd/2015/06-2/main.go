package main

import (
	"fmt"
	"strings"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	lights := aoc.NewGrid2D[int](1000, 1000)
	for _, instruction := range InputToInstructions() {
		for x := instruction.TopLeft.X; x <= instruction.BottomRight.X; x++ {
			for y := instruction.TopLeft.Y; y <= instruction.BottomRight.Y; y++ {
				lights.Add(x, y, instruction.Op(lights.Get(x, y)))
			}
		}
	}

	var brightness int
	for y := 0; y < lights.Height; y++ {
		for x := 0; x < lights.Width; x++ {
			brightness += lights.Get(x, y)
		}
	}
	fmt.Println(brightness)
}

type Instruction struct {
	Op                   func(int) int
	TopLeft, BottomRight aoc.Point2D
}

var Ops = map[string]func(int) int{
	"on":     func(n int) int { return n + 1 },
	"off":    func(n int) int { return aoc.Max(n-1, 0) },
	"toggle": func(n int) int { return n + 2 },
}

func InputToInstructions() []Instruction {
	return aoc.InputLinesTo(2015, 6, func(line string) (Instruction, error) {
		// Transform the operation into a single word, this allows it to be parsed by Sscanf
		line = strings.ReplaceAll(line, "turn on", "on")
		line = strings.ReplaceAll(line, "turn off", "off")

		var op string
		var p1x, p1y, p2x, p2y int
		_, err := fmt.Sscanf(line, "%s %d,%d through %d,%d", &op, &p1x, &p1y, &p2x, &p2y)
		if err != nil {
			return Instruction{}, err
		}

		return Instruction{
			Op:          Ops[op],
			TopLeft:     aoc.Point2D{X: aoc.Min(p1x, p2x), Y: aoc.Min(p1y, p2y)},
			BottomRight: aoc.Point2D{X: aoc.Max(p1x, p2x), Y: aoc.Max(p1y, p2y)},
		}, nil
	})
}
