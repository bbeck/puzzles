package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	s := aoc.InputToString(2018, 5)
	alphabet := "abcdefghijklmnopqrstuvwxyz"

	shortest := math.MaxInt64
	for _, c := range alphabet {
		s1 := strings.ReplaceAll(s, string(c), "")
		s2 := strings.ReplaceAll(s1, strings.ToUpper(string(c)), "")
		length := len(Collapse(s2))

		if length < shortest {
			shortest = length
		}
	}

	fmt.Printf("shortest: %d\n", shortest)
}

func Collapse(s string) string {
	replacements := []string{
		"aA", "bB", "cC", "dD", "eE", "fF", "gG", "hH", "iI", "jJ", "kK", "lL", "mM",
		"nN", "oO", "pP", "qQ", "rR", "sS", "tT", "uU", "vV", "wW", "xX", "yY", "zZ",
		"Aa", "Bb", "Cc", "Dd", "Ee", "Ff", "Gg", "Hh", "Ii", "Jj", "Kk", "Ll", "Mm",
		"Nn", "Oo", "Pp", "Qq", "Rr", "Ss", "Tt", "Uu", "Vv", "Ww", "Xx", "Yy", "Zz",
	}

	for {
		size := len(s)
		for _, replacement := range replacements {
			s = strings.ReplaceAll(s, replacement, "")
		}

		if len(s) == size {
			break
		}
	}

	return s
}
