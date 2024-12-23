package main

import (
	"fmt"
	"github.com/bbeck/puzzles/lib"
)

func main() {
	var sum int
	for _, n := range lib.InputToInts() {
		for range 2000 {
			n = Next(n)
		}
		sum += n
	}
	fmt.Println(sum)
}

func Next(secret int) int {
	secret = (secret ^ (secret * 64)) % 16777216
	secret = (secret ^ (secret / 32)) % 16777216
	secret = (secret ^ (secret * 2048)) % 16777216
	return secret
}
