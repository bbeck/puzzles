package main

import (
	"fmt"

	. "github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	ns := in.Ints()
	goal := Min(ns...)

	var hits int
	for _, n := range ns {
		hits += n - goal
	}
	fmt.Println(hits)
}
