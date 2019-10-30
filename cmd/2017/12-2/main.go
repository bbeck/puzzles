package main

import (
	"fmt"
	"strings"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	programs := InputToPrograms(2017, 12)

	sets := make(map[string]*aoc.DisjointSet)
	for _, program := range programs {
		sets[program.ID()] = aoc.NewDisjointSet(program)
	}

	// Union the sets together based on the children of the programs
	for _, program := range programs {
		for _, child := range program.children {
			sets[program.ID()].Union(sets[child.ID()])
		}
	}

	// Now count the number of unique representative programs to determine how
	// many groups there are
	groups := make(map[string]bool)
	for _, set := range sets {
		groups[set.Find().Data.(*Program).ID()] = true
	}

	fmt.Printf("number of groups: %d\n", len(groups))
}

type Program struct {
	id       int
	children []*Program
}

func (p *Program) ID() string {
	return fmt.Sprintf("%d", p.id)
}

func InputToPrograms(year, day int) []*Program {
	connections := make(map[int][]int)
	for _, line := range aoc.InputToLines(year, day) {
		sides := strings.Split(line, " <-> ")
		a := aoc.ParseInt(sides[0])

		for _, rhs := range strings.Split(strings.ReplaceAll(sides[1], ",", ""), " ") {
			b := aoc.ParseInt(rhs)

			connections[a] = append(connections[a], b)
		}
	}

	programs := make(map[int]*Program)

	get := func(id int) *Program {
		program := programs[id]
		if program == nil {
			program = &Program{id: id}
			programs[id] = program
		}

		return program
	}

	for id, children := range connections {
		parent := get(id)

		for _, cid := range children {
			child := get(cid)
			parent.children = append(parent.children, child)
		}
	}

	ps := make([]*Program, 0)
	for _, program := range programs {
		ps = append(ps, program)
	}

	return ps
}
