package main

import (
	"fmt"

	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	instructions := InputToInstructions()

	lroot, rroot := instructions[0].Left, instructions[0].Right
	for i := 1; i < len(instructions); i++ {
		Insert(lroot, instructions[i].Left)
		Insert(rroot, instructions[i].Right)
	}

	var lbest, rbest string
	for _, s := range Levels(lroot) {
		if len(s) > len(lbest) {
			lbest = s
		}
	}
	for _, s := range Levels(rroot) {
		if len(s) > len(rbest) {
			rbest = s
		}
	}

	fmt.Println(lbest + rbest)
}

func Insert(root *Node, leaf *Node) {
	for root != nil {
		if leaf.Rank <= root.Rank && root.Left != nil {
			root = root.Left
			continue
		}
		if leaf.Rank > root.Rank && root.Right != nil {
			root = root.Right
			continue
		}
		break
	}

	if leaf.Rank <= root.Rank {
		root.Left = leaf
	} else {
		root.Right = leaf
	}
}

func Levels(root *Node) []string {
	var m = make(map[int]string)

	var helper func(root *Node, level int)
	helper = func(root *Node, level int) {
		if root == nil {
			return
		}

		m[level] += root.Symbol
		helper(root.Left, level+1)
		helper(root.Right, level+1)
	}

	helper(root, 0)

	var levels = make([]string, len(m))
	for level, s := range m {
		levels[level] = s
	}
	return levels
}

type Instruction struct {
	Kind        string
	ID          int
	Left, Right *Node
}

type Node struct {
	Rank        int
	Symbol      string
	Left, Right *Node
}

func InputToInstructions() []Instruction {
	in.Remove("[", "]", ",")

	return in.LinesToS(func(in in.Scanner[Instruction]) Instruction {
		var instruction Instruction
		instruction.Kind = in.String()
		instruction.ID = in.Int()
		instruction.Left = &Node{Rank: in.Int(), Symbol: in.String()}
		instruction.Right = &Node{Rank: in.Int(), Symbol: in.String()}
		return instruction
	})
}
