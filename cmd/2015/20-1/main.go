package main

import (
	"fmt"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	desired := aoc.InputToInt(2015, 20)

	homes := make([]int, desired+1)
	for elf := 1; elf < desired; elf++ {
		for home := elf; home < desired; home = home + elf {
			homes[home] += elf * 10
		}
	}

	for home, presents := range homes {
		if presents >= desired {
			fmt.Printf("home: %d\n", home)
			break
		}
	}
}
