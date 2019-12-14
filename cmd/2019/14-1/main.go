package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	reactions := make(map[string]Reaction)
	for _, reaction := range InputToReactions(2019, 14) {
		reactions[reaction.output] = reaction
	}

	// These are the chemicals we have and the quantities that we have them in,
	// we want to keep breaking these down until the only positive quantities that
	// we have is for ORE.
	chemicals := map[string]int{
		"FUEL": 1,
	}

	// Choose a chemical to break down, if there are none left then the empty
	// string is returned.
	choose := func(chemicals map[string]int) string {
		for chemical, quantity := range chemicals {
			if chemical != "ORE" && quantity > 0 {
				return chemical
			}
		}

		return ""
	}

	for {
		chemical := choose(chemicals)
		if chemical == "" {
			break
		}

		reaction := reactions[chemical]
		chemicals[chemical] -= reaction.quantity
		for input, quantity := range reaction.inputs {
			chemicals[input] += quantity
		}
	}

	fmt.Printf("ore needed: %d\n", chemicals["ORE"])
}

type Reaction struct {
	inputs   map[string]int
	output   string
	quantity int
}

func InputToReactions(year, day int) []Reaction {
	var reactions []Reaction
	for _, line := range aoc.InputToLines(year, day) {
		sides := strings.Split(line, " => ")

		inputs := make(map[string]int)
		for _, part := range strings.Split(sides[0], ", ") {
			var quantity int
			var chemical string
			if _, err := fmt.Sscanf(part, "%d %s", &quantity, &chemical); err != nil {
				log.Fatalf("unable to parse input part: %s", part)
			}

			inputs[chemical] = quantity
		}

		var quantity int
		var output string
		if _, err := fmt.Sscanf(sides[1], "%d %s", &quantity, &output); err != nil {
			log.Fatalf("unable to parse output part: %s", sides[1])
		}

		reactions = append(reactions, Reaction{inputs, output, quantity})
	}

	return reactions
}
