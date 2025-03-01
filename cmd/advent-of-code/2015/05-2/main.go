package main

import (
	"fmt"
	"strings"

	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	var count int
	for in.HasNext() {
		if IsNice(in.String()) {
			count++
		}
	}

	fmt.Println(count)
}

func IsNice(s string) bool {
	return ContainsRepeatPair(s) && ContainsSplitRepeat(s)
}

func ContainsRepeatPair(s string) bool {
	for i := range len(s) - 3 {
		if strings.Contains(s[i+2:], s[i:i+2]) {
			return true
		}
	}

	return false
}

func ContainsSplitRepeat(s string) bool {
	for i := range len(s) - 2 {
		if s[i] == s[i+2] {
			return true
		}
	}

	return false
}
