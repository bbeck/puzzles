package main

import (
	"fmt"
	"strings"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	programs := InputToPrograms(2017, 7)

	// The program that's the root of the tree is the one that never appears as
	// a child of any other program.
	children := make(map[string]bool)
	for _, program := range programs {
		for _, child := range program.children {
			children[child.id] = true
		}
	}

	for _, program := range programs {
		if !children[program.id] {
			fmt.Printf("root: %s\n", program.id)
			break
		}
	}
}

type Program struct {
	id       string
	weight   int
	children []*Program
}

func InputToPrograms(year, day int) []*Program {
	catalog := make(map[string]*Program)
	for _, line := range aoc.InputToLines(year, day) {
		line = strings.ReplaceAll(line, ",", "")
		line = strings.ReplaceAll(line, "(", "")
		line = strings.ReplaceAll(line, ")", "")

		tokens := strings.Split(line, " ")
		id := tokens[0]
		weight := aoc.ParseInt(tokens[1])

		program := catalog[id]
		if program == nil {
			program = &Program{id: id}
			catalog[id] = program
		}

		program.weight = weight

		// children
		for i := 3; i < len(tokens); i++ {
			child := catalog[tokens[i]]
			if child == nil {
				child = &Program{id: tokens[i]}
				catalog[tokens[i]] = child
			}

			program.children = append(program.children, child)
		}
	}

	var programs []*Program
	for _, program := range catalog {
		programs = append(programs, program)
	}

	return programs
}
