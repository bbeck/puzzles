package main

import (
	"fmt"
	"strings"

	. "github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	var sum int
	for _, scale := range InputToScales() {
		if scale.G > scale.R && scale.G > scale.B {
			sum += scale.ID
		}
	}
	fmt.Println(sum)
}

type Scale struct {
	ID      int
	R, G, B int
}

func InputToScales() []Scale {
	var replacer = strings.NewReplacer(
		"r", "0", "R", "1",
		"g", "0", "G", "1",
		"b", "0", "B", "1",
	)

	in.Remove(":")
	return in.LinesToS(func(in in.Scanner[Scale]) Scale {
		return Scale{
			ID: in.Int(),
			R:  ParseIntWithBase(replacer.Replace(in.String()), 2),
			G:  ParseIntWithBase(replacer.Replace(in.String()), 2),
			B:  ParseIntWithBase(replacer.Replace(in.String()), 2),
		}
	})
}
