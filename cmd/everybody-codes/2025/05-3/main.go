package main

import (
	"fmt"
	"sort"

	. "github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	fbs := InputToBones()

	sort.Slice(fbs, func(i, j int) bool {
		fbi, fbj := fbs[i], fbs[j]

		if fbi.Quality != fbj.Quality {
			return fbi.Quality > fbj.Quality
		}
		for n := range Min(len(fbi.LevelValues), len(fbj.LevelValues)) {
			if fbi.LevelValues[n] != fbj.LevelValues[n] {
				return fbi.LevelValues[n] > fbj.LevelValues[n]
			}
		}
		return fbi.ID > fbj.ID
	})

	var checksum int
	for i, fb := range fbs {
		checksum += (i + 1) * fb.ID
	}
	fmt.Println(checksum)
}

const (
	LEFT = iota
	CENTER
	RIGHT
)

type Fishbone struct {
	ID          int
	LevelValues []int
	Quality     int
}

func InputToBones() []Fishbone {
	return in.LinesToS(func(in in.Scanner[Fishbone]) Fishbone {
		id := in.Int()

		var levels [][3]int

	outer:
		for in.HasNext() {
			n := in.Int()

			for current := 0; current < len(levels); current++ {
				if n < levels[current][CENTER] && levels[current][LEFT] == 0 {
					levels[current][LEFT] = n
					continue outer
				}
				if n > levels[current][CENTER] && levels[current][RIGHT] == 0 {
					levels[current][RIGHT] = n
					continue outer
				}
			}

			levels = append(levels, [3]int{0, n, 0})
		}

		// Convert levels to level values.
		var values []int
		for _, level := range levels {
			var digits []int
			for _, value := range level {
				if value != 0 {
					digits = append(digits, Digits(value)...)
				}
			}
			values = append(values, JoinDigits(digits))
		}

		// Calculate quality.
		var digits []int
		for _, level := range levels {
			digits = append(digits, Digits(level[CENTER])...)
		}
		quality := JoinDigits(digits)

		return Fishbone{ID: id, LevelValues: values, Quality: quality}
	})
}
