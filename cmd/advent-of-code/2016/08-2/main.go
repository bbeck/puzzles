package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/puz"
	"log"
)

func main() {
	screen := puz.NewGrid2D[bool](50, 6)
	for _, instruction := range InputToInstructions() {
		if instruction.Kind == "rect" {
			Rect(screen, instruction.Width, instruction.Height)
		}
		if instruction.Kind == "rotate" && instruction.Width > 0 {
			RotateRow(screen, instruction.Row, instruction.Width)
		}
		if instruction.Kind == "rotate" && instruction.Height > 0 {
			RotateCol(screen, instruction.Col, instruction.Height)
		}
	}

	Show(screen)
}

func Show(screen puz.Grid2D[bool]) {
	for y := 0; y < screen.Height; y++ {
		for x := 0; x < screen.Width; x++ {
			if screen.Get(x, y) {
				fmt.Print("█")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func Rect(screen puz.Grid2D[bool], width, height int) {
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			screen.Set(x, y, true)
		}
	}
}

func RotateRow(screen puz.Grid2D[bool], y int, distance int) {
	var row []bool
	for x := 0; x < screen.Width; x++ {
		row = append(row, screen.Get(x, y))
	}

	for x := 0; x < screen.Width; x++ {
		screen.Set(x, y, row[(x-distance+screen.Width)%screen.Width])
	}
}

func RotateCol(screen puz.Grid2D[bool], x int, distance int) {
	var col []bool
	for y := 0; y < screen.Height; y++ {
		col = append(col, screen.Get(x, y))
	}

	for y := 0; y < screen.Height; y++ {
		screen.Set(x, y, col[(y-distance+screen.Height)%screen.Height])
	}
}

type Instruction struct {
	Kind          string
	Width, Height int
	Col, Row      int
}

func InputToInstructions() []Instruction {
	return puz.InputLinesTo(2016, 8, func(line string) Instruction {
		var instruction Instruction
		if _, err := fmt.Sscanf(line, "%s %dx%d", &instruction.Kind, &instruction.Width, &instruction.Height); err == nil {
			return instruction
		}
		if _, err := fmt.Sscanf(line, "%s row y=%d by %d", &instruction.Kind, &instruction.Row, &instruction.Width); err == nil {
			return instruction
		}
		if _, err := fmt.Sscanf(line, "%s column x=%d by %d", &instruction.Kind, &instruction.Col, &instruction.Height); err == nil {
			return instruction
		}
		log.Fatalf("unable to parse line: %s", line)
		return instruction
	})
}
