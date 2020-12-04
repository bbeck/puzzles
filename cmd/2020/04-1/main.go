package main

import (
	"fmt"
	"strings"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	var count int
	for _, passport := range InputToPassports(2020, 04) {
		if IsValid(passport) {
			count++
		}
	}

	fmt.Println(count)
}

func IsValid(passport Passport) bool {
	if len(passport) == 8 {
		return true
	}

	if len(passport) == 7 {
		_, found := passport["cid"]
		return !found
	}

	return false
}

type Passport map[string]string

func InputToPassports(year, day int) []Passport {
	var passports []Passport

	passport := make(Passport)
	for _, line := range aoc.InputToLines(year, day) {
		if len(line) == 0 {
			passports = append(passports, passport)
			passport = make(Passport)
			continue
		}

		for _, field := range strings.Split(line, " ") {
			parts := strings.Split(field, ":")
			key := parts[0]
			value := parts[1]

			passport[key] = value
		}
	}
	passports = append(passports, passport)

	return passports
}
