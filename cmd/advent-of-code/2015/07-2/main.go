package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/lib"
	"strconv"
	"strings"
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
		return uint16(n)
	}

	var value uint16
	switch expr := circuit[variable]; expr.Kind {
	case "nullary":
		value = Evaluate(circuit, expr.RHS, memory)

	case "urnary":
		switch expr.Op {
		case "NOT":
			value = ^Evaluate(circuit, expr.RHS, memory)
		}

	case "binary":
		lhs, rhs := Evaluate(circuit, expr.LHS, memory), Evaluate(circuit, expr.RHS, memory)
		switch expr.Op {
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

type Circuit map[string]Expression

func InputToCircuit() Circuit {
	circuit := make(Circuit)
	for _, line := range lib.InputToLines() {
		parts := strings.Split(line, " -> ")
		circuit[parts[1]] = ParseExpression(parts[0])
	}

	return circuit
}

type Expression struct {
	Kind         string
	LHS, Op, RHS string
}

func ParseExpression(s string) Expression {
	terms := strings.Fields(s)

	if len(terms) == 1 {
		return Expression{
			Kind: "nullary",
			RHS:  terms[0],
		}
	}

	if len(terms) == 2 {
		return Expression{
			Kind: "urnary",
			Op:   terms[0],
			RHS:  terms[1],
		}
	}

	return Expression{
		Kind: "binary",
		LHS:  terms[0],
		Op:   terms[1],
		RHS:  terms[2],
	}
}
