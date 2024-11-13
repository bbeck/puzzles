package main

import (
	"fmt"
	"strings"

	"github.com/bbeck/advent-of-code/puz"
)

func main() {
	instructions, nodes := InputToInstructionsAndGraph()

	lcm := 1
	for _, node := range nodes {
		if strings.HasSuffix(node.ID, "A") {
			lcm = puz.LCM(lcm, Length(instructions, node))
		}
	}

	fmt.Println(lcm)
}

func Length(instructions string, root *Node) int {
	var n int
	for n = 0; !strings.HasSuffix(root.ID, "Z"); n++ {
		instruction := instructions[n%len(instructions)]
		if instruction == 'L' {
			root = root.Left
		} else {
			root = root.Right
		}
	}

	return n
}

type Node struct {
	ID    string
	Left  *Node
	Right *Node
}

func InputToInstructionsAndGraph() (string, []*Node) {
	lines := puz.InputToLines(2023, 8)
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

	return instructions, puz.GetMapValues(nodes)
}
