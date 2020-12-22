package main

import (
	"crypto/sha256"
	"fmt"
	"strings"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	d1, d2 := InputToDecks(2020, 22)
	_, score := Play(d1, d2)
	fmt.Println(score)
}

func Play(d1, d2 Deck) (winner int, score int) {
	seen := aoc.NewSet()

	for !d1.Empty() && !d2.Empty() {
		// Check for infinite recursion
		hash := d1.Hash() + d2.Hash()
		if seen.Contains(hash) {
			return 1, -1
		}
		seen.Add(hash)

		c1, c2 := d1.Deal(), d2.Deal()
		switch {
		case int(c1) <= len(d1) && int(c2) <= len(d2):
			// We have enough cards in each deck to recurse, so start a new game to
			// determine the winner.
			if winner, _ := Play(d1.Copy(c1), d2.Copy(c2)); winner == 1 {
				d1.Add(c1, c2)
			} else {
				d2.Add(c2, c1)
			}

		case c1 < c2:
			d2.Add(c2, c1)

		case c2 < c1:
			d1.Add(c1, c2)
		}
	}

	if d1.Empty() {
		return 2, d2.Score()
	} else {
		return 1, d1.Score()
	}
}

type Deck []byte

func (d Deck) Empty() bool {
	return len(d) == 0
}

func (d Deck) Hash() string {
	return fmt.Sprintf("%x", sha256.Sum256(d))
}

func (d Deck) Score() int {
	N := len(d)

	var score int
	for i, c := range d {
		score += (N - i) * int(c)
	}
	return score
}

func (d *Deck) Deal() byte {
	c := (*d)[0]
	*d = (*d)[1:]
	return c
}

func (d *Deck) Add(cards ...byte) {
	*d = append(*d, cards...)
}

func (d Deck) Copy(n byte) Deck {
	return append([]byte(nil), d[:n]...)
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

		decks[current].Add(byte(aoc.ParseInt(line)))
	}

	return decks[0], decks[1]
}
