package main

import (
	"fmt"

	"github.com/bbeck/puzzles/lib/in"
)

var Target = map[string]func(int) bool{
	"children":    func(n int) bool { return n == 3 },
	"cats":        func(n int) bool { return n > 7 },
	"samoyeds":    func(n int) bool { return n == 2 },
	"pomeranians": func(n int) bool { return n < 3 },
	"akitas":      func(n int) bool { return n == 0 },
	"vizslas":     func(n int) bool { return n == 0 },
	"goldfish":    func(n int) bool { return n < 5 },
	"trees":       func(n int) bool { return n > 3 },
	"cars":        func(n int) bool { return n == 2 },
	"perfumes":    func(n int) bool { return n == 1 },
}

func main() {
	for _, aunt := range InputToAunts() {
		if Matches(aunt, Target) {
			fmt.Println(aunt.Id)
		}
	}
}

func Matches(aunt Aunt, target map[string]func(int) bool) bool {
	for field, actual := range aunt.Fields {
		if !target[field](actual) {
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
	return in.LinesToS(func(in in.Scanner[Aunt]) Aunt {
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
