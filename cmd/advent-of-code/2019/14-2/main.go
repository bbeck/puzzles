package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/bbeck/advent-of-code/lib"
)

func main() {
	reactions := make(map[string]Reaction)
	for _, reaction := range InputToReactions() {
		reactions[reaction.Output.Symbol] = reaction
	}

	fuel := sort.Search(1e12, func(fuel int) bool {
		chemicals := map[string]int{"FUEL": fuel}
		Reduce(chemicals, reactions)
		return chemicals["ORE"] > 1e12
	})

	// The search returns the first fuel amount that requires more than 1e12 ore.
	fmt.Println(fuel - 1)
}

func Reduce(chemicals map[string]int, reactions map[string]Reaction) {
	changed := true
	for changed {
		changed = false

		for _, reaction := range reactions {
			if chemicals[reaction.Output.Symbol] <= 0 {
				continue
			}

			multiplier := lib.Max(1, chemicals[reaction.Output.Symbol]/reaction.Output.Quantity)
			chemicals[reaction.Output.Symbol] -= multiplier * reaction.Output.Quantity
			for _, input := range reaction.Inputs {
				chemicals[input.Symbol] += multiplier * input.Quantity
			}

			changed = true
		}
	}
}

type Chemical struct {
	Symbol   string
	Quantity int
}

type Reaction struct {
	Inputs []Chemical
	Output Chemical
}

func InputToReactions() []Reaction {
	return lib.InputLinesTo(func(line string) Reaction {
		lhs, rhs, _ := strings.Cut(line, " => ")

		var reaction Reaction
		for _, s := range strings.Split(lhs, ", ") {
			quantity, symbol, _ := strings.Cut(s, " ")
			reaction.Inputs = append(reaction.Inputs, Chemical{
				Symbol:   symbol,
				Quantity: lib.ParseInt(quantity),
			})
		}

		quantity, symbol, _ := strings.Cut(rhs, " ")
		reaction.Output.Symbol = symbol
		reaction.Output.Quantity = lib.ParseInt(quantity)

		return reaction
	})
}
