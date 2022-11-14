package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/aoc"
	"strings"
)

func main() {
	points, instructions := InputToPaper()
	for _, i := range instructions {
		points = Fold(points, i)
	}

	tl, br := aoc.GetBounds(points.Entries())
	for y := tl.Y; y <= br.Y; y++ {
		for x := tl.X; x <= br.X; x++ {
			if points.Contains(aoc.Point2D{X: x, Y: y}) {
				fmt.Print("█")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func Fold(points aoc.Set[aoc.Point2D], instr Instruction) aoc.Set[aoc.Point2D] {
	var next aoc.Set[aoc.Point2D]
	for point := range points {
		if instr.Axis == "x" && point.X >= instr.Offset {
			point = aoc.Point2D{X: 2*instr.Offset - point.X, Y: point.Y}
		} else if instr.Axis == "y" && point.Y >= instr.Offset {
			point = aoc.Point2D{X: point.X, Y: 2*instr.Offset - point.Y}
		}

		next.Add(point)
	}

	return next
}

type Instruction struct {
	Axis   string
	Offset int
}

func InputToPaper() (aoc.Set[aoc.Point2D], []Instruction) {
	var points aoc.Set[aoc.Point2D]
	var instructions []Instruction
	for _, line := range aoc.InputToLines(2021, 13) {
		if line == "" {
			continue
		}

		if !strings.HasPrefix(line, "fold") {
			var p aoc.Point2D
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
