package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/aoc"
	"strings"
)

func main() {
	fish := InputToFish()

	for day := 1; day <= 80; day++ {
		N := len(fish)
		for i := 0; i < N; i++ {
			if fish[i] == 0 {
				fish[i] = 6
				fish = append(fish, 8)
				continue
			}

			fish[i]--
		}
	}

	fmt.Println(len(fish))
}

func InputToFish() []int {
	line := aoc.InputToString(2021, 6)

	var fs []int
	for _, s := range strings.Split(strings.TrimSpace(line), ",") {
		fs = append(fs, aoc.ParseInt(s))
	}
	return fs
}
