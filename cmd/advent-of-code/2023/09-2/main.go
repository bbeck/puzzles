package main

import (
	"fmt"
	"strings"

	"github.com/bbeck/puzzles/lib"
)

func main() {
	var sum int
	for _, line := range InputToIntLines() {
		sum += Extrapolate(lib.Reversed(line))
	}
	fmt.Println(sum)
}

func Extrapolate(ns []int) int {
	if len(ns) == 1 {
		return ns[0]
	}

	var deltas []int
	for i := 1; i < len(ns); i++ {
		deltas = append(deltas, ns[i]-ns[i-1])
	}
	return Extrapolate(deltas) + ns[len(ns)-1]
}

func InputToIntLines() [][]int {
	return lib.InputLinesTo(func(line string) []int {
		var ns []int
		for _, s := range strings.Fields(line) {
			ns = append(ns, lib.ParseInt(s))
		}
		return ns
	})
}
