package main

import (
	"fmt"
	"strings"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	programs := InputToPrograms(2017, 12)

	var root *Program
	for _, program := range programs {
		if program.id == 0 {
			root = program
			break
		}
	}

	seen := make(map[int]bool)
	mark := func(node aoc.Node) bool {
		seen[node.(*Program).id] = true
		return false
	}

	aoc.BreadthFirstSearch(root, mark)
	fmt.Printf("size of group: %d\n", len(seen))
}

type Program struct {
	id       int
	children []*Program
}

func (p *Program) ID() string {
	return fmt.Sprintf("%d", p.id)
}

func (p *Program) Children() []aoc.Node {
	var children []aoc.Node
	for _, child := range p.children {
		children = append(children, child)
	}

	return children
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
