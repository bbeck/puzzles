package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/aoc"
	"sort"
	"strings"
)

var Letters = []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g'}

type Mapping map[string]int

var StandardMapping = Mapping{
	"abcefg":  0,
	"cf":      1,
	"acdeg":   2,
	"acdfg":   3,
	"bcdf":    4,
	"abdfg":   5,
	"abdefg":  6,
	"acf":     7,
	"abcdefg": 8,
	"abcdfg":  9,
}

func main() {
	// We're going to brute force this by trying all possible mappings because there are only
	// 7! == 5040 possibilities.  We'll precompute all the mappings and then figure out which
	// should be used for each entry.
	var mappings []Mapping
	aoc.EnumeratePermutations(len(Letters), func(perm []int) {
		mapper := NewMapper(perm)

		mapping := make(Mapping)
		for s, n := range StandardMapping {
			mapping[Canonicalize(strings.Map(mapper, s))] = n
		}
		mappings = append(mappings, mapping)
	})

	// Now process each entry, summing the outputs
	var sum int
	for _, entry := range InputToEntries() {
		mapping := FindMapping(mappings, entry)

		var n int
		for _, output := range entry.Outputs {
			n = n*10 + mapping[output]
		}
		sum += n
	}
	fmt.Println(sum)
}

func NewMapper(perm []int) func(rune) rune {
	mapping := make(map[rune]rune)
	for from, to := range perm {
		mapping[Letters[from]] = Letters[to]
	}

	return func(r rune) rune {
		return mapping[r]
	}
}

func FindMapping(mappings []Mapping, entry Entry) Mapping {
	for _, mapping := range mappings {
		works := true
		for _, digit := range entry.Digits {
			if _, ok := mapping[digit]; !ok {
				works = false
				break
			}
		}

		if works {
			return mapping
		}
	}

	return nil
}

type Entry struct {
	Digits  []string // All digits, both inputs and outputs
	Outputs []string // Only the outputs
}

func InputToEntries() []Entry {
	var entries []Entry
	for _, line := range aoc.InputToLines(2021, 8) {
		fields := strings.Split(line, " | ")

		var inputs []string
		for _, s := range strings.Split(fields[0], " ") {
			inputs = append(inputs, Canonicalize(s))
		}

		var outputs []string
		for _, s := range strings.Split(fields[1], " ") {
			outputs = append(outputs, Canonicalize(s))
		}

		entries = append(entries, Entry{
			Digits:  append(inputs, outputs...),
			Outputs: outputs,
		})
	}
	return entries
}

func Canonicalize(s string) string {
	var letters []rune
	for _, r := range s {
		letters = append(letters, r)
	}

	sort.Slice(letters, func(i, j int) bool {
		return letters[i] < letters[j]
	})

	return string(letters)
}
