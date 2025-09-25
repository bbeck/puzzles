package main

import (
	"fmt"
	"sort"

	. "github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	ns := in.Ints()
	sort.Ints(ns)

	goal := ns[len(ns)/2]

	var hits int
	for _, n := range ns {
		hits += Abs(goal - n)
	}
	fmt.Println(hits)
}
