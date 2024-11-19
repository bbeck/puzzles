package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/bbeck/puzzles/lib"
)

func main() {
	lights := lib.NewGrid2D[bool](1000, 1000)
	for _, instruction := range InputToInstructions() {
		for x := instruction.TopLeft.X; x <= instruction.BottomRight.X; x++ {
			for y := instruction.TopLeft.Y; y <= instruction.BottomRight.Y; y++ {
				lights.Set(x, y, instruction.Op(lights.Get(x, y)))
			}
		}
	}

	var on int
	for y := 0; y < lights.Height; y++ {
		for x := 0; x < lights.Width; x++ {
			if lights.Get(x, y) {
				on++
			}
		}
	}
	fmt.Println(on)
}

type Instruction struct {
	Op                   func(bool) bool
	TopLeft, BottomRight lib.Point2D
}

var Ops = map[string]func(bool) bool{
	"on":     func(b bool) bool { return true },
	"off":    func(b bool) bool { return false },
	"toggle": func(b bool) bool { return !b },
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
