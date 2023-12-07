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
		ri, rj := Rank(hands[i].Text), Rank(hands[j].Text)
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

var Choices = []rune{
	'A', 'K', 'Q', 'T', '9', '8', '7', '6', '5', '4', '3', '2',
}

func Rank(h string) int {
	index := strings.Index(h, "J")
	if index == -1 {
		return RankInner(h)
	}

	var ranks []int
	for _, c := range Choices {
		ranks = append(ranks,
			Rank(fmt.Sprintf("%s%c%s", h[:index], c, h[index+1:])),
		)
	}
	return aoc.Max(ranks...)
}

func RankInner(h string) int {
	var fc aoc.FrequencyCounter[rune]
	for _, c := range h {
		fc.Add(c)
	}

	entries := fc.Entries()

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
	Text string
	Bid  int
}

func InputToHands() []Hand {
	return aoc.InputLinesTo(2023, 7, func(line string) (Hand, error) {
		lhs, rhs, _ := strings.Cut(line, " ")

		return Hand{
			Text: lhs,
			Bid:  aoc.ParseInt(rhs),
		}, nil
	})
}
