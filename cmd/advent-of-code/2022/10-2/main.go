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

	for y := 0; y < 6; y++ {
		for x := 0; x < 40; x++ {
			if cycle := y*40 + x; x-1 <= signals[cycle] && signals[cycle] <= x+1 {
				fmt.Print("█")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
