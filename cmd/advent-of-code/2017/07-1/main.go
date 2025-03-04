package main

import (
	"fmt"
	. "github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	programs := InputToPrograms()

	// The program that's the root of the tree is the one that never appears as
	// a child of any other program.
	var all, children Set[string]
	for _, program := range programs {
		all.Add(program.ID)
		children.Add(program.Children...)
	}

	root := all.Difference(children).Entries()[0]
	fmt.Println(root)
}

type Program struct {
	ID       string
	Weight   int
	Children []string
}

func InputToPrograms() []Program {
	return in.LinesToS(func(in in.Scanner[Program]) Program {
		in.Remove("(", ")", "->", ",")

		var fields = in.Fields()
		return Program{
			ID:       fields[0],
			Weight:   ParseInt(fields[1]),
			Children: fields[2:],
		}
	})
}
