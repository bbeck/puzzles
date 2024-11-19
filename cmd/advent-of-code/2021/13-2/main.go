package main

import (
	"fmt"
	"github.com/bbeck/puzzles/lib"
	"strings"
)

func main() {
	points, instructions := InputToPaper()
	for _, i := range instructions {
		points = Fold(points, i)
	}

	tl, br := lib.GetBounds(points.Entries())
	for y := tl.Y; y <= br.Y; y++ {
		for x := tl.X; x <= br.X; x++ {
			if points.Contains(lib.Point2D{X: x, Y: y}) {
				fmt.Print("█")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func Fold(points lib.Set[lib.Point2D], instr Instruction) lib.Set[lib.Point2D] {
	var next lib.Set[lib.Point2D]
	for point := range points {
		if instr.Axis == "x" && point.X >= instr.Offset {
			point = lib.Point2D{X: 2*instr.Offset - point.X, Y: point.Y}
		} else if instr.Axis == "y" && point.Y >= instr.Offset {
			point = lib.Point2D{X: point.X, Y: 2*instr.Offset - point.Y}
		}

		next.Add(point)
	}

	return next
}

type Instruction struct {
	Axis   string
	Offset int
}

func InputToPaper() (lib.Set[lib.Point2D], []Instruction) {
	var points lib.Set[lib.Point2D]
	var instructions []Instruction
	for _, line := range lib.InputToLines() {
		if line == "" {
			continue
		}

		if !strings.HasPrefix(line, "fold") {
			var p lib.Point2D
			fmt.Sscanf(line, "%d,%d", &p.X, &p.Y)
			points.Add(p)
		}

		if strings.HasPrefix(line, "fold") {
			line = strings.ReplaceAll(line, "=", " ")

			var i Instruction
			fmt.Sscanf(line, "fold along %s %d", &i.Axis, &i.Offset)
			instructions = append(instructions, i)
		}
	}

	return points, instructions
}
