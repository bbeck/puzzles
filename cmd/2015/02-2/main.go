package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/aoc"
)

type Box struct {
	l, w, h int
}

func (b Box) Volume() int {
	return b.l * b.w * b.h
}
func (b Box) RibbonLength() int {
	dimensions := []int{b.l, b.w, b.h}
	return 2*(aoc.Sum(dimensions...)-aoc.Max(dimensions...)) + b.Volume()
}

func main() {
	length := 0
	for _, box := range InputToBoxes() {
		length += box.RibbonLength()
	}
	fmt.Println(length)
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
