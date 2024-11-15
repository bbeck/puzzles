package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/puz"
)

const N = 4

func main() {
	s := puz.InputToString()
	for end := N; end < len(s); end++ {
		bs := []byte(s[end-N : end])
		if len(puz.SetFrom(bs...)) == N {
			fmt.Println(end)
			break
		}
	}
}
