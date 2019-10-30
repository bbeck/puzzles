package main

import (
	"fmt"
	"log"
	"math"

	"github.com/bbeck/advent-of-code/aoc"
)

type Ingredient struct {
	name       string
	capacity   int
	durability int
	flavor     int
	texture    int
	calories   int
}

func main() {
	ingredients := InputToIngredients(2015, 15)
	score := Optimize(ingredients)
	fmt.Printf("score: %d\n", score)
}

func InputToIngredients(year, day int) []Ingredient {
	var ingredients []Ingredient
	for _, line := range aoc.InputToLines(year, day) {
		var name string
		var capacity, durability, flavor, texture, calories int

		if _, err := fmt.Sscanf(line, "%s capacity %d, durability %d, flavor %d, texture %d, calories %d", &name, &capacity, &durability, &flavor, &texture, &calories); err != nil {
			log.Fatalf("unable to parse ingredient: %s", line)
		}

		ingredients = append(ingredients, Ingredient{
			name:       name[:len(name)-1], // name has a trailing colon
			capacity:   capacity,
			durability: durability,
			flavor:     flavor,
			texture:    texture,
			calories:   calories,
		})
	}

	return ingredients
}

func Optimize(ingredients []Ingredient) int {
	var helper func(remaining int, index int, amounts []int, ingredients []Ingredient) int
	helper = func(remaining int, index int, amounts []int, ingredients []Ingredient) int {
		if index >= len(amounts) {
			if remaining == 0 {
				return Score(amounts, ingredients)
			}

			// We chose all of the ingredients but still need more, so this is a
			// non-solution.
			return math.MinInt64
		}

		best := math.MinInt64
		for amount := 0; amount <= remaining; amount++ {
			amounts[index] = amount
			score := helper(remaining-amount, index+1, amounts, ingredients)
			if score > best {
				best = score
			}
		}

		return best
	}

	return helper(100, 0, make([]int, len(ingredients)), ingredients)
}

func Score(amounts []int, ingredients []Ingredient) int {
	var scoreCapacity int
	var scoreDurability int
	var scoreFlavor int
	var scoreTexture int
	for i := 0; i < len(amounts); i++ {
		scoreCapacity += amounts[i] * ingredients[i].capacity
		scoreDurability += amounts[i] * ingredients[i].durability
		scoreFlavor += amounts[i] * ingredients[i].flavor
		scoreTexture += amounts[i] * ingredients[i].texture
	}

	if scoreCapacity < 0 || scoreDurability < 0 || scoreFlavor < 0 || scoreTexture < 0 {
		return 0
	}

	return scoreCapacity * scoreDurability * scoreFlavor * scoreTexture
}
