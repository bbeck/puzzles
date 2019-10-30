package main

import (
	"fmt"
	"strings"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	ops := InputToOperations(2016, 21)
	s := "abcdefgh"
	lines := aoc.InputToLines(2016, 21)

	for i, op := range ops {
		s2 := op(s)
		fmt.Printf("%2d (%-40s): %s -> %s\n", i, lines[i], s, s2)
		s = s2
	}
}

type Operation func(s string) string

func InputToOperations(year, day int) []Operation {
	var operations []Operation
	for _, line := range aoc.InputToLines(year, day) {
		tokens := strings.Split(line, " ")
		switch {
		case tokens[0] == "swap" && tokens[1] == "position":
			x := aoc.ParseInt(tokens[2])
			y := aoc.ParseInt(tokens[5])

			operations = append(operations, func(s string) string {
				bs := []byte(s)
				bs[x], bs[y] = bs[y], bs[x]
				return string(bs)
			})

		case tokens[0] == "swap" && tokens[1] == "letter":
			x := tokens[2]
			y := tokens[5]

			operations = append(operations, func(s string) string {
				s = strings.ReplaceAll(s, x, ".")
				s = strings.ReplaceAll(s, y, x)
				s = strings.ReplaceAll(s, ".", y)
				return s
			})

		case tokens[0] == "rotate" && tokens[1] == "left":
			x := aoc.ParseInt(tokens[2])

			operations = append(operations, func(s string) string {
				for i := 0; i < x; i++ {
					s = s[1:] + string(s[0])
				}

				return s
			})

		case tokens[0] == "rotate" && tokens[1] == "right":
			x := aoc.ParseInt(tokens[2])

			operations = append(operations, func(s string) string {
				for i := 0; i < x; i++ {
					s = string(s[len(s)-1]) + s[:len(s)-1]
				}

				return s
			})

		case tokens[0] == "rotate" && tokens[1] == "based":
			x := tokens[6]

			operations = append(operations, func(s string) string {
				index := strings.Index(s, x)
				if index >= 4 {
					index++
				}
				index++

				x := index
				for i := 0; i < x; i++ {
					s = string(s[len(s)-1]) + s[:len(s)-1]
				}

				return s
			})

		case tokens[0] == "reverse" && tokens[1] == "positions":
			x := aoc.ParseInt(tokens[2])
			y := aoc.ParseInt(tokens[4])

			operations = append(operations, func(s string) string {
				bs := []byte(s)
				for start, end := x, y; start < end; start, end = start+1, end-1 {
					bs[start], bs[end] = bs[end], bs[start]
				}

				return string(bs)
			})

		case tokens[0] == "move" && tokens[1] == "position":
			x := aoc.ParseInt(tokens[2])
			y := aoc.ParseInt(tokens[5])

			operations = append(operations, func(s string) string {
				c := s[x]
				s = s[:x] + s[x+1:]
				if y < len(s) {
					s = s[:y] + string(c) + s[y:]
				} else {
					s = s[:y] + string(c)
				}
				return s
			})
		}
	}

	return operations
}
