package main

import (
	"fmt"
	"strings"

	. "github.com/bbeck/puzzles/lib"
)

func main() {
	tree := InputToTree()

	paths := make(map[int][][]string)

	var walk func(node string, path []string)
	walk = func(node string, path []string) {
		if node == "ANT" || node == "BUG" {
			return
		}
		
		if node == "@" {
			p := make([]string, len(path))
			copy(p, path)

			paths[len(path)] = append(paths[len(path)], p)
			return
		}

		for child := range tree[node] {
			walk(child, append(path, string(child[0])))
		}
	}
	walk("RR", []string{"R"})

	for _, ps := range paths {
		if len(ps) == 1 {
			fmt.Println(strings.Join(ps[0], ""))
		}
	}
}

func InputToTree() map[string]Set[string] {
	tree := make(map[string]Set[string])
	for _, line := range InputToLines() {
		parent, rhs, _ := strings.Cut(line, ":")
		children := strings.Split(rhs, ",")
		tree[parent] = SetFrom(children...)
	}
	return tree
}
