package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/aoc"
	"strings"
)

func main() {
	counts := make(map[int]int)
	for _, card := range InputToCards() {
		counts[card.ID]++

		common := len(card.Winning.Intersect(card.Numbers))
		for i := 1; i <= common; i++ {
			counts[card.ID+i] += counts[card.ID]
		}
	}

	var total int
	for _, count := range counts {
		total += count
	}
	fmt.Println(total)
}

type Card struct {
	ID      int
	Winning aoc.Set[int]
	Numbers aoc.Set[int]
}

func InputToCards() []Card {
	return aoc.InputLinesTo(2023, 4, func(line string) (Card, error) {
		line = strings.ReplaceAll(line, "Card ", "")
		line = strings.ReplaceAll(line, ":", "")

		fields := aoc.DequeFrom(strings.Fields(line)...)
		id := aoc.ParseInt(fields.PopFront())

		var winning aoc.Set[int]
		for field := fields.PopFront(); field != "|"; field = fields.PopFront() {
			winning.Add(aoc.ParseInt(field))
		}

		var numbers aoc.Set[int]
		for !fields.Empty() {
			numbers.Add(aoc.ParseInt(fields.PopFront()))
		}

		return Card{ID: id, Winning: winning, Numbers: numbers}, nil
	})
}
