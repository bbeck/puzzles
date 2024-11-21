package main

import (
	"fmt"
	. "github.com/bbeck/puzzles/lib"
)

func main() {
	ns := InputToInts()
	goal := Min(ns...)

	var hits int
	for _, n := range ns {
		hits += n - goal
	}
	fmt.Println(hits)
}
