package main

import (
	"fmt"
	"github.com/bbeck/puzzles/lib"
)

func main() {
	ns := lib.SetFrom(lib.InputToLines()...)

	o2 := ns
	for pos := 0; len(o2) > 1; pos++ {
		zs, os := Partition(o2, pos)
		if len(os) >= len(zs) {
			o2 = os
		} else {
			o2 = zs
		}
	}

	co2 := ns
	for pos := 0; len(co2) > 1; pos++ {
		zs, os := Partition(co2, pos)
		if len(os) < len(zs) {
			co2 = os
		} else {
			co2 = zs
		}
	}

	a := lib.ParseIntWithBase(o2.Entries()[0], 2)
	b := lib.ParseIntWithBase(co2.Entries()[0], 2)
	fmt.Println(a * b)
}

func Partition(ns lib.Set[string], position int) (lib.Set[string], lib.Set[string]) {
	var zs, os lib.Set[string]
	for n := range ns {
		if n[position] == '0' {
			zs.Add(n)
		} else {
			os.Add(n)
		}
	}

	return zs, os
}
