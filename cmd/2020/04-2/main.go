package main

import (
	"fmt"
	"regexp"
	"strconv"
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
	isInRange := func(name string, min, max int) bool {
		value, ok := passport[name]
		if !ok {
			return false
		}

		n, err := strconv.Atoi(value)
		if err != nil {
			return false
		}

		return min <= n && n <= max
	}

	isHeight := func(name string) bool {
		value, ok := passport[name]
		if !ok {
			return false
		}

		L := len(value)

		switch {
		case strings.HasSuffix(value, "cm"):
			n := aoc.ParseInt(value[0 : L-2])
			return 150 <= n && n <= 193
		case strings.HasSuffix(value, "in"):
			n := aoc.ParseInt(value[0 : L-2])
			return 59 <= n && n <= 76
		default:
			return false
		}
	}

	isColor := func(name string) bool {
		value, ok := passport[name]
		if !ok {
			return false
		}

		match, _ := regexp.MatchString("^#[0-9a-f]{6}$", value)
		return match
	}

	isEyeColor := func(name string) bool {
		value, ok := passport[name]
		if !ok {
			return false
		}

		colors := map[string]bool{
			"amb": true, "blu": true, "brn": true, "gry": true,
			"grn": true, "hzl": true, "oth": true,
		}
		return colors[value]
	}

	isPassportId := func(name string) bool {
		value, ok := passport[name]
		if !ok {
			return false
		}

		match, _ := regexp.MatchString("^[0-9]{9}$", value)
		return match
	}

	return isInRange("byr", 1920, 2002) &&
		isInRange("iyr", 2010, 2020) &&
		isInRange("eyr", 2020, 2030) &&
		isHeight("hgt") &&
		isColor("hcl") &&
		isEyeColor("ecl") &&
		isPassportId("pid")
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
