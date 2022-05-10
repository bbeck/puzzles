package main

import (
	"fmt"
	"strings"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	var count int
	for _, line := range aoc.InputToLines(2015, 5) {
		if IsNice(line) {
			count++
		}
	}

	fmt.Println(count)
}

func IsNice(s string) bool {
	return ContainsThreeVowels(s) && ContainsDoubleLetter(s) && ContainsNoBadStrings(s)
}

func ContainsThreeVowels(s string) bool {
	var count int
	for _, c := range s {
		if c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u' {
			count++
		}
	}

	return count >= 3
}

func ContainsDoubleLetter(s string) bool {
	for i := 0; i < len(s)-1; i++ {
		if s[i] == s[i+1] {
			return true
		}
	}

	return false
}

func ContainsNoBadStrings(s string) bool {
	banned := []string{"ab", "cd", "pq", "xy"}
	for _, ban := range banned {
		if strings.Contains(s, ban) {
			return false
		}
	}

	return true
}
