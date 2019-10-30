package main

import (
	"fmt"
	"log"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	generators := InputToGenerators(2017, 15)

	var count int
	for round := 0; round < 40_000_000; round++ {
		seen := make(map[uint64]bool)
		for _, g := range generators {
			seen[g.Next()&0xFFFF] = true
		}

		if len(seen) == 1 {
			count++
		}
	}

	fmt.Printf("count: %d\n", count)
}

type Generator struct {
	id     string
	factor uint64
	value  uint64
}

func (g *Generator) Next() uint64 {
	g.value = (g.value * g.factor) % 2147483647
	return g.value
}

func InputToGenerators(year, day int) []*Generator {
	var generators []*Generator
	for _, line := range aoc.InputToLines(year, day) {
		var id string
		var value uint64
		if _, err := fmt.Sscanf(line, "Generator %s starts with %d", &id, &value); err != nil {
			log.Fatalf("unable to parse input line: %s", line)
		}

		switch id {
		case "A":
			generators = append(generators, &Generator{id: id, factor: 16807, value: value})

		case "B":
			generators = append(generators, &Generator{id: id, factor: 48271, value: value})
		}
	}

	return generators
}
