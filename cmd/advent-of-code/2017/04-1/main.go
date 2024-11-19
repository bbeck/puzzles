package main

import (
	"fmt"
	"strings"

	"github.com/bbeck/advent-of-code/lib"
)

func main() {
	var count int
	for _, line := range lib.InputToLines() {
		if IsValid(line) {
			count++
		}
	}
	fmt.Println(count)
}

func IsValid(s string) bool {
	var seen lib.Set[string]
	for _, word := range strings.Fields(s) {
		if !seen.Add(word) {
			return false
		}
	}

	return true
}
