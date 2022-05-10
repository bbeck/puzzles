package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/aoc"
)

type Box struct {
	l, w, h int
}

func (b Box) Area() int {
	var sides = []int{
		b.l * b.w,
		b.w * b.h,
		b.h * b.l,
	}

	return 2*aoc.Sum(sides...) + aoc.Min(sides...)
}

func main() {
	area := 0
	for _, box := range InputToBoxes() {
		area += box.Area()
	}
	fmt.Println(area)
}

func InputToBoxes() []Box {
	parser := func(line string) (Box, error) {
		var box Box
		if _, err := fmt.Sscanf(line, "%dx%dx%d", &box.l, &box.w, &box.h); err != nil {
			return box, err
		}
		return box, nil
	}
	return aoc.InputLinesTo(2015, 2, parser)
}
