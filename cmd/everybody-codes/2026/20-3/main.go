package main

import (
	"fmt"
	"strconv"
	"strings"

	. "github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	var count int
	for in.HasNext() {
		var line = in.Line()
		if !IsCorrect(line) {
			var ok bool
			if line, ok = Correct(line); !ok {
				continue
			}
		}

		for idx := range line {
			count += Segments[line[idx]]
		}
	}
	fmt.Println(strconv.FormatInt(int64(count), 11))
}

func IsCorrect(line string) bool {
	eval := func(a, op, b string) int {
		if op == "-" {
			return ParseIntWithBase(a, 11) - ParseIntWithBase(b, 11)
		}
		return ParseIntWithBase(a, 11) + ParseIntWithBase(b, 11)
	}

	var fs = strings.Fields(line)
	return eval(fs[0], fs[1], fs[2]) == eval(fs[4], fs[5], fs[6])
}

func Correct(line string) (string, bool) {
	// Start by considering if we can move a segment within a digit.
	for idx := range line {
		for _, changed := range Move[line[idx]] {
			s := line[:idx] + changed + line[idx+1:]
			if IsCorrect(s) {
				return s, true
			}
		}
	}

	// Next see about removing a segment from one digit and adding it to another.
	for idxRemove := range line {
		for idxAdd := range line {
			if idxRemove == idxAdd {
				continue
			}

			for _, ch1 := range Remove[line[idxRemove]] {
				for _, ch2 := range Add[line[idxAdd]] {
					var s = line[:idxRemove] + ch1 + line[idxRemove+1:]
					s = s[:idxAdd] + ch2 + s[idxAdd+1:]
					if IsCorrect(s) {
						return s, true
					}
				}
			}
		}
	}

	return "", false
}

var Segments = map[byte]int{
	'0': 6, '1': 2, '2': 5, '3': 5, '4': 4, '5': 5, '6': 6,
	'7': 3, '8': 7, '9': 6, 'A': 6, '+': 2, '-': 1, '=': 2,
}

var Remove = map[byte][]string{
	'6': {"5"},
	'7': {"1"},
	'8': {"0", "6", "9", "A"},
	'9': {"3", "5"},
	'+': {"-"},
}

var Add = map[byte][]string{
	'0': {"8"},
	'1': {"7"},
	'3': {"9"},
	'5': {"6", "9"},
	'6': {"8"},
	'9': {"8"},
	'A': {"8"},
	'-': {"+"},
}

var Move = map[byte][]string{
	'0': {"6", "9", "A"},
	'2': {"3"},
	'3': {"2", "5"},
	'5': {"3"},
	'6': {"0", "9", "A"},
	'9': {"0", "6", "A"},
	'A': {"0", "6", "9"},
}
