package main

import (
	"fmt"
	"strings"

	. "github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	var tree = InputToTree()

	var count int
	BreadthFirstSearch(
		"NULL",
		func(s string) []string { return tree[s] },
		func(s string) bool { count++; return false },
	)
	fmt.Println(count - 1) // Don't report NULL
}

func InputToTree() map[string][]string {
	var nodes = make(map[string][]string)
	for in.HasNext() {
		lhs, rhs := in.Cut(":")
		nodes[lhs] = strings.Split(rhs, ",")
	}
	return nodes
}
