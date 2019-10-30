package main

import (
	"fmt"
	"strings"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	pipes := InputToPipes(2017, 24)
	length, strength := MaximizeStrength(pipes, 0, 0)
	fmt.Printf("length: %d\n", length)
	fmt.Printf("strength: %d\n", strength)
}

func MaximizeStrength(pipes []Pipe, needed int, used int) (int, int) {
	var length, strength int
	for i, pipe := range pipes {
		var next int
		if used&(1<<i) > 0 {
			continue
		} else if pipe.lhs == needed {
			next = pipe.rhs
		} else if pipe.rhs == needed {
			next = pipe.lhs
		} else {
			continue
		}

		l, s := MaximizeStrength(pipes, next, used|(1<<i))
		l += 1
		s += pipe.lhs + pipe.rhs

		if l > length || (l == length && s > strength) {
			length = l
			strength = s
		}
	}

	return length, strength
}

type Pipe struct {
	lhs, rhs int
}

func InputToPipes(year, day int) []Pipe {
	var pipes []Pipe
	for _, line := range aoc.InputToLines(year, day) {
		parts := strings.Split(line, "/")
		lhs := aoc.ParseInt(parts[0])
		rhs := aoc.ParseInt(parts[1])

		pipes = append(pipes, Pipe{lhs, rhs})
	}

	return pipes
}
