package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/aoc"
	"strings"
)

func main() {
	var points int
	for _, card := range InputToCards() {
		common := len(card.Winning.Intersect(card.Numbers))
		if common == 0 {
			continue
		}

		points += aoc.Pow(2, uint(common-1))
	}

	fmt.Println(points)
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
