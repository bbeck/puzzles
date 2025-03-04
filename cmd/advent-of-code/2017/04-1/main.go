package main

import (
	"fmt"
	"github.com/bbeck/puzzles/lib/in"
	"strings"

	. "github.com/bbeck/puzzles/lib"
)

func main() {
	var count int
	for in.HasNext() {
		if IsValid(in.Line()) {
			count++
		}
	}
	fmt.Println(count)
}

func IsValid(s string) bool {
	var seen Set[string]
	for _, word := range strings.Fields(s) {
		if !seen.Add(word) {
			return false
		}
	}

	return true
}
