package main

import (
	"fmt"
	"strings"

	"github.com/bbeck/advent-of-code/puz"
)

func main() {
	reactions := make(map[string]Reaction)
	for _, reaction := range InputToReactions() {
		reactions[reaction.Output.Symbol] = reaction
	}

	chemicals := map[string]int{"FUEL": 1}
	Reduce(chemicals, reactions)
	fmt.Println(chemicals["ORE"])
}

func Reduce(chemicals map[string]int, reactions map[string]Reaction) {
	changed := true
	for changed {
		changed = false

		for _, reaction := range reactions {
			if chemicals[reaction.Output.Symbol] <= 0 {
				continue
			}

			multiplier := puz.Max(1, chemicals[reaction.Output.Symbol]/reaction.Output.Quantity)
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
	return puz.InputLinesTo(func(line string) Reaction {
		lhs, rhs, _ := strings.Cut(line, " => ")

		var reaction Reaction
		for _, s := range strings.Split(lhs, ", ") {
			quantity, symbol, _ := strings.Cut(s, " ")
			reaction.Inputs = append(reaction.Inputs, Chemical{
				Symbol:   symbol,
				Quantity: puz.ParseInt(quantity),
			})
		}

		quantity, symbol, _ := strings.Cut(rhs, " ")
		reaction.Output.Symbol = symbol
		reaction.Output.Quantity = puz.ParseInt(quantity)

		return reaction
	})
}
