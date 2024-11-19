package main

import (
	"fmt"
	"strings"

	"github.com/bbeck/advent-of-code/lib"
)

func main() {
	var count int
	for _, line := range lib.InputToLines() {
		if IsNice(line) {
			count++
		}
	}

	fmt.Println(count)
}

func IsNice(s string) bool {
	return ContainsRepeatPair(s) && ContainsSplitRepeat(s)
}

func ContainsRepeatPair(s string) bool {
	for i := 0; i < len(s)-3; i++ {
		if strings.Contains(s[i+2:], s[i:i+2]) {
			return true
		}
	}

	return false
}

func ContainsSplitRepeat(s string) bool {
	for i := 0; i < len(s)-2; i++ {
		if s[i] == s[i+2] {
			return true
		}
	}

	return false
}
