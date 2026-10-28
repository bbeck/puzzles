package main

import (
	"fmt"
	"slices"

	. "github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	var sum int
	for _, orderings := range InputToPartialOrderings() {
		sum += Sort(orderings)
	}
	fmt.Println(sum)
}

func Sort(less Set[Less]) int {
	var seen Set[int]
	for ordering := range less {
		seen.Add(ordering.A, ordering.B)
	}

	var ids = seen.Entries()
	slices.SortFunc(ids, func(i, j int) int {
		if less.Contains(Less{i, j}) {
			return 1 // Hardest to softest
		}
		return -1
	})

	return JoinDigits(ids)
}

type Less struct {
	A, B int
}

func InputToPartialOrderings() []Set[Less] {
	return in.LinesToS(func(in in.Scanner[Set[Less]]) Set[Less] {
		var orderings Set[Less]
		for _, field := range in.FieldsS() {
			a, op, b := field.Int(), field.Byte(), field.Int()
			if op == '>' {
				a, b = b, a
			}
			orderings.Add(Less{a, b})
		}
		return orderings
	})
}
