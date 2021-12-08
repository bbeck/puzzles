package main

import (
	"errors"
	"fmt"
	"github.com/bbeck/advent-of-code/aoc"
	"log"
	"sort"
	"strings"
)

var digits = map[int]string{
	0: "ABCEFG",  // 6
	1: "CF",      // 2
	2: "ACDEG",   // 5
	3: "ACDFG",   // 5
	4: "BCDF",    // 4
	5: "ABDFG",   // 5
	6: "ABDEFG",  // 6
	7: "ACF",     // 3
	8: "ABCDEFG", // 7
	9: "ABCDFG",  // 6
}

func main() {
	entries := InputToEntries()

	var sum int
	for _, entry := range entries {
		mapping := Solve(entry)

		var n int
		for _, output := range entry.Outputs {
			d, _ := Eval(output, mapping)
			n = n*10 + d
		}

		sum += n
	}
	fmt.Println(sum)
}

var alphabet = []string{"A", "B", "C", "D", "E", "F", "G"}

func Solve(entry Entry) map[string]string {
	var solution map[string]string
	aoc.EnumeratePermutations(7, func(perm []int) {
		mapping := map[string]string{
			"a": alphabet[perm[0]],
			"b": alphabet[perm[1]],
			"c": alphabet[perm[2]],
			"d": alphabet[perm[3]],
			"e": alphabet[perm[4]],
			"f": alphabet[perm[5]],
			"g": alphabet[perm[6]],
		}

		if Works(entry, mapping) {
			solution = mapping
		}
	})

	return solution
}

func Works(entry Entry, mapping map[string]string) bool {
	var numbers []string
	numbers = append(numbers, entry.Signals...)
	numbers = append(numbers, entry.Outputs...)

	for _, number := range numbers {
		_, err := Eval(number, mapping)
		if err != nil {
			return false
		}
	}

	return true
}

func Eval(s string, mapping map[string]string) (int, error) {
	var segments []string
	for _, c := range s {
		segments = append(segments, mapping[string(c)])
	}
	sort.Strings(segments)

	key := strings.Join(segments, "")
	for n := 0; n < 10; n++ {
		if digits[n] == key {
			return n, nil
		}
	}

	return -1, errors.New("mapping doesn't work")
}

type Entry struct {
	Signals []string
	Outputs []string
}

func InputToEntries() []Entry {
	var entries []Entry
	for _, line := range aoc.InputToLines(2021, 8) {
		var s0, s1, s2, s3, s4, s5, s6, s7, s8, s9 string
		var o0, o1, o2, o3 string
		_, err := fmt.Sscanf(line, "%s %s %s %s %s %s %s %s %s %s | %s %s %s %s",
			&s0, &s1, &s2, &s3, &s4, &s5, &s6, &s7, &s8, &s9, &o0, &o1, &o2, &o3)
		if err != nil {
			log.Fatal(err)
		}

		entries = append(entries, Entry{
			Signals: []string{s0, s1, s2, s3, s4, s5, s6, s7, s8, s9},
			Outputs: []string{o0, o1, o2, o3},
		})
	}
	return entries
}
