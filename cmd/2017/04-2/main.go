package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	var count int
	for _, line := range aoc.InputToLines(2017, 4) {
		if IsValid(line) {
			count++
		}
	}

	fmt.Printf("count: %d\n", count)
}

func IsValid(line string) bool {
	seen := make(map[string]bool)

	for _, word := range strings.Split(line, " ") {
		word = Canonicalize(word)
		if seen[word] {
			return false
		}

		seen[word] = true
	}

	return true
}

func Canonicalize(s string) string {
	bs := []byte(s)
	sort.Slice(bs, func(i, j int) bool {
		return bs[i] < bs[j]
	})

	return string(bs)
}
