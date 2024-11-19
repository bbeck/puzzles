package main

import (
	"fmt"
	"github.com/bbeck/puzzles/lib"
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
	a := p.Value[p.Min-1] == p.C
	b := p.Value[p.Max-1] == p.C
	return a != b
}

type Password struct {
	Min, Max int
	C        byte
	Value    string
}

func InputToPasswords() []Password {
	return lib.InputLinesTo(func(line string) Password {
		line = strings.ReplaceAll(line, ":", "")

		var password Password
		fmt.Sscanf(line, "%d-%d %c %s", &password.Min, &password.Max, &password.C, &password.Value)
		return password
	})
}
