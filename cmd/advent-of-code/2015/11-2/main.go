package main

import (
	"fmt"

	"github.com/bbeck/advent-of-code/puz"
)

func main() {
	password := puz.InputToBytes()

	for i := 0; i < 2; i++ {
		for {
			password = NextPassword(password)
			if IsValid(password) {
				break
			}
		}
	}

	fmt.Println(string(password))
}

func NextPassword(password []byte) []byte {
	for i := len(password) - 1; i >= 0; i-- {
		password[i]++
		if password[i] <= 'z' {
			break
		}

		password[i] = 'a'
	}

	return password
}

func IsValid(password []byte) bool {
	return HasStraight(password) && HasValidLetters(password) && HasPairs(password)
}

func HasStraight(bs []byte) bool {
	for i := 0; i < len(bs)-3; i++ {
		if bs[i+1] == bs[i]+1 && bs[i+2] == bs[i+1]+1 {
			return true
		}
	}
	return false
}

func HasValidLetters(bs []byte) bool {
	for _, b := range bs {
		if b == 'i' || b == 'o' || b == 'l' {
			return false
		}
	}

	return true
}

func HasPairs(bs []byte) bool {
	var seen byte
	for i := 0; i < len(bs)-1; i++ {
		if bs[i] == bs[i+1] {
			if seen != 0 && bs[i] != seen {
				return true
			}

			seen = bs[i]
			i++
		}
	}

	return false
}
