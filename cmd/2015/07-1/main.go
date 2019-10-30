package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/bbeck/advent-of-code/aoc"
)

type Expression func(c *Circuit) uint16

type Circuit struct {
	state       map[string]uint16
	expressions map[string]Expression
}

func main() {
	circuit := InputToCircuit(2015, 7)
	fmt.Printf("a: %d\n", circuit.Read("a"))
}

func InputToCircuit(year, day int) *Circuit {
	expressions := make(map[string]Expression)
	for _, line := range aoc.InputToLines(year, day) {
		sides := strings.Split(line, " -> ")
		if len(sides) != 2 {
			log.Fatalf("unable to parse line: %s into sides", line)
		}

		expressions[sides[1]] = ParseExpression(sides[0])
	}

	return &Circuit{
		state:       make(map[string]uint16),
		expressions: expressions,
	}
}

func (c *Circuit) Read(wire string) uint16 {
	// Sometimes we ask to read a constant value.
	if n, err := strconv.Atoi(wire); err == nil {
		return uint16(n)
	}

	if value, ok := c.state[wire]; ok {
		return value
	}

	expr := c.expressions[wire]
	value := expr(c)
	c.state[wire] = value
	return value
}

func ParseExpression(lhs string) Expression {
	terms := strings.Split(lhs, " ")

	if len(terms) == 3 {
		if terms[1] == "AND" {
			return func(c *Circuit) uint16 {
				lhs := c.Read(terms[0])
				rhs := c.Read(terms[2])
				return lhs & rhs
			}
		}

		if terms[1] == "OR" {
			return func(c *Circuit) uint16 {
				lhs := c.Read(terms[0])
				rhs := c.Read(terms[2])
				return lhs | rhs
			}
		}

		if terms[1] == "LSHIFT" {
			return func(c *Circuit) uint16 {
				lhs := c.Read(terms[0])
				rhs := c.Read(terms[2])
				return lhs << rhs
			}
		}

		if terms[1] == "RSHIFT" {
			return func(c *Circuit) uint16 {
				lhs := c.Read(terms[0])
				rhs := c.Read(terms[2])
				return lhs >> rhs
			}
		}
	}

	if len(terms) == 2 {
		if terms[0] == "NOT" {
			return func(c *Circuit) uint16 {
				rhs := c.Read(terms[1])
				return ^rhs
			}
		}
	}

	return func(c *Circuit) uint16 {
		return c.Read(terms[0])
	}
}
