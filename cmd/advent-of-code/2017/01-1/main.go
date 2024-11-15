package main

import (
	"fmt"

	"github.com/bbeck/advent-of-code/puz"
)

func main() {
	s := puz.InputToString()
	N := len(s)

	var sum int
	for i := 0; i < N; i++ {
		j := (i + 1 + N) % N
		if s[i] == s[j] {
			sum += puz.ParseInt(string(s[i]))
		}
	}
	fmt.Println(sum)
}
