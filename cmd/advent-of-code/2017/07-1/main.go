package main

import (
	"fmt"
	"strings"

	"github.com/bbeck/advent-of-code/puz"
)

func main() {
	programs := InputToPrograms()

	// The program that's the root of the tree is the one that never appears as
	// a child of any other program.
	var all, children puz.Set[string]
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
	return puz.InputLinesTo(2017, 7, func(line string) Program {
		line = strings.ReplaceAll(line, ",", "")
		line = strings.ReplaceAll(line, "->", "")
		line = strings.ReplaceAll(line, "(", "")
		line = strings.ReplaceAll(line, ")", "")

		fields := strings.Fields(line)
		return Program{
			ID:       fields[0],
			Weight:   puz.ParseInt(fields[1]),
			Children: fields[2:],
		}
	})
}
