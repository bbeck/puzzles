package main

import (
	"fmt"
	"strings"

	. "github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	scales := InputToScales()

	group := func(s Scale) *string {
		var color string
		switch {
		case s.R > s.G && s.R > s.B:
			color = "red"
		case s.G > s.R && s.G > s.B:
			color = "green"
		case s.B > s.R && s.B > s.G:
			color = "blue"
		}

		var brightness string
		switch {
		case 0 <= s.S && s.S <= 30:
			brightness = "matte"
		case 33 <= s.S:
			brightness = "shiny"
		}

		var group *string
		if color != "" && brightness != "" {
			key := color + "-" + brightness
			group = &key
		}
		return group
	}

	var fc FrequencyCounter[string]
	var sums = make(map[string]int)
	for _, scale := range scales {
		if g := group(scale); g != nil {
			fc.Add(*g)
			sums[*g] += scale.ID
		}
	}

	var largest = fc.Entries()[0].Value
	fmt.Println(sums[largest])
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
