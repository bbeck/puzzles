package main

import (
	"fmt"

	. "github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	screen := NewGrid2D[bool](50, 6)
	for _, instruction := range InputToInstructions() {
		switch instruction.Kind {
		case "rect":
			Rect(screen, instruction.Width, instruction.Height)

		case "rotate row":
			RotateRow(screen, instruction.Row, instruction.Width)

		case "rotate column":
			RotateCol(screen, instruction.Col, instruction.Height)
		}
	}

	var count int
	screen.ForEach(func(x, y int, value bool) {
		if value {
			count++
		}
	})
	fmt.Println(count)
}

func Rect(screen Grid2D[bool], width, height int) {
	for y := range height {
		for x := range width {
			screen.Set(x, y, true)
		}
	}
}

func RotateRow(screen Grid2D[bool], y int, distance int) {
	var row []bool
	for x := range screen.Width {
		row = append(row, screen.Get(x, y))
	}

	for x := range screen.Width {
		screen.Set(x, y, row[(x-distance+screen.Width)%screen.Width])
	}
}

func RotateCol(screen Grid2D[bool], x int, distance int) {
	var col []bool
	for y := range screen.Height {
		col = append(col, screen.Get(x, y))
	}

	for y := range screen.Height {
		screen.Set(x, y, col[(y-distance+screen.Height)%screen.Height])
	}
}

type Instruction struct {
	Kind          string
	Width, Height int
	Col, Row      int
}

func InputToInstructions() []Instruction {
	return in.LinesToS(func(in in.Scanner[Instruction]) Instruction {
		switch {
		case in.HasPrefix("rect"):
			var width, height int
			in.Scanf("rect %dx%d", &width, &height)
			return Instruction{Kind: "rect", Width: width, Height: height}

		case in.HasPrefix("rotate row"):
			var row, distance int
			in.Scanf("rotate row y=%d by %d", &row, &distance)
			return Instruction{Kind: "rotate row", Row: row, Width: distance}

		case in.HasPrefix("rotate column"):
			var col, distance int
			in.Scanf("rotate column x=%d by %d", &col, &distance)
			return Instruction{Kind: "rotate column", Col: col, Height: distance}

		default:
			panic("unsupported prefix")
		}
	})
}
