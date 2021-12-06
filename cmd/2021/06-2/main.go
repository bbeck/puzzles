package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/aoc"
	"strings"
)

func main() {
	fish := make(map[Fish]int)
	for _, f := range InputToFish() {
		fish[f]++
	}

	for day := 1; day <= 256; day++ {
		next := make(map[Fish]int)
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

type Fish int

func InputToFish() []Fish {
	line := aoc.InputToString(2021, 6)

	var fs []Fish
	for _, s := range strings.Split(strings.TrimSpace(line), ",") {
		n := aoc.ParseInt(s)
		fs = append(fs, Fish(n))
	}
	return fs
}
