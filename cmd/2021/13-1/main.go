package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/aoc"
	"log"
)

func main() {
	points, instructions := InputToPaper()
	for _, instruction := range instructions {
		points = Fold(points, instruction.axis, instruction.offset)
		break
	}

	fmt.Println(points.Size())
}

func Fold(points aoc.Set, axis string, offset int) aoc.Set {
	next := aoc.NewSet()
	for _, o := range points.Entries() {
		point := o.(aoc.Point2D)

		if axis == "x" && point.X >= offset {
			point = aoc.Point2D{X: 2*offset - point.X, Y: point.Y}
		} else if axis == "y" && point.Y >= offset {
			point = aoc.Point2D{X: point.X, Y: 2*offset - point.Y}
		}

		next.Add(point)
	}

	return next
}

type Instruction struct {
	axis   string
	offset int
}

func InputToPaper() (aoc.Set, []Instruction) {
	var points = aoc.NewSet()
	var instructions []Instruction
	for _, line := range aoc.InputToLines(2021, 13) {
		if line == "" {
			continue
		}

		var point aoc.Point2D
		if _, err := fmt.Sscanf(line, "%d,%d", &point.X, &point.Y); err == nil {
			points.Add(point)
			continue
		}

		var instruction Instruction
		if _, err := fmt.Sscanf(line, "fold along x=%d", &instruction.offset); err == nil {
			instruction.axis = "x"
			instructions = append(instructions, instruction)
			continue
		}

		if _, err := fmt.Sscanf(line, "fold along y=%d", &instruction.offset); err == nil {
			instruction.axis = "y"
			instructions = append(instructions, instruction)
			continue
		}

		log.Fatalf("unable to parse line: %s", line)
	}

	return points, instructions
}
