package main

import (
	"fmt"
	"math"
	"strings"

	. "github.com/bbeck/puzzles/lib"
)

func main() {
	memo := make(map[string]int)

	var sum int
	for _, code := range InputToLines() {
		length := MinLength(code, 2+1, NumericKeypad, memo)
		n := ParseInt(code[:len(code)-1])
		sum += length * n
	}
	fmt.Println(sum)
}

type Keypad map[Point2D]string

var NumericKeypad = ToKeypad(`789456123.0A`)
var ArrowKeypad = ToKeypad(`.^A<v>`)

func ToKeypad(s string) Keypad {
	pts := make(map[string]Point2D)
	for y, line := range Chunks([]byte(s), 3) {
		for x, ch := range line {
			if ch != '.' {
				pts[string(ch)] = Point2D{X: x, Y: y}
			}
		}
	}

	keypad := make(map[Point2D]string)
	for c, p := range pts {
		keypad[Point2D{X: p.X - pts["A"].X, Y: p.Y - pts["A"].Y}] = c
	}
	return keypad
}

func MinLength(code string, n int, keypad Keypad, memo map[string]int) int {
	if n < 0 {
		return 1
	}

	key := fmt.Sprintf("%s|%d", code, n)
	if v, ok := memo[key]; ok {
		return v
	}

	var current Point2D
	var total int
	for _, ch := range code {
		paths, next := FindShortestPaths(keypad, current, string(ch))

		var best = math.MaxInt
		for _, p := range paths {
			best = Min(best, MinLength(p, n-1, ArrowKeypad, memo))
		}

		total += best
		current = next
	}

	memo[key] = total
	return total
}

func FindShortestPaths(keypad Keypad, start Point2D, s string) ([]string, Point2D) {
	allPaths, _ := AllShortestPaths(
		start,
		func(p Point2D) []Point2D {
			var children []Point2D
			for _, child := range p.OrthogonalNeighbors() {
				if _, ok := keypad[child]; ok {
					children = append(children, child)
				}
			}
			return children
		},
		func(p Point2D) bool { return keypad[p] == s },
		func(Point2D, Point2D) int { return 1 },
	)

	var current Point2D
	var paths []string
	for _, path := range allPaths {
		current = start
		var steps []string
		for _, p := range path[1:] {
			switch p {
			case current.Up():
				steps = append(steps, "^")
			case current.Right():
				steps = append(steps, ">")
			case current.Down():
				steps = append(steps, "v")
			case current.Left():
				steps = append(steps, "<")
			}
			current = p
		}
		steps = append(steps, "A")
		paths = append(paths, strings.Join(steps, ""))
	}

	return paths, current
}
