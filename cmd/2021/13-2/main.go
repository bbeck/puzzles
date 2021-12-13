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
	}

	var maxX, maxY int
	for _, o := range points.Entries() {
		point := o.(aoc.Point2D)

		maxX = aoc.MaxInt(maxX, point.X)
		maxY = aoc.MaxInt(maxY, point.Y)
	}

	for y := 0; y <= maxY; y++ {
		for x := 0; x <= maxX; x++ {
			if points.Contains(aoc.Point2D{X: x, Y: y}) {
				fmt.Print("█")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
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
