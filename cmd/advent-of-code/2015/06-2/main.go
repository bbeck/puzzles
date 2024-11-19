package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/bbeck/puzzles/lib"
)

func main() {
	lights := lib.NewGrid2D[int](1000, 1000)
	for _, instruction := range InputToInstructions() {
		for x := instruction.TopLeft.X; x <= instruction.BottomRight.X; x++ {
			for y := instruction.TopLeft.Y; y <= instruction.BottomRight.Y; y++ {
				lights.Set(x, y, instruction.Op(lights.Get(x, y)))
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
	TopLeft, BottomRight lib.Point2D
}

var Ops = map[string]func(int) int{
	"on":     func(n int) int { return n + 1 },
	"off":    func(n int) int { return lib.Max(n-1, 0) },
	"toggle": func(n int) int { return n + 2 },
}

func InputToInstructions() []Instruction {
	return lib.InputLinesTo(func(line string) Instruction {
		// Transform the operation into a single word, this allows it to be parsed by Sscanf
		line = strings.ReplaceAll(line, "turn on", "on")
		line = strings.ReplaceAll(line, "turn off", "off")

		var op string
		var p1x, p1y, p2x, p2y int
		_, err := fmt.Sscanf(line, "%s %d,%d through %d,%d", &op, &p1x, &p1y, &p2x, &p2y)
		if err != nil {
			log.Fatalf("unable to parse line: %v", err)
		}

		return Instruction{
			Op:          Ops[op],
			TopLeft:     lib.Point2D{X: lib.Min(p1x, p2x), Y: lib.Min(p1y, p2y)},
			BottomRight: lib.Point2D{X: lib.Max(p1x, p2x), Y: lib.Max(p1y, p2y)},
		}
	})
}
