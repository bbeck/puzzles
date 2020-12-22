package main

import (
	"fmt"
	"strings"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	d1, d2 := InputToDecks(2020, 22)

	for !d1.Empty() && !d2.Empty() {
		c1, c2 := d1.Deal(), d2.Deal()

		switch {
		case c1 < c2:
			d2.Add(c2, c1)

		case c1 > c2:
			d1.Add(c1, c2)
		}
	}

	var deck Deck
	if d1.Empty() {
		deck = d2
	} else {
		deck = d1
	}
	fmt.Println(deck.Score())
}

type Deck []int

func (d Deck) Empty() bool {
	return len(d) == 0
}

func (d Deck) Score() int {
	N := len(d)

	var score int
	for i, c := range d {
		score += (N - i) * c
	}
	return score
}

func (d *Deck) Deal() int {
	c := (*d)[0]
	*d = (*d)[1:]
	return c
}

func (d *Deck) Add(cards ...int) {
	*d = append(*d, cards...)
}

func InputToDecks(year, day int) (Deck, Deck) {
	decks := []Deck{
		make(Deck, 0),
		make(Deck, 0),
	}

	current := -1
	for _, line := range aoc.InputToLines(year, day) {
		if len(line) == 0 {
			continue
		}

		if strings.HasPrefix(line, "Player") {
			current++
			continue
		}

		decks[current].Add(aoc.ParseInt(line))
	}

	return decks[0], decks[1]
}
