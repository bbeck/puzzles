package main

import (
	"fmt"
	"strings"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	var count int
	for _, p := range InputToPassports() {
		delete(p, "cid")
		if len(p) == 7 {
			count++
		}
	}
	fmt.Println(count)
}

type Passport map[string]string

func InputToPassports() []Passport {
	var passports []Passport

	current := make(Passport)
	for _, line := range aoc.InputToLines(2020, 4) {
		if len(line) == 0 {
			passports = append(passports, current)
			current = make(Passport)
			continue
		}

		for _, field := range strings.Fields(line) {
			key, value, _ := strings.Cut(field, ":")
			current[key] = value
		}
	}
	passports = append(passports, current)

	return passports
}
