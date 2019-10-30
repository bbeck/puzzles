package main

import (
	"fmt"
	"strings"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	s := aoc.InputToString(2018, 5)
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

	fmt.Printf("length: %d\n", len(s))
}
