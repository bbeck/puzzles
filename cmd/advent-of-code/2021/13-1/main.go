package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/puz"
	"strings"
)

func main() {
	points, instructions := InputToPaper()
	points = Fold(points, instructions[0])
	fmt.Println(len(points))
}

func Fold(points puz.Set[puz.Point2D], instruction Instruction) puz.Set[puz.Point2D] {
	var next puz.Set[puz.Point2D]
	for point := range points {
		if instruction.Axis == "x" && point.X >= instruction.Offset {
			point = puz.Point2D{X: 2*instruction.Offset - point.X, Y: point.Y}
		} else if instruction.Axis == "y" && point.Y >= instruction.Offset {
			point = puz.Point2D{X: point.X, Y: 2*instruction.Offset - point.Y}
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
