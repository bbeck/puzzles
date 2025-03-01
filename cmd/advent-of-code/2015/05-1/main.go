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
	return ContainsThreeVowels(s) && ContainsDoubleLetter(s) && ContainsNoBadStrings(s)
}

func ContainsThreeVowels(s string) bool {
	var count int
	for _, c := range s {
		if c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u' {
			count++
			if count >= 3 {
				return true
			}
		}
	}

	return false
}

func ContainsDoubleLetter(s string) bool {
	for i := range len(s) - 1 {
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
