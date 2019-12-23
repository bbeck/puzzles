package main

import (
	"fmt"
	"log"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	var deck []int
	N := 10007
	for i := 0; i < N; i++ {
		deck = append(deck, i)
	}

	instructions := InputToInstructions(2019, 22)
	for _, instruction := range instructions {
		switch instruction.kind {
		case DealNewStackKind:
			deck = DealNewStack(deck)
		case DealWithIncrementKind:
			deck = DealWithIncrement(deck, instruction.arg)
		case CutKind:
			deck = Cut(deck, instruction.arg)
		}
	}

	for i, card := range deck {
		if card == 2019 {
			fmt.Println("position of card 2019:", i)
		}
	}
}

func DealNewStack(deck []int) []int {
	var next []int
	for _, card := range deck {
		next = append([]int{card}, next...)
	}
	return next
}

func DealWithIncrement(deck []int, increment int) []int {
	next := make([]int, len(deck))

	var index int
	for _, card := range deck {
		next[index] = card
		index = (index + increment) % len(deck)
	}
	return next
}

func Cut(deck []int, cut int) []int {
	if cut == 0 {
		return deck
	}

	var next []int
	if cut > 0 {
		next = append(deck[cut:], deck[:cut]...)
	}
	if cut < 0 {
		cut = -cut
		next = append(deck[len(deck)-cut:], deck[:len(deck)-cut]...)
	}

	return next
}

const (
	DealNewStackKind      string = "DealNewStack"
	CutKind               string = "Cut"
	DealWithIncrementKind string = "DealWithIncrement"
)

type Instruction struct {
	kind string
	arg  int
}

func InputToInstructions(year, day int) []Instruction {
	var instructions []Instruction
	for _, line := range aoc.InputToLines(year, day) {
		var argument int

		if _, err := fmt.Sscanf(line, "deal with increment %d", &argument); err == nil {
			instructions = append(instructions, Instruction{
				kind: DealWithIncrementKind,
				arg:  argument,
			})
			continue
		}

		if line == "deal into new stack" {
			instructions = append(instructions, Instruction{
				kind: DealNewStackKind,
			})
			continue
		}

		if _, err := fmt.Sscanf(line, "cut %d", &argument); err == nil {
			instructions = append(instructions, Instruction{
				kind: CutKind,
				arg:  argument,
			})
			continue
		}

		log.Fatalf("unable to parse line: %s", line)
	}

	return instructions
}
