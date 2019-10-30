package main

import (
	"fmt"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	keypad := map[aoc.Point2D]int{
		aoc.Point2D{}.Up().Left():  1,
		aoc.Point2D{}.Up():         2,
		aoc.Point2D{}.Up().Right(): 3,

		aoc.Point2D{}.Left():  4,
		aoc.Point2D{}:         5,
		aoc.Point2D{}.Right(): 6,

		aoc.Point2D{}.Down().Left():  7,
		aoc.Point2D{}.Down():         8,
		aoc.Point2D{}.Down().Right(): 9,
	}

	current := aoc.Point2D{1, 1}
	for _, line := range aoc.InputToLines(2016, 2) {
		for _, c := range line {
			switch c {
			case 'U':
				next := current.Up()
				if keypad[next] != 0 {
					current = next
				}
			case 'D':
				next := current.Down()
				if keypad[next] != 0 {
					current = next
				}
			case 'L':
				next := current.Left()
				if keypad[next] != 0 {
					current = next
				}
			case 'R':
				next := current.Right()
				if keypad[next] != 0 {
					current = next
				}
			}
		}

		num := keypad[current]
		fmt.Printf("%d", num)
	}
	fmt.Println()
}
