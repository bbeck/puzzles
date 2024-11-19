package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/bbeck/puzzles/lib"
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
	Id     string
	Fields map[string]int
}

func InputToAunts() []Aunt {
	return lib.InputLinesTo(func(line string) Aunt {
		line = strings.ReplaceAll(line, ":", "")
		line = strings.ReplaceAll(line, ",", "")

		var id, compound1, compound2, compound3 string
		var value1, value2, value3 int
		_, err := fmt.Sscanf(
			line,
			"Sue %s %s %d %s %d %s %d",
			&id,
			&compound1, &value1,
			&compound2, &value2,
			&compound3, &value3,
		)
		if err != nil {
			log.Fatalf("unable to parse line: %v", err)
		}

		return Aunt{
			Id: id,
			Fields: map[string]int{
				compound1: value1,
				compound2: value2,
				compound3: value3,
			},
		}
	})
}
