package main

import (
	"fmt"
	"strings"

	"github.com/bbeck/advent-of-code/puz"
)

func main() {
	replacements := InputToReplacements()
	molecule := InputToMolecule()

	var molecules puz.Set[string]
	for start := 0; start < len(molecule); start++ {
		for prefix, values := range replacements {
			if !strings.HasPrefix(molecule[start:], prefix) {
				continue
			}

			for _, value := range values {
				updated := molecule[:start] + strings.Replace(molecule[start:], prefix, value, 1)
				molecules.Add(updated)
			}
		}
	}

	fmt.Println(len(molecules))
}

func InputToReplacements() map[string][]string {
	replacements := make(map[string][]string)
	for _, line := range puz.InputToLines() {
		lhs, rhs, ok := strings.Cut(line, " => ")
		if ok {
			replacements[lhs] = append(replacements[lhs], rhs)
		}
	}

	return replacements
}

func InputToMolecule() string {
	for _, line := range puz.InputToLines() {
		if len(line) > 0 && !strings.Contains(line, "=>") {
			return line
		}
	}

	return ""
}
