package main

import (
	"fmt"
	. "github.com/bbeck/puzzles/lib"
	"math"
	"strings"
)

func main() {
	stones := InputToStones()

	memo := make(map[Key]int)

	var count int
	for _, stone := range stones {
		count += Count(stone, 75, memo)
	}
	fmt.Println(count)
}

type Key struct {
	Stone int
	Times int
}

func Count(stone, n int, memo map[Key]int) int {
	key := Key{Stone: stone, Times: n}
	if c, ok := memo[key]; ok {
		return c
	}

	if n == 0 {
		return 1
	}

	digits := 1 + uint(math.Floor(math.Log10(float64(stone))))

	var count int
	switch {
	case stone == 0:
		count = Count(1, n-1, memo)

	case digits%2 == 0:
		d := Pow(10, digits/2)
		lhs := stone / d
		rhs := stone % d
		count = Count(lhs, n-1, memo) + Count(rhs, n-1, memo)

	default:
		count = Count(2024*stone, n-1, memo)
	}

	memo[key] = count
	return count
}

func InputToStones() []int {
	var ns []int
	for _, s := range strings.Fields(InputToString()) {
		ns = append(ns, ParseInt(s))
	}

	return ns
}
