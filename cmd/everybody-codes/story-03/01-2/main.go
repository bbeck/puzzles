package main

import (
	"cmp"
	"fmt"
	"slices"
	"strings"

	. "github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	scales := InputToScales()
	slices.SortFunc(scales, func(a, b Scale) int {
		if c := cmp.Compare(a.S, b.S); c != 0 {
			return -c
		}
		return cmp.Compare(a.R+a.G+a.B, b.R+b.G+b.B)
	})
	fmt.Println(scales[0].ID)
}

type Scale struct {
	ID         int
	R, G, B, S int
}

func InputToScales() []Scale {
	var replacer = strings.NewReplacer(
		"r", "0", "R", "1",
		"g", "0", "G", "1",
		"b", "0", "B", "1",
		"s", "0", "S", "1",
	)

	in.Remove(":")
	return in.LinesToS(func(in in.Scanner[Scale]) Scale {
		return Scale{
			ID: in.Int(),
			R:  ParseIntWithBase(replacer.Replace(in.String()), 2),
			G:  ParseIntWithBase(replacer.Replace(in.String()), 2),
			B:  ParseIntWithBase(replacer.Replace(in.String()), 2),
			S:  ParseIntWithBase(replacer.Replace(in.String()), 2),
		}
	})
}
