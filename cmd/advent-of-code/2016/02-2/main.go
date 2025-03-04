package main

import (
	"fmt"
	"strings"

	. "github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	digits := in.LinesTo[string](func(s string) string {
		return KeypadDigit(s)
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

var Start = Point2D{X: 1, Y: 3}

func KeypadDigit(s string) string {
	p := Start
	for _, c := range s {
		var next Point2D
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
