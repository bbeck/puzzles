package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	ns := aoc.SetFrom(aoc.InputToLines(2021, 3)...)

	o2 := aoc.SetFrom(ns.Entries()...)
	for pos := 0; len(o2) > 1; pos++ {
		zs, os := Partition(o2, pos)
		if len(os) >= len(zs) {
			o2 = os
		} else {
			o2 = zs
		}
	}

	co2 := aoc.SetFrom(ns.Entries()...)
	for pos := 0; len(co2) > 1; pos++ {
		zs, os := Partition(co2, pos)
		if len(os) < len(zs) {
			co2 = os
		} else {
			co2 = zs
		}
	}

	a := aoc.ParseIntWithBase(o2.Entries()[0], 2)
	b := aoc.ParseIntWithBase(co2.Entries()[0], 2)
	fmt.Println(a * b)
}

func Partition(ns aoc.Set[string], position int) (aoc.Set[string], aoc.Set[string]) {
	var zs, os aoc.Set[string]
	for n := range ns {
		if n[position] == '0' {
			zs.Add(n)
		} else {
			os.Add(n)
		}
	}

	return zs, os
}
