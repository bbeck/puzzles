package main

import (
	"fmt"
	"strings"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	ingredients := InputToIngredients()

	var best int
	EnumeratePossibleRecipes(ingredients, func(amounts []int) {
		var capacity, durability, flavor, texture int
		for i := 0; i < len(amounts); i++ {
			capacity += amounts[i] * ingredients[i].Capacity
			durability += amounts[i] * ingredients[i].Durability
			flavor += amounts[i] * ingredients[i].Flavor
			texture += amounts[i] * ingredients[i].Texture
		}

		if capacity > 0 && durability > 0 && flavor > 0 && texture > 0 {
			best = aoc.Max(best, capacity*durability*flavor*texture)
		}
	})

	fmt.Println(best)
}

func EnumeratePossibleRecipes(ingredients []Ingredient, fn func(amounts []int)) {
	amounts := make([]int, len(ingredients))

	var helper func(index int, remaining int)
	helper = func(index int, remaining int) {
		if index == len(amounts)-1 {
			amounts[index] = remaining
			fn(amounts)
			return
		}

		for amount := 0; amount <= remaining; amount++ {
			amounts[index] = amount
			helper(index+1, remaining-amount)
		}
	}

	helper(0, 100)
}

type Ingredient struct {
	Name       string
	Capacity   int
	Durability int
	Flavor     int
	Texture    int
	Calories   int
}

func InputToIngredients() []Ingredient {
	return aoc.InputLinesTo(2015, 15, func(line string) (Ingredient, error) {
		line = strings.ReplaceAll(line, ",", "")
		line = strings.ReplaceAll(line, "capacity", "")
		line = strings.ReplaceAll(line, "durability", "")
		line = strings.ReplaceAll(line, "flavor", "")
		line = strings.ReplaceAll(line, "texture", "")
		line = strings.ReplaceAll(line, "calories", "")

		var ingredient Ingredient
		_, err := fmt.Sscanf(line,
			"%s %d %d %d %d %d",
			&ingredient.Name,
			&ingredient.Capacity,
			&ingredient.Durability,
			&ingredient.Flavor,
			&ingredient.Texture,
			&ingredient.Calories,
		)
		return ingredient, err
	})
}
