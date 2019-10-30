package main

import (
	"fmt"
	"log"

	"github.com/bbeck/advent-of-code/aoc"
)

type Prism struct {
	l, w, h uint
}

func (p Prism) Area() uint {
	s1 := p.l * p.w
	s2 := p.l * p.h
	s3 := p.w * p.h

	var smallest uint
	if s1 <= s2 && s1 <= s3 {
		smallest = s1
	} else if s2 <= s1 && s2 <= s3 {
		smallest = s2
	} else {
		smallest = s3
	}

	return 2*s1 + 2*s2 + 2*s3 + smallest
}

func (p Prism) RibbonLength() uint {
	p1 := 2*p.l + 2*p.w
	p2 := 2*p.l + 2*p.h
	p3 := 2*p.w + 2*p.h

	var smallest uint
	if p1 <= p2 && p1 <= p3 {
		smallest = p1
	} else if p2 <= p1 && p2 <= p3 {
		smallest = p2
	} else {
		smallest = p3
	}

	return smallest + p.l*p.w*p.h
}

func main() {
	total := uint(0)
	for _, prism := range InputToPrisms() {
		total += prism.RibbonLength()
	}

	fmt.Printf("total: %d\n", total)
}

func InputToPrisms() []Prism {
	prisms := make([]Prism, 0)
	for _, line := range aoc.InputToLines(2015, 2) {
		var l, w, h uint
		if _, err := fmt.Sscanf(line, "%dx%dx%d", &l, &w, &h); err != nil {
			log.Fatalf("unable to parse line '%s': %+v", line, err)
		}

		prisms = append(prisms, Prism{l, w, h})
	}

	return prisms
}
