package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/aoc"
	"strings"
)

func main() {
	var count int
	for _, p := range InputToPasswords() {
		if IsValid(p) {
			count++
		}
	}
	fmt.Println(count)
}

func IsValid(p Password) bool {
	n := strings.Count(p.Value, p.C)
	return p.Min <= n && n <= p.Max
}

type Password struct {
	Min, Max int
	C        string
	Value    string
}

func InputToPasswords() []Password {
	return aoc.InputLinesTo(2020, 2, func(line string) (Password, error) {
		line = strings.ReplaceAll(line, ":", "")

		var password Password
		_, err := fmt.Sscanf(line, "%d-%d %s %s", &password.Min, &password.Max, &password.C, &password.Value)
		return password, err
	})
}
