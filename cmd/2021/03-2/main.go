package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	ns := aoc.InputToLines(2021, 3)
	L := len(ns[0])

	oxygen := append([]string{}, ns...)
	for bit := 0; len(oxygen) > 1 && bit < L; bit++ {
		zeroes, ones := counts(oxygen, bit)

		var keep uint8
		if zeroes > ones {
			keep = '0'
		} else {
			keep = '1'
		}

		oxygen = filter(oxygen, func(n string) bool { return n[bit] != keep })
	}

	co2 := append([]string{}, ns...)
	for bit := 0; len(co2) > 1 && bit < L; bit++ {
		zeroes, ones := counts(co2, bit)

		var keep uint8
		if zeroes > ones {
			keep = '1'
		} else {
			keep = '0'
		}

		co2 = filter(co2, func(n string) bool { return n[bit] != keep })
	}

	fmt.Println(aoc.ParseIntWithBase(oxygen[0], 2) * aoc.ParseIntWithBase(co2[0], 2))
}

func counts(ns []string, bit int) (int, int) {
	var zeroes, ones int
	for _, n := range ns {
		if n[bit] == '0' {
			zeroes++
		} else {
			ones++
		}
	}
	return zeroes, ones
}

func filter(ns []string, fn func(string) bool) []string {
	var r []string
	for _, n := range ns {
		if !fn(n) {
			r = append(r, n)
		}
	}
	return r
}
