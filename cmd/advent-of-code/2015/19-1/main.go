package main

import (
	"fmt"
	"strings"

	. "github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	replacements, molecule := InputToReplacementsAndMolecule()

	var molecules Set[string]
	for start := range len(molecule) {
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

func InputToReplacementsAndMolecule() (map[string][]string, string) {
	replacements := make(map[string][]string)

	chunk := in.ChunkS()
	for chunk.HasNext() {
		lhs, rhs := chunk.Cut(" => ")
		replacements[lhs] = append(replacements[lhs], rhs)
	}

	molecules := in.Line()

	return replacements, molecules
}
