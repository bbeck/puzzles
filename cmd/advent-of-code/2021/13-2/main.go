package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/puz"
	"strings"
)

func main() {
	points, instructions := InputToPaper()
	for _, i := range instructions {
		points = Fold(points, i)
	}

	tl, br := puz.GetBounds(points.Entries())
	for y := tl.Y; y <= br.Y; y++ {
		for x := tl.X; x <= br.X; x++ {
			if points.Contains(puz.Point2D{X: x, Y: y}) {
				fmt.Print("█")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func Fold(points puz.Set[puz.Point2D], instr Instruction) puz.Set[puz.Point2D] {
	var next puz.Set[puz.Point2D]
	for point := range points {
		if instr.Axis == "x" && point.X >= instr.Offset {
			point = puz.Point2D{X: 2*instr.Offset - point.X, Y: point.Y}
		} else if instr.Axis == "y" && point.Y >= instr.Offset {
			point = puz.Point2D{X: point.X, Y: 2*instr.Offset - point.Y}
		}

		next.Add(point)
	}

	return next
}

type Instruction struct {
	Axis   string
	Offset int
}

func InputToPaper() (puz.Set[puz.Point2D], []Instruction) {
	var points puz.Set[puz.Point2D]
	var instructions []Instruction
	for _, line := range puz.InputToLines(2021, 13) {
		if line == "" {
			continue
		}

		if !strings.HasPrefix(line, "fold") {
			var p puz.Point2D
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
