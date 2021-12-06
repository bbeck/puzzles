package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/aoc"
	"strings"
)

func main() {
	// mapping of fish timer to the count of fish with that timer
	fish := make(map[int]int)
	for _, f := range InputToFish() {
		fish[f]++
	}

	for day := 1; day <= 256; day++ {
		next := make(map[int]int)
		for f, count := range fish {
			if f == 0 {
				next[6] += count
				next[8] += count
				continue
			}
			next[f-1] += count
		}

		fish = next
	}

	var count int
	for _, c := range fish {
		count += c
	}

	fmt.Println(count)
}

func InputToFish() []int {
	line := aoc.InputToString(2021, 6)

	var fs []int
	for _, s := range strings.Split(strings.TrimSpace(line), ",") {
		fs = append(fs, aoc.ParseInt(s))
	}
	return fs
}
