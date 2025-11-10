package main

import (
	"fmt"

	. "github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	f := InputToBone()

	var digits []int
	for _, level := range f.Levels {
		digits = append(digits, Digits(level[CENTER])...)
	}
	fmt.Println(JoinDigits(digits))
}

type Fishbone struct {
	ID     int
	Levels [][3]int
}

const (
	LEFT = iota
	CENTER
	RIGHT
)

func InputToBone() Fishbone {
	return in.LinesToS(func(in in.Scanner[Fishbone]) Fishbone {
		id := in.Int()

		var levels [][3]int

	outer:
		for in.HasNext() {
			n := in.Int()

			for current := 0; current < len(levels); current++ {
				if n < levels[current][CENTER] && levels[current][LEFT] == 0 {
					levels[current][LEFT] = n
					continue outer
				}
				if n > levels[current][CENTER] && levels[current][RIGHT] == 0 {
					levels[current][RIGHT] = n
					continue outer
				}
			}

			levels = append(levels, [3]int{0, n, 0})
		}

		return Fishbone{ID: id, Levels: levels}
	})[0]
}
