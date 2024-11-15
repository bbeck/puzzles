package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/puz"
	"strings"
)

func main() {
	food := InputToFood()

	// Determine the possible ingredients for each allergen.  If an allergen's
	// set is empty then we haven't yet processed any foods containing it and
	// initialize it to this food's ingredients.  Otherwise, restrict the set of
	// ingredients that could contain the allergen by this food's ingredients.
	mapping := make(map[string]puz.Set[string])
	for _, f := range food {
		ingredients := puz.SetFrom(f.Ingredients...)

		for _, a := range f.Allergens {
			if len(mapping[a]) == 0 {
				mapping[a] = mapping[a].Union(ingredients)
			} else {
				mapping[a] = mapping[a].Intersect(ingredients)
			}
		}
	}

	// Build a set of possible allergens.
	var possible puz.Set[string]
	for _, fs := range mapping {
		possible = possible.Union(fs)
	}

	var count int
	for _, f := range food {
		for _, i := range f.Ingredients {
			if !possible.Contains(i) {
				count++
			}
		}
	}
	fmt.Println(count)
}

type Food struct {
	Ingredients, Allergens []string
}

func InputToFood() []Food {
	return puz.InputLinesTo(func(line string) Food {
		line = strings.ReplaceAll(line, "(", "")
		line = strings.ReplaceAll(line, ")", "")
		line = strings.ReplaceAll(line, ",", "")
		lhs, rhs, _ := strings.Cut(line, " contains ")

		return Food{
			Ingredients: strings.Fields(lhs),
			Allergens:   strings.Fields(rhs),
		}
	})
}
