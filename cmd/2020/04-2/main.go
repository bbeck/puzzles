package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	var count int
	for _, p := range InputToPassports() {
		if IsValid(p) {
			count++
		}
	}
	fmt.Println(count)
}

func IsValid(p Passport) bool {
	delete(p, "cid")
	if len(p) != 7 {
		return false
	}

	if byr := aoc.ParseInt(p["byr"]); byr < 1920 || byr > 2002 {
		return false
	}

	if iyr := aoc.ParseInt(p["iyr"]); iyr < 2010 || iyr > 2020 {
		return false
	}

	if eyr := aoc.ParseInt(p["eyr"]); eyr < 2020 || eyr > 2030 {
		return false
	}

	switch hgt := p["hgt"]; {
	case strings.HasSuffix(hgt, "cm"):
		if value := aoc.ParseInt(hgt[:len(hgt)-2]); value < 150 || value > 193 {
			return false
		}

	case strings.HasSuffix(hgt, "in"):
		if value := aoc.ParseInt(hgt[:len(hgt)-2]); value < 59 || value > 76 {
			return false
		}

	default:
		return false
	}

	if matched, err := regexp.MatchString("^#[0-9a-f]{6}$", p["hcl"]); err != nil || !matched {
		return false
	}

	switch ecl := p["ecl"]; ecl {
	case "amb":
	case "blu":
	case "brn":
	case "gry":
	case "grn":
	case "hzl":
	case "oth":
	default:
		return false
	}

	if matched, err := regexp.MatchString("^[0-9]{9}$", p["pid"]); err != nil || !matched {
		return false
	}

	return true
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
