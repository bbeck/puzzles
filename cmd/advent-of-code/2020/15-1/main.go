package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/puz"
	"strings"
)

func main() {
	ns := InputToInts()
	memory := make(map[int]int)

	var last int
	for turn := 0; turn < 2020; turn++ {
		var next int
		if turn < len(ns) {
			next = ns[turn]
		} else if when, found := memory[last]; found {
			next = turn - when
		} else {
			next = 0
		}
		memory[last] = turn
		last = next
	}

	fmt.Println(last)
}

func InputToInts() []int {
	var ns []int
	for _, s := range strings.Split(puz.InputToString(), ",") {
		ns = append(ns, puz.ParseInt(s))
	}
	return ns
}
