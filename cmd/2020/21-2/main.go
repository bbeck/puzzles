package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	ingredients := InputToIngredients(2020, 21)

	allFoods := aoc.NewSet()
	allAllergens := aoc.NewSet()
	for _, ingredient := range ingredients {
		allFoods = allFoods.Union(ingredient.foods)
		allAllergens = allAllergens.Union(ingredient.allergens)
	}

	// Mapping of allergen to the foods that could cause it.  Initially it's
	// all foods.
	causes := make(map[string]aoc.Set)
	for _, entry := range allAllergens.Entries() {
		causes[entry.(string)] = allFoods.Union(aoc.NewSet())
	}

	// Reduce the possible causes of each allergen by only the foods in each
	// ingredients list.
	for _, ingredient := range ingredients {
		for _, entry := range ingredient.allergens.Entries() {
			allergen := entry.(string)
			causes[allergen] = causes[allergen].Intersect(ingredient.foods)
		}
	}

	// Now that we have a mapping of possible foods for each allergen keep looping
	// through the mappings adding assignments for singleton sets and removing
	// assigned entries from the remaining sets.
	assigned := make([]string, 0)          // Foods
	assignments := make(map[string]string) // food -> allergen
	for len(assigned) < allAllergens.Size() {
		for allergen, foods := range causes {
			for _, food := range assigned {
				foods.Remove(food)
			}
			if foods.Size() == 1 {
				food := foods.Entries()[0].(string)
				assigned = append(assigned, food)
				assignments[food] = allergen
			}
		}
	}

	sort.Slice(assigned, func(i, j int) bool {
		return assignments[assigned[i]] < assignments[assigned[j]]
	})
	fmt.Println(strings.Join(assigned, ","))
}

type Ingredient struct {
	foods     aoc.Set
	allergens aoc.Set
}

func InputToIngredients(year, day int) []Ingredient {
	var ingredients []Ingredient
	for _, line := range aoc.InputToLines(year, day) {
		line = strings.ReplaceAll(line, "contains ", "")
		line = strings.ReplaceAll(line, ",", "")
		line = strings.ReplaceAll(line, ")", "")

		parts := strings.Split(line, " (")
		lhs := parts[0]
		rhs := parts[1]

		foods := aoc.NewSet()
		for _, food := range strings.Split(lhs, " ") {
			foods.Add(food)
		}

		allergens := aoc.NewSet()
		for _, allergen := range strings.Split(rhs, " ") {
			allergens.Add(allergen)
		}

		ingredients = append(ingredients, Ingredient{
			foods:     foods,
			allergens: allergens,
		})
	}

	return ingredients
}
