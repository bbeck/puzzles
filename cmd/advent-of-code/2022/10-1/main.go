package main

import (
	"fmt"
	"strings"

	"github.com/bbeck/advent-of-code/lib"
)

func main() {
	// The signal value for each cycle.  Seeded with 1 since x starts at 1.
	signals := []int{1}

	for _, line := range lib.InputToLines() {
		x := signals[len(signals)-1]

		switch op, arg, _ := strings.Cut(line, " "); op {
		case "addx":
			signals = append(signals, []int{x, x + lib.ParseInt(arg)}...)

		case "noop":
			signals = append(signals, x)
		}
	}

	var sum int
	for _, cycle := range []int{20, 60, 100, 140, 180, 220} {
		sum += cycle * signals[cycle-1]
	}
	fmt.Println(sum)
}
