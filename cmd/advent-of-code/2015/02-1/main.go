package main

import (
	"fmt"
	"github.com/bbeck/puzzles/lib"
	"log"
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

	return 2*lib.Sum(sides...) + lib.Min(sides...)
}

func main() {
	area := 0
	for _, box := range InputToBoxes() {
		area += box.Area()
	}
	fmt.Println(area)
}

func InputToBoxes() []Box {
	parser := func(line string) Box {
		var box Box
		if _, err := fmt.Sscanf(line, "%dx%dx%d", &box.l, &box.w, &box.h); err != nil {
			log.Fatalf("unable to parse line: %v", err)
		}
		return box
	}
	return lib.InputLinesTo(parser)
}
