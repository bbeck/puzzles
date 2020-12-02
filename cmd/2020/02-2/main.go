package main

import (
	"fmt"
	"log"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	var count int
	for _, password := range InputToPasswords(2020, 2) {
		if IsValid(password) {
			count += 1
		}
	}

	fmt.Println(count)
}

func IsValid(p Password) bool {
	min := p.policy.min <= len(p.password) && p.password[p.policy.min-1] == p.policy.letter
	max := p.policy.max <= len(p.password) && p.password[p.policy.max-1] == p.policy.letter
	return min != max
}

type Password struct {
	policy   Policy
	password string
}

type Policy struct {
	min, max int
	letter   byte
}

func InputToPasswords(year, day int) []Password {
	var passwords []Password
	for _, line := range aoc.InputToLines(year, day) {
		var min, max int
		var letter byte
		var password string
		if _, err := fmt.Sscanf(line, "%d-%d %c: %s", &min, &max, &letter, &password); err != nil {
			log.Fatalf("unable to parse line: %s, %v", line, err)
		}

		passwords = append(passwords, Password{
			policy: Policy{
				min:    min,
				max:    max,
				letter: letter,
			},
			password: password,
		})
	}

	return passwords
}
