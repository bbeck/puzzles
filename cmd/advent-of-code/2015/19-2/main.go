package main

import (
	"fmt"
	"math/rand"
	"strings"

	"github.com/bbeck/puzzles/lib"
)

func main() {
	replacements := InputToReplacements()
	molecule := InputToMolecule()

	// Keep attempting to find a path, if we fail we'll reorder the replacements
	// in the hope that trying a different order will succeed.
	var count int
	for {
		if steps, ok := Synthesize(molecule, replacements); ok {
			count = steps
			break
		}

		rand.Shuffle(len(replacements), func(i, j int) {
			replacements[i], replacements[j] = replacements[j], replacements[i]
		})
	}

	fmt.Println(count)
}

func Synthesize(molecule string, replacements []Replacement) (int, bool) {
	// This technically won't work for every input, but we're assuming that
	// there's only one solution to get from an "e" to our molecule.  Instead
	// of trying to build up to the molecule in the forward direction we'll
	// move backwards from our molecule to the "e".  If we reach a dead end
	// we'll fail and the caller may try again.
	var count int
	for molecule != "e" {
		start := molecule

		for _, replacement := range replacements {
			updated := strings.Replace(molecule, replacement.RHS, replacement.LHS, 1)
			if updated != molecule {
				count++
				molecule = updated
				break
			}
		}

		if start == molecule {
			// None of our replacements worked, return a failure
			return 0, false
		}
	}

	return count, true
}

type Replacement struct {
	LHS, RHS string
}

func InputToReplacements() []Replacement {
	var replacements []Replacement
	for _, line := range lib.InputToLines() {
		lhs, rhs, ok := strings.Cut(line, " => ")
		if ok {
			replacements = append(replacements, Replacement{LHS: lhs, RHS: rhs})
		}
	}

	return replacements
}

func InputToMolecule() string {
	for _, line := range lib.InputToLines() {
		if len(line) > 0 && !strings.Contains(line, "=>") {
			return line
		}
	}

	return ""
}
