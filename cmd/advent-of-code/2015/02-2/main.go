package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/puz"
	"log"
)

type Box struct {
	l, w, h int
}

func (b Box) Volume() int {
	return b.l * b.w * b.h
}
func (b Box) RibbonLength() int {
	dimensions := []int{b.l, b.w, b.h}
	return 2*(puz.Sum(dimensions...)-puz.Max(dimensions...)) + b.Volume()
}

func main() {
	length := 0
	for _, box := range InputToBoxes() {
		length += box.RibbonLength()
	}
	fmt.Println(length)
}

func InputToBoxes() []Box {
	parser := func(line string) Box {
		var box Box
		if _, err := fmt.Sscanf(line, "%dx%dx%d", &box.l, &box.w, &box.h); err != nil {
			log.Fatalf("unable to parse line: %v", err)
		}
		return box
	}
	return puz.InputLinesTo(parser)
}
