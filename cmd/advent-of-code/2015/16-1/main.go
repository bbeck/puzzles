package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/bbeck/advent-of-code/puz"
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
	Id     string
	Fields map[string]int
}

func InputToAunts() []Aunt {
	return puz.InputLinesTo(func(line string) Aunt {
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
