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
	count := 0
	for _, c := range p.password {
		if c == p.policy.letter {
			count++
		}
	}

	return p.policy.min <= count && count <= p.policy.max
}

type Password struct {
	policy   Policy
	password string
}

type Policy struct {
	min, max int
	letter   int32
}

func InputToPasswords(year, day int) []Password {
	var passwords []Password
	for _, line := range aoc.InputToLines(year, day) {
		var min, max int
		var letter int32
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
