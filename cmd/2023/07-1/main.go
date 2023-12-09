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
		ti, tj := hands[i].Type, hands[j].Type
		if ti != tj {
			return ti < tj
		}

		for n := 0; n < 5; n++ {
			ci, cj := hands[i].Cards[n], hands[j].Cards[n]
			if ci != cj {
				return CardStrengths[ci] < CardStrengths[cj]
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

func Type(cards string) int {
	var fc aoc.FrequencyCounter[rune]
	for _, c := range cards {
		fc.Add(c)
	}
	entries := fc.Entries()

	if entries[0].Count == 5 {
		return 7
	}
	if entries[0].Count == 4 {
		return 6
	}
	if entries[0].Count == 3 && entries[1].Count == 2 {
		return 5
	}
	if entries[0].Count == 3 {
		return 4
	}
	if entries[0].Count == 2 && entries[1].Count == 2 {
		return 3
	}
	if entries[0].Count == 2 {
		return 2
	}
	return 1
}

type Hand struct {
	Cards string
	Bid   int
	Type  int
}

func InputToHands() []Hand {
	return aoc.InputLinesTo(2023, 7, func(line string) Hand {
		lhs, rhs, _ := strings.Cut(line, " ")

		return Hand{
			Cards: lhs,
			Bid:   aoc.ParseInt(rhs),
			Type:  Type(lhs),
		}
	})
}
