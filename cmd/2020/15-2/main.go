package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/aoc"
	"strings"
)

func main() {
	memory := make([]int, 30000000)
	for i := 0; i < len(memory); i++ {
		memory[i] = -1
	}

	ns := InputToInts()

	var last int
	for turn := 0; turn < 30000000; turn++ {
		var next int
		if turn < len(ns) {
			next = ns[turn]
		} else if when := memory[last]; when != -1 {
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
	for _, s := range strings.Split(aoc.InputToString(2020, 15), ",") {
		ns = append(ns, aoc.ParseInt(s))
	}
	return ns
}
