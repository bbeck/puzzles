package main

import (
	"fmt"
	"strings"

	"github.com/bbeck/advent-of-code/lib"
)

func main() {
	d1, d2 := InputToDecks()
	for !d1.Empty() && !d2.Empty() {
		c1, c2 := d1.PopFront(), d2.PopFront()
		if c1 > c2 {
			d1.PushBack(c1)
			d1.PushBack(c2)
		} else {
			d2.PushBack(c2)
			d2.PushBack(c1)
		}
	}

	var winner = d1
	if !d2.Empty() {
		winner = d2
	}

	var sum int
	for i, c := range winner.Entries() {
		sum += (winner.Len() - i) * c
	}
	fmt.Println(sum)
}

type Deck struct{ lib.Deque[int] }

func InputToDecks() (Deck, Deck) {
	var decks [2]lib.Deque[int]

	current := -1
	for _, line := range lib.InputToLines() {
		if len(line) == 0 {
			continue
		}

		if strings.HasPrefix(line, "Player") {
			current++
			continue
		}

		decks[current].PushBack(lib.ParseInt(line))
	}

	return Deck{decks[0]}, Deck{decks[1]}
}
