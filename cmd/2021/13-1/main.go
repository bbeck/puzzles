package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/aoc"
	"log"
	"strings"
)

func main() {
	points, instructions := InputToPaper()

	for _, instruction := range instructions {
		if instruction.axis == "x" {
			points = FoldX(points, instruction.offset)
		} else {
			points = FoldY(points, instruction.offset)
		}
		break
	}

	fmt.Println(points.Size())
}

func FoldX(points aoc.Set, offset int) aoc.Set {
	next := aoc.NewSet()
	for _, o := range points.Entries() {
		point := o.(aoc.Point2D)
		if point.X < offset {
			next.Add(point)
		} else {
			next.Add(aoc.Point2D{X: offset - (point.X - offset), Y: point.Y})
		}
	}
	return next
}

func FoldY(points aoc.Set, offset int) aoc.Set {
	next := aoc.NewSet()
	for _, o := range points.Entries() {
		point := o.(aoc.Point2D)
		if point.Y < offset {
			next.Add(point)
		} else {
			next.Add(aoc.Point2D{X: point.X, Y: offset - (point.Y - offset)})
		}
	}
	return next
}

func Show(points aoc.Set) {
	var maxX, maxY int
	for _, p := range points.Entries() {
		maxX = aoc.MaxInt(maxX, p.(aoc.Point2D).X)
		maxY = aoc.MaxInt(maxY, p.(aoc.Point2D).Y)
	}

	for y := 0; y <= maxY; y++ {
		for x := 0; x <= maxX; x++ {
			if points.Contains(aoc.Point2D{x, y}) {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

type Instruction struct {
	axis   string
	offset int
}

func InputToPaper() (aoc.Set, []Instruction) {
	lines := aoc.InputToLines(2021, 13)

	points := aoc.NewSet()
	var i int
	for i < len(lines) {
		if lines[i] == "" {
			i++
			break
		}

		var x, y int
		if _, err := fmt.Sscanf(lines[i], "%d,%d", &x, &y); err != nil {
			log.Fatal(err)
		}

		points.Add(aoc.Point2D{X: x, Y: y})
		i++
	}

	var instructions []Instruction
	for i < len(lines) {
		var s string
		if _, err := fmt.Sscanf(lines[i], "fold along %s", &s); err != nil {
			log.Fatal(err)
		}

		fields := strings.Split(s, "=")
		axis := fields[0]
		offset := aoc.ParseInt(fields[1])

		instructions = append(instructions, Instruction{axis: axis, offset: offset})
		i++
	}

	return points, instructions
}
