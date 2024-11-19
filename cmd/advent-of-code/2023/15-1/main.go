package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/lib"
	"strings"
)

func main() {
	line := lib.InputToString()

	var sum int
	for _, field := range strings.Split(line, ",") {
		sum += Hash(field)
	}
	fmt.Println(sum)
}

func Hash(s string) int {
	var hash int32
	for _, c := range s {
		hash = 17 * (hash + c) % 256
	}
	return int(hash)
}
