package main

import (
	"fmt"

	. "github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	sequences := InputToSequences()

	var sum int
	for _, s := range sequences {
		sum += Eval(s)
	}
	fmt.Println(sum)
}

func Eval(ns []int) int {
	var seen = SetFrom(0)

	var position int
	for _, n := range ns {
		var next int

		if position-n > 0 && !seen.Contains(position-n) {
			next = position - n
		} else {
			next = position + n
			for seen.Contains(next) {
				next++
			}
		}

		position = next
		seen.Add(position)
	}
	return position
}

func InputToSequences() [][]int {
	return in.LinesToS(func(in in.Scanner[[]int]) []int {
		return in.Ints()
	})
}
