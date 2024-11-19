package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/bbeck/puzzles/lib"
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
	'T': 9,
	'9': 8,
	'8': 7,
	'7': 6,
	'6': 5,
	'5': 4,
	'4': 3,
	'3': 2,
	'2': 1,
	'J': 0,
}

func Type(cards string) int {
	c0, c1, j := Counts(cards)

	// The strongest hand is always applying the jokers to the most frequently
	// occurring card.
	c0 += j

	if c0 == 5 {
		return 7
	}
	if c0 == 4 {
		return 6
	}
	if c0 == 3 && c1 == 2 {
		return 5
	}
	if c0 == 3 {
		return 4
	}
	if c0 == 2 && c1 == 2 {
		return 3
	}
	if c0 == 2 {
		return 2
	}
	return 1
}

func Counts(cards string) (int, int, int) {
	// Count the cards and return the two most frequent counts as well as the
	// number of jokers.
	var fc lib.FrequencyCounter[rune]
	for _, c := range cards {
		fc.Add(c)
	}

	var numJokers int
	var counts []int
	for _, entry := range fc.Entries() {
		if entry.Value == 'J' {
			numJokers = entry.Count
			continue
		}
		counts = append(counts, entry.Count)
	}
	for len(counts) < 2 {
		counts = append(counts, 0)
	}

	return counts[0], counts[1], numJokers
}

type Hand struct {
	Cards string
	Bid   int
	Type  int
}

func InputToHands() []Hand {
	return lib.InputLinesTo(func(line string) Hand {
		lhs, rhs, _ := strings.Cut(line, " ")

		return Hand{
			Cards: lhs,
			Bid:   lib.ParseInt(rhs),
			Type:  Type(lhs),
		}
	})
}
