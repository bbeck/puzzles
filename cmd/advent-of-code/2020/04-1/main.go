package main

import (
	"fmt"
	"strings"

	"github.com/bbeck/puzzles/lib"
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
	for _, line := range lib.InputToLines() {
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
