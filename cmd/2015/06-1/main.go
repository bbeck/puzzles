package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/bbeck/advent-of-code/aoc"
)

type Operation func(lights [][]bool, x int, y int)

var (
	On Operation = func(lights [][]bool, x, y int) {
		lights[x][y] = true
	}
	Off Operation = func(lights [][]bool, x, y int) {
		lights[x][y] = false
	}
	Toggle Operation = func(lights [][]bool, x, y int) {
		lights[x][y] = !lights[x][y]
	}
)

type Instruction struct {
	Op                   Operation
	TopLeft, BottomRight aoc.Point2D
}

func main() {
	lights := make([][]bool, 1000)
	for i := 0; i < len(lights); i++ {
		lights[i] = make([]bool, 1000)
	}

	for _, instruction := range InputToInstructions(2015, 6) {
		minX := aoc.MinInt(instruction.TopLeft.X, instruction.BottomRight.X)
		maxX := aoc.MaxInt(instruction.TopLeft.X, instruction.BottomRight.X)
		minY := aoc.MinInt(instruction.TopLeft.Y, instruction.BottomRight.Y)
		maxY := aoc.MaxInt(instruction.TopLeft.Y, instruction.BottomRight.Y)

		for y := minY; y <= maxY; y++ {
			for x := minX; x <= maxX; x++ {
				instruction.Op(lights, x, y)
			}
		}
	}

	var count int
	for x := 0; x < len(lights); x++ {
		for y := 0; y < len(lights[x]); y++ {
			if lights[x][y] {
				count++
			}
		}
	}

	fmt.Printf("number of lit lights: %d\n", count)
}

func InputToInstructions(year, day int) []Instruction {
	instructions := make([]Instruction, 0)
	for _, line := range aoc.InputToLines(year, day) {
		instructions = append(instructions, NewInstruction(line))
	}

	return instructions
}

func NewInstruction(s string) Instruction {
	// Normalize the string first so it can be parsed by Sscanf
	s = strings.ReplaceAll(s, "turn on", "turn_on")
	s = strings.ReplaceAll(s, "turn off", "turn_off")

	var opcode string
	var tlx, tly, brx, bry int
	_, err := fmt.Sscanf(s, "%s %d,%d through %d,%d", &opcode, &tlx, &tly, &brx, &bry)
	if err != nil {
		log.Fatalf("error parsing instruction from line: %s: %+v", s, err)
	}

	var op Operation
	switch opcode {
	case "turn_on":
		op = On
	case "turn_off":
		op = Off
	case "toggle":
		op = Toggle
	default:
		log.Fatalf("unrecognized opcode: %s", opcode)
	}

	return Instruction{
		Op:          op,
		TopLeft:     aoc.Point2D{tlx, tly},
		BottomRight: aoc.Point2D{brx, bry},
	}
}
