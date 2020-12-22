package main

import (
	"fmt"
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

	// The set of foods that may have allergens
	suspects := aoc.NewSet()
	for _, foods := range causes {
		suspects = suspects.Union(foods)
	}

	// Count how many non-allergen foods there are in the ingredients lists
	var sum int
	for _, ingredient := range ingredients {
		for _, food := range ingredient.foods.Entries() {
			if !suspects.Contains(food) {
				sum++
			}
		}
	}
	fmt.Println(sum)
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
