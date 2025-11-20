package main

import (
	"fmt"

	. "github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	var ring = InputToRing()

	var current int
	var offset = 202520252025
	for offset >= ring[current].Len() {
		offset -= ring[current].Len()
		current = (current + 1) % len(ring)
	}

	fmt.Println(ring[current].Value(offset))
}

type Range struct {
	Start, End int
}

func (r Range) Len() int {
	return Abs(r.End-r.Start) + 1
}

func (r Range) Value(offset int) int {
	if r.Start < r.End {
		return r.Start + offset
	} else {
		return r.Start - offset
	}
}

func InputToRing() []Range {
	next := func(forward bool) Range {
		var r Range
		if forward {
			in.Scanf("%d-%d", &r.Start, &r.End)
		} else {
			in.Scanf("%d-%d", &r.End, &r.Start)
		}
		return r
	}

	var front, back []Range
	for in.HasNext() {
		if in.HasNext() {
			front = append(front, next(true))
		}
		if in.HasNext() {
			back = append(back, next(false))
		}
	}

	return append([]Range{{1, 1}}, append(front, Reversed(back)...)...)
}
