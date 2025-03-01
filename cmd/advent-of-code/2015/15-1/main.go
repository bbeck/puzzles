package main

import (
	"fmt"

	. "github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	ingredients := InputToIngredients()

	var best int
	EnumeratePossibleRecipes(ingredients, func(amounts []int) {
		var capacity, durability, flavor, texture int
		for i := range amounts {
			capacity += amounts[i] * ingredients[i].Capacity
			durability += amounts[i] * ingredients[i].Durability
			flavor += amounts[i] * ingredients[i].Flavor
			texture += amounts[i] * ingredients[i].Texture
		}

		if capacity > 0 && durability > 0 && flavor > 0 && texture > 0 {
			best = Max(best, capacity*durability*flavor*texture)
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
	return in.LinesTo(func(in *in.Scanner[Ingredient]) Ingredient {
		return Ingredient{
			Name:       in.String(),
			Capacity:   in.Int(),
			Durability: in.Int(),
			Flavor:     in.Int(),
			Texture:    in.Int(),
			Calories:   in.Int(),
		}
	})
}
