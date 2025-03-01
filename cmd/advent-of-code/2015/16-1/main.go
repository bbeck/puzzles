package main

import (
	"fmt"

	"github.com/bbeck/puzzles/lib/in"
)

var Target = map[string]int{
	"children":    3,
	"cats":        7,
	"samoyeds":    2,
	"pomeranians": 3,
	"akitas":      0,
	"vizslas":     0,
	"goldfish":    5,
	"trees":       3,
	"cars":        2,
	"perfumes":    1,
}

func main() {
	for _, aunt := range InputToAunts() {
		if Matches(aunt, Target) {
			fmt.Println(aunt.Id)
		}
	}
}

func Matches(aunt Aunt, target map[string]int) bool {
	for field, actual := range aunt.Fields {
		if target[field] != actual {
			return false
		}
	}

	return true
}

type Aunt struct {
	Id     int
	Fields map[string]int
}

func InputToAunts() []Aunt {
	return in.LinesTo(func(in *in.Scanner[Aunt]) Aunt {
		var id, v1, v2, v3 int
		var f1, f2, f3 string
		in.Scanf("Sue %d: %s: %d, %s: %d, %s: %d", &id, &f1, &v1, &f2, &v2, &f3, &v3)

		return Aunt{
			Id: id,
			Fields: map[string]int{
				f1: v1,
				f2: v2,
				f3: v3,
			},
		}
	})
}
