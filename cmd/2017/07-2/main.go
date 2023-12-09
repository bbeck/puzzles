package main

import (
	"fmt"
	"strings"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	// The program that's the root of the tree is the one that never appears as
	// a child of any other program.
	programs := make(map[string]Program)
	var all, children aoc.Set[string]
	for _, program := range InputToPrograms() {
		programs[program.ID] = program
		all.Add(program.ID)
		children.Add(program.Children...)
	}
	root := all.Difference(children).Entries()[0]

	_, weight := FindImbalance(root, 0, programs)
	fmt.Println(weight)
}

func FindImbalance(id string, expected int, programs map[string]Program) (string, int) {
	program := programs[id]

	// Index the weights of our children.
	weights := make(map[int][]string)
	for _, child := range program.Children {
		weight := Weight(child, programs)
		weights[weight] = append(weights[weight], child)
	}

	// If there was only a single weight discovered then our children are in
	// balance, the imbalance must lie in this node's weight.
	if len(weights) == 1 {
		return id, expected - (Weight(id, programs) - program.Weight)
	}

	// We know there was more than one weight discovered.  Find the outlier one
	// and seek the imbalance in that branch.  If we end up in a situation where
	// a node only has two children, and they have different weights, that is
	// okay, either child can be adjusted to achieve balance.
	var outlier string
	for weight, children := range weights {
		if len(children) == 1 {
			outlier = children[0]
		}
		if len(children) > 1 {
			expected = weight
		}
	}

	return FindImbalance(outlier, expected, programs)
}

func Weight(id string, programs map[string]Program) int {
	program := programs[id]

	weight := program.Weight
	for _, child := range program.Children {
		weight += Weight(child, programs)
	}

	return weight
}

type Program struct {
	ID       string
	Weight   int
	Children []string
}

func InputToPrograms() []Program {
	return aoc.InputLinesTo(2017, 7, func(line string) Program {
		line = strings.ReplaceAll(line, ",", "")
		line = strings.ReplaceAll(line, "->", "")
		line = strings.ReplaceAll(line, "(", "")
		line = strings.ReplaceAll(line, ")", "")

		fields := strings.Fields(line)
		return Program{
			ID:       fields[0],
			Weight:   aoc.ParseInt(fields[1]),
			Children: fields[2:],
		}
	})
}
