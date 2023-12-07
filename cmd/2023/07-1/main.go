package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	hands := InputToHands()
	sort.Slice(hands, func(i, j int) bool {
		ri, rj := Rank(hands[i]), Rank(hands[j])
		if ri != rj {
			return ri < rj
		}

		for n := 0; n < 5; n++ {
			si, sj := CardStrengths[hands[i].Text[n]], CardStrengths[hands[j].Text[n]]
			if si != sj {
				return si < sj
			}
		}

		return false
	})

	var sum int
	for i, hand := range hands {
		sum += (i + 1) * hand.Bid
	}
	fmt.Println(sum)
}

var CardStrengths = map[uint8]int{
	'A': 12,
	'K': 11,
	'Q': 10,
	'J': 9,
	'T': 8,
	'9': 7,
	'8': 6,
	'7': 5,
	'6': 4,
	'5': 3,
	'4': 2,
	'3': 1,
	'2': 0,
}

func Rank(h Hand) int {
	entries := h.Cards.Entries()

	// 5 of a kind
	if entries[0].Count == 5 {
		return 7
	}

	// 4 of a kind
	if entries[0].Count == 4 {
		return 6
	}

	// full house
	if entries[0].Count == 3 && entries[1].Count == 2 {
		return 5
	}

	// 3 of a kind
	if entries[0].Count == 3 {
		return 4
	}

	// 2 pair
	if entries[0].Count == 2 && entries[1].Count == 2 {
		return 3
	}

	// 1 pair
	if entries[0].Count == 2 {
		return 2
	}

	return 1
}

type Hand struct {
	Text  string
	Cards aoc.FrequencyCounter[rune]
	Bid   int
}

func InputToHands() []Hand {
	return aoc.InputLinesTo(2023, 7, func(line string) (Hand, error) {
		lhs, rhs, _ := strings.Cut(line, " ")

		var fc aoc.FrequencyCounter[rune]
		for _, c := range lhs {
			fc.Add(c)
		}

		return Hand{
			Text:  lhs,
			Cards: fc,
			Bid:   aoc.ParseInt(rhs),
		}, nil
	})
}
