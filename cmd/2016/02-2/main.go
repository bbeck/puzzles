package main

import (
	"fmt"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	keypad := map[aoc.Point2D]int{
		aoc.Point2D{}.Up().Up(): 0x1,

		aoc.Point2D{}.Up().Left():  0x2,
		aoc.Point2D{}.Up():         0x3,
		aoc.Point2D{}.Up().Right(): 0x4,

		aoc.Point2D{}.Left().Left():   0x5,
		aoc.Point2D{}.Left():          0x6,
		aoc.Point2D{}:                 0x7,
		aoc.Point2D{}.Right():         0x8,
		aoc.Point2D{}.Right().Right(): 0x9,

		aoc.Point2D{}.Down().Left():  0xA,
		aoc.Point2D{}.Down():         0xB,
		aoc.Point2D{}.Down().Right(): 0xC,

		aoc.Point2D{}.Down().Down(): 0xD,
	}

	current := aoc.Point2D{}.West().West()
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
		fmt.Printf("%X", num)
	}
	fmt.Println()
}
