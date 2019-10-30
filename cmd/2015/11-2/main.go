package main

import (
	"fmt"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	password := aoc.InputToString(2015, 11)

	for i := 0; i < 2; i++ {
		for {
			password = NextPassword(password)
			if IsValid(password) {
				break
			}
		}
	}

	fmt.Printf("new password: %s\n", password)
}

func NextPassword(password string) string {
	bs := []byte(password)
	for i := len(bs) - 1; i >= 0; i-- {
		bs[i]++
		if bs[i] <= 'z' {
			break
		}

		bs[i] = 'a'
	}

	return string(bs)
}

func IsValid(password string) bool {
	return HasStraight(password) && HasValidLetters(password) && HasPairs(password)
}

func HasStraight(s string) bool {
	for i := 0; i < len(s)-3; i++ {
		if s[i+1] == s[i]+1 && s[i+2] == s[i+1]+1 {
			return true
		}
	}
	return false
}

func HasValidLetters(s string) bool {
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case 'i':
			return false
		case 'o':
			return false
		case 'l':
			return false
		}
	}

	return true
}

func HasPairs(s string) bool {
	var seen byte
	for i := 0; i < len(s)-1; i++ {
		if s[i] == s[i+1] {
			if seen != 0 && s[i] != seen {
				return true
			}

			seen = s[i]
			i++
		}
	}

	return false
}
