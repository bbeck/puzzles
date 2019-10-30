package main

import (
	"fmt"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	desired := aoc.InputToInt(2015, 20)

	maxHome := desired
	homes := make(map[int]int)
	for elf := 1; elf < maxHome; elf++ {
		for count := 0; count < 50; count++ {
			home := elf * (count + 1)
			if home > maxHome {
				break
			}

			homes[home] += elf * 11
			if homes[home] >= desired {
				maxHome = home
			}
		}
	}

	for home := 1; ; home++ {
		if homes[home] >= desired {
			fmt.Printf("home: %d\n", home)
			break
		}
	}
}
