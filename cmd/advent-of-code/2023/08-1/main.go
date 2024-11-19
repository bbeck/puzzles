package main

import (
	"fmt"
	"strings"

	"github.com/bbeck/advent-of-code/lib"
)

func main() {
	instructions, root := InputToInstructionsAndGraph()

	var n int
	for n = 0; root.ID != "ZZZ"; n++ {
		instruction := instructions[n%len(instructions)]
		if instruction == 'L' {
			root = root.Left
		} else {
			root = root.Right
		}
	}

	fmt.Println(n)
}

type Node struct {
	ID    string
	Left  *Node
	Right *Node
}

func InputToInstructionsAndGraph() (string, *Node) {
	lines := lib.InputToLines()
	instructions := lines[0]

	nodes := make(map[string]*Node)
	get := func(id string) *Node {
		if node := nodes[id]; node == nil {
			nodes[id] = &Node{ID: id}
		}
		return nodes[id]
	}

	for _, line := range lines[2:] {
		line = strings.ReplaceAll(line, "=", "")
		line = strings.ReplaceAll(line, "(", "")
		line = strings.ReplaceAll(line, ",", "")
		line = strings.ReplaceAll(line, ")", "")
		fields := strings.Fields(line)

		node := get(fields[0])
		node.Left = get(fields[1])
		node.Right = get(fields[2])
	}

	return instructions, nodes["AAA"]
}
