package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/lib"
)

const N = 1_000_000
const MOVES = 10_000_000

func main() {
	circle, current := InputToCircle()

	for i := 1; i <= MOVES; i++ {
		// We start with the following
		//   ... current v1 v2 v3 x ... destination y ...
		//
		// Then we remove v1, v2, and v3, and place them after destination.
		//   ... current x ... destination v1 v2 v3 y ...
		v1 := circle[current]
		v2 := circle[v1]
		v3 := circle[v2]
		x := circle[v3]

		destination := Destination(current, v1, v2, v3)
		y := circle[destination]

		circle[current] = x
		circle[destination] = v1
		circle[v3] = y
		current = x
	}

	a := circle[1]
	b := circle[a]
	fmt.Println(a * b)
}

func Destination(current, r1, r2, r3 int) int {
	destination := current
	for {
		destination -= 1
		if destination == 0 {
			destination = N
		}

		if destination != r1 && destination != r2 && destination != r3 {
			break
		}
	}

	return destination
}

func InputToCircle() ([]int, int) {
	digits := lib.Digits(lib.InputToInt())

	circle := make([]int, N+1)

	current := N
	for _, d := range digits {
		circle[current] = d
		current = d
	}
	for d := 10; d <= N; d++ {
		circle[current] = d
		current = d
	}
	circle[current] = digits[0]

	return circle, digits[0]
}
