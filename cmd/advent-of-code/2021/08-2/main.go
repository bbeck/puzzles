package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/puz"
	"strings"
)

var StandardDigits = map[Digit]int{
	DigitFrom("abcefg"):  0,
	DigitFrom("cf"):      1,
	DigitFrom("acdeg"):   2,
	DigitFrom("acdfg"):   3,
	DigitFrom("bcdf"):    4,
	DigitFrom("abdfg"):   5,
	DigitFrom("abdefg"):  6,
	DigitFrom("acf"):     7,
	DigitFrom("abcdefg"): 8,
	DigitFrom("abcdfg"):  9,
}

func main() {
	// Build a list of all possible mappers
	var mappers []Mapper
	puz.EnumeratePermutations(7, func(perm []int) bool {
		p := append([]int(nil), perm...)

		mappers = append(mappers, func(from Digit) Digit {
			var to puz.BitSet
			for i, n := range p {
				if from.Contains(i) {
					to = to.Add(n)
				}
			}
			return Digit{to}
		})
		return false
	})

	var sum int
	for _, entry := range InputToEntries() {
		mapper := FindMapper(mappers, entry.Digits)

		var ns []int
		for _, digit := range entry.Outputs {
			ns = append(ns, StandardDigits[mapper(digit)])
		}
		sum += puz.JoinDigits(ns)
	}

	fmt.Println(sum)
}

func FindMapper(mappers []Mapper, digits []Digit) Mapper {
outer:
	for _, mapper := range mappers {
		for _, digit := range digits {
			if _, ok := StandardDigits[mapper(digit)]; !ok {
				continue outer
			}
		}
		return mapper
	}

	return nil
}

type Mapper func(Digit) Digit
type Digit struct{ puz.BitSet }

func DigitFrom(s string) Digit {
	var bs puz.BitSet
	for _, c := range s {
		bs = bs.Add(int(c - 'a'))
	}
	return Digit{bs}
}

func DigitsFrom(s string) []Digit {
	var digits []Digit
	for _, d := range strings.Fields(s) {
		digits = append(digits, DigitFrom(d))
	}
	return digits
}

type Entry struct {
	Digits  []Digit // All digits, both inputs and outputs
	Outputs []Digit // Only the outputs
}

func InputToEntries() []Entry {
	return puz.InputLinesTo(func(line string) Entry {
		lhs, rhs, _ := strings.Cut(line, " | ")
		return Entry{Digits: DigitsFrom(lhs + " " + rhs), Outputs: DigitsFrom(rhs)}
	})
}
