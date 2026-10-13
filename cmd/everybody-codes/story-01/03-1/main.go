package main

import (
	"fmt"

	. "github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	var snails = InputToSnails()
	for range 100 {
		for i := range snails {
			snails[i] = Step(snails[i])
		}
	}

	var sum int
	for _, snail := range snails {
		sum += snail.X + 100*snail.Y
	}
	fmt.Println(sum)
}

func Step(p Point2D) Point2D {
	p.X += 1
	p.Y -= 1

	if p.Y <= 0 {
		p.Y += p.X - 1
		p.X = 1
	}

	return p
}

func InputToSnails() []Point2D {
	return in.LinesToS(func(in in.Scanner[Point2D]) Point2D {
		return Point2D{X: in.Int(), Y: in.Int()}
	})
}
