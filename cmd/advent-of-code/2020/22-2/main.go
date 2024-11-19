package main

import (
	"fmt"
	"strings"

	"github.com/bbeck/puzzles/lib"
)

func main() {
	d1, d2 := InputToDecks()
	_, deck := Play(d1, d2)

	var sum int
	for i, c := range deck.Entries() {
		sum += (deck.Deque.Len() - i) * int(c)
	}
	fmt.Println(sum)
}

func Play(d1, d2 Deck) (int, Deck) {
	var seen lib.Set[string]

	for !d1.Empty() && !d2.Empty() {
		if !seen.Add(d1.ID() + "|" + d2.ID()) {
			// Repeated rounds end immediately with a win for player 1
			return 1, d1
		}

		var winner int
		c1, c2 := d1.PopFront(), d2.PopFront()
		if c1 <= d1.Len() && c2 <= d2.Len() {
			winner, _ = Play(d1.Copy(c1), d2.Copy(c2))
		} else {
			// One or both players don't have enough cards to recurse
			if c1 > c2 {
				winner = 1
			} else {
				winner = 2
			}
		}

		if winner == 1 {
			d1.PushBack(c1)
			d1.PushBack(c2)
		} else {
			d2.PushBack(c2)
			d2.PushBack(c1)
		}
	}

	if d2.Empty() {
		return 1, d1
	} else {
		return 2, d2
	}
}

type Deck struct{ lib.Deque[byte] }

func (d Deck) Len() byte {
	return byte(d.Deque.Len())
}

func (d Deck) ID() string {
	var sb strings.Builder
	sb.Write(d.Entries())
	return sb.String()
}

func (d Deck) Copy(n byte) Deck {
	var c Deck
	for _, card := range d.Entries()[:n] {
		c.PushBack(card)
	}
	return c
}

func InputToDecks() (Deck, Deck) {
	var decks [2]lib.Deque[byte]

	current := -1
	for _, line := range lib.InputToLines() {
		if len(line) == 0 {
			continue
		}

		if strings.HasPrefix(line, "Player") {
			current++
			continue
		}

		card := byte(lib.ParseInt(line))
		decks[current].PushBack(card)
	}

	return Deck{decks[0]}, Deck{decks[1]}
}
