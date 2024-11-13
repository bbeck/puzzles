package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/bbeck/advent-of-code/puz"
)

func main() {
	var count int
	for _, line := range puz.InputToLines(2017, 4) {
		if IsValid(line) {
			count++
		}
	}
	fmt.Println(count)
}

func IsValid(s string) bool {
	var seen puz.Set[string]
	for _, word := range strings.Fields(s) {
		if !seen.Add(Canonicalize(word)) {
			return false
		}
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
