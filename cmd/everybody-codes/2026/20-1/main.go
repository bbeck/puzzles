package main

import (
	"fmt"
	"strings"

	. "github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	var count int
	for in.HasNext() {
		var line = in.Line()
		if !IsCorrect(line) {
			continue
		}

		for idx := range line {
			count += Segments[line[idx]]
		}
	}
	fmt.Println(count)
}

func IsCorrect(line string) bool {
	eval := func(a, op, b string) int {
		if op == "-" {
			return ParseInt(a) - ParseInt(b)
		}
		return ParseInt(a) + ParseInt(b)
	}

	var fs = strings.Fields(line)
	return eval(fs[0], fs[1], fs[2]) == eval(fs[4], fs[5], fs[6])
}

var Segments = map[byte]int{
	'0': 6, '1': 2, '2': 5, '3': 5, '4': 4, '5': 5, '6': 6,
	'7': 3, '8': 7, '9': 6, '+': 2, '-': 1, '=': 2,
}
