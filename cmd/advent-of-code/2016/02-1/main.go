package main

import (
	"fmt"
	"strings"

	"github.com/bbeck/advent-of-code/puz"
)

func main() {
	digits := puz.InputLinesTo(func(line string) string {
		return KeypadDigit(line)
	})

	fmt.Println(strings.Join(digits, ""))
}

var Keypad = []string{
	".....",
	".123.",
	".456.",
	".789.",
	".....",
}

var Start = puz.Point2D{X: 2, Y: 2}

func KeypadDigit(s string) string {
	p := Start
	for _, c := range s {
		var next puz.Point2D
		switch c {
		case 'U':
			next = p.Up()
		case 'R':
			next = p.Right()
		case 'D':
			next = p.Down()
		case 'L':
			next = p.Left()
		}

		if Keypad[next.Y][next.X] != '.' {
			p = next
		}
	}

	return string(Keypad[p.Y][p.X])
}
