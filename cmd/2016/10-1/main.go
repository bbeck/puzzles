package main

import (
	"fmt"
	"log"

	"github.com/bbeck/advent-of-code/aoc"
)

const (
	target1 = 17
	target2 = 61
)

func main() {
	inits, bots := InputToBots(2016, 10)
	for _, initialization := range inits {
		bots[initialization.bot].Take(initialization.value, bots)
	}
}

type Initialization struct {
	// value # goes to bot #
	value int
	bot   int
}

type Bot struct {
	// bot # gives low to (bot|output) # and high to (bot|output) #
	id                int
	lowType, highType string
	low, high         int
	value             int
}

func (b *Bot) Take(n int, bots map[int]*Bot) {
	if b.value == 0 {
		b.value = n
		return
	}

	if (n == target1 && b.value == target2) || (n == target2 && b.value == target1) {
		fmt.Printf("id: %d, comparing %d and %d\n", b.id, target1, target2)
	}

	var low, high int
	if n < b.value {
		low = n
		high = b.value
	} else {
		low = b.value
		high = n
	}
	b.value = 0

	if b.lowType == "bot" {
		bots[b.low].Take(low, bots)
	}

	if b.highType == "bot" {
		bots[b.high].Take(high, bots)
	}
}

func InputToBots(year, day int) ([]Initialization, map[int]*Bot) {
	var initializations []Initialization
	bots := make(map[int]*Bot)

	for _, line := range aoc.InputToLines(year, day) {
		var value, bot int
		if _, err := fmt.Sscanf(line, "value %d goes to bot %d", &value, &bot); err == nil {
			initializations = append(initializations, Initialization{value, bot})
			continue
		}

		var low, high int
		var lowType, highType string
		if _, err := fmt.Sscanf(line, "bot %d gives low to %s %d and high to %s %d", &bot, &lowType, &low, &highType, &high); err == nil {
			bots[bot] = &Bot{bot, lowType, highType, low, high, 0}
			continue
		}

		log.Fatalf("unable to parse instruction: %s", line)
	}

	return initializations, bots
}
