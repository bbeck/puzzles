package main

import (
	"fmt"
	"strings"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	digits := aoc.InputLinesTo(2016, 2, func(line string) string {
		return KeypadDigit(line)
	})

	fmt.Println(strings.Join(digits, ""))
}

var Keypad = []string{
	".......",
	"...1...",
	"..234..",
	".56789.",
	"..ABC..",
	"...D...",
	".......",
}

var Start = aoc.Point2D{X: 1, Y: 3}

func KeypadDigit(s string) string {
	p := Start
	for _, c := range s {
		var next aoc.Point2D
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
