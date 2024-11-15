package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/puz"
)

const N = 14

func main() {
	s := puz.InputToString(2022, 6)
	for end := N; end < len(s); end++ {
		bs := []byte(s[end-N : end])
		if len(puz.SetFrom(bs...)) == N {
			fmt.Println(end)
			break
		}
	}
}
