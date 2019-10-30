package main

import (
	"fmt"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	var count int
	for _, line := range aoc.InputToLines(2015, 5) {
		if IsNice(line) {
			count++
		}
	}

	fmt.Printf("count: %d\n", count)
}

func IsNice(s string) bool {
	return ContainsThreeVowels(s) && ContainsDoubleLetter(s) && !ContainsBadString(s)
}

func ContainsThreeVowels(s string) bool {
	var count int
	for _, c := range s {
		if c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u' {
			count++
		}
	}

	return count > 2
}

func ContainsDoubleLetter(s string) bool {
	for i := 0; i < len(s)-1; i++ {
		if s[i] == s[i+1] {
			return true
		}
	}

	return false
}

func ContainsBadString(s string) bool {
	for i := 0; i < len(s)-1; i++ {
		switch {
		case s[i] == 'a' && s[i+1] == 'b':
			return true
		case s[i] == 'c' && s[i+1] == 'd':
			return true
		case s[i] == 'p' && s[i+1] == 'q':
			return true
		case s[i] == 'x' && s[i+1] == 'y':
			return true
		}
	}

	return false
}
