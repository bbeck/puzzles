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

	var root *Program
	for _, program := range programs {
		if !children[program.id] {
			root = program
			break
		}
	}

	// Walk down the tree constantly moving down the branch of imbalance until we
	// reach a node that has all equal weight children.  At that point we know
	// that we've found the node with an incorrect weight.
	var lastDelta int
	for root != nil {
		// choose the child that's out of balance
		choice, delta := Choose(root.children)
		if choice == nil {
			fmt.Printf("%s weight should be %d\n", root.id, root.weight-lastDelta)
			break
		}

		root = choice
		lastDelta = delta
	}
}

func Choose(children []*Program) (*Program, int) {
	var weights []int
	frequency := make(map[int]int)
	for _, child := range children {
		weight := Weight(child)
		weights = append(weights, weight)
		frequency[weight]++
	}

	// If all of the children have the same weight or there are no children then
	// we can't choose a child.
	if len(frequency) == 0 {
		return nil, 0
	}

	// Now that we know the weights, we need to find the one that has the lowest
	// frequency.
	for i := 0; i < len(weights); i++ {
		weight := weights[i]
		if frequency[weight] == 1 {
			if i != 0 {
				return children[i], Weight(children[i]) - Weight(children[0])
			} else {
				return children[i], Weight(children[i]) - Weight(children[1])
			}
		}
	}

	// If we get here then we know that all of the children had the same weight.
	return nil, 0
}

// Weight computes the weight of a program and all of its children
func Weight(p *Program) int {
	weight := p.weight
	for _, child := range p.children {
		weight += Weight(child)
	}

	return weight
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
