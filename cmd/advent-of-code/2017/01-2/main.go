package main

import (
	"fmt"

	"github.com/bbeck/advent-of-code/lib"
)

func main() {
	s := lib.InputToString()
	N := len(s)

	var sum int
	for i := 0; i < N; i++ {
		j := (i + N/2 + N) % N
		if s[i] == s[j] {
			sum += lib.ParseInt(string(s[i]))
		}
	}
	fmt.Println(sum)
}
