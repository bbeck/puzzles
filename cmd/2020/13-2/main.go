package main

import (
	"fmt"
	"strings"

	"github.com/bbeck/advent-of-code/aoc"
)

// Part 2 asks us to solve a system of congruences of the form:
//   x = a_i (mod n_i) for i = 1, ..., k
//
// To do this we can use the chinese remainder theorem.

func main() {
	buses := InputToBuses(2020, 13)

	// Convert the buses into our system of congruences
	var as, ns []int
	for i, bus := range buses {
		if bus == -1 {
			continue
		}

		as = append(as, -i)
		ns = append(ns, bus)
	}

	// Now apply the chinese remainder theory to solve the system
	fmt.Println(ChineseRemainderTheorem(as, ns))
}

func ChineseRemainderTheorem(as, ns []int) int {
	var prod = 1
	for _, n := range ns {
		prod *= n
	}

	var sum int
	for i := 0; i < len(as); i++ {
		p := prod / ns[i]
		sum += as[i] * MulInv(p, ns[i]) * p
	}

	for sum < 0 {
		sum += prod
	}

	return sum % prod
}

func MulInv(a, b int) int {
	if b == 1 {
		return 1
	}

	x0, x1 := 0, 1
	for a > 1 {
		q := a / b
		a, b = b, a%b
		x0, x1 = x1-q*x0, x0
	}

	return x1
}

func InputToBuses(year, day int) []int {
	lines := aoc.InputToLines(year, day)

	var ids []int
	for _, id := range strings.Split(lines[1], ",") {
		if id == "x" {
			ids = append(ids, -1)
		} else {
			ids = append(ids, aoc.ParseInt(id))
		}
	}

	return ids
}
