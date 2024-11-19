package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/lib"
	"sort"
	"strings"
)

func main() {
	food := InputToFood()

	// Determine the possible ingredients for each allergen.  If an allergen's
	// set is empty then we haven't yet processed any foods containing it and
	// initialize it to this food's ingredients.  Otherwise, restrict the set of
	// ingredients that could contain the allergen by this food's ingredients.
	mapping := make(map[string]lib.Set[string])
	for _, f := range food {
		ingredients := lib.SetFrom(f.Ingredients...)

		for _, a := range f.Allergens {
			if len(mapping[a]) == 0 {
				mapping[a] = mapping[a].Union(ingredients)
			} else {
				mapping[a] = mapping[a].Intersect(ingredients)
			}
		}
	}

	assignments := make(map[string]string)
	for len(mapping) > 0 {
		for a, fs := range mapping {
			fs = fs.DifferenceElems(lib.GetMapValues(assignments)...)

			if len(fs) == 1 {
				food := fs.Entries()[0]

				assignments[a] = food
				delete(mapping, a)
			}
		}
	}

	allergens := lib.GetMapKeys(assignments)
	sort.Strings(allergens)

	var foods []string
	for _, a := range allergens {
		foods = append(foods, assignments[a])
	}
	fmt.Println(strings.Join(foods, ","))
}

type Food struct {
	Ingredients, Allergens []string
}

func InputToFood() []Food {
	return lib.InputLinesTo(func(line string) Food {
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
