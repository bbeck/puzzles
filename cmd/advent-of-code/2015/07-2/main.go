package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	circuit := InputToCircuit()
	a := Evaluate(circuit, "a", make(map[string]uint16))
	a = Evaluate(circuit, "a", map[string]uint16{"b": a})
	fmt.Println(a)
}

func Evaluate(circuit Circuit, variable string, memory map[string]uint16) uint16 {
	if value, found := memory[variable]; found {
		return value
	}

	// Sometimes we're asked to evaluate a numeric constant.
	if n, err := strconv.Atoi(variable); err == nil {
		memory[variable] = uint16(n)
		return uint16(n)
	}

	var value uint16
	switch expr := circuit[variable]; len(expr) {
	case 1:
		value = Evaluate(circuit, expr[0], memory)

	case 2:
		switch op := expr[0]; op {
		case "NOT":
			value = ^Evaluate(circuit, expr[1], memory)
		}

	case 3:
		lhs, rhs := Evaluate(circuit, expr[0], memory), Evaluate(circuit, expr[2], memory)
		switch op := expr[1]; op {
		case "AND":
			value = lhs & rhs
		case "OR":
			value = lhs | rhs
		case "LSHIFT":
			value = lhs << rhs
		case "RSHIFT":
			value = lhs >> rhs
		}
	}

	memory[variable] = value
	return value
}

type Circuit map[string][]string

func InputToCircuit() Circuit {
	circuit := make(Circuit)
	for in.HasNext() {
		lhs, rhs := in.Cut(" -> ")
		circuit[rhs] = strings.Fields(lhs)
	}

	return circuit
}
