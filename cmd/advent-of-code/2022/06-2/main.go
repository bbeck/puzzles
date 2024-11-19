package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/lib"
)

const N = 14

func main() {
	s := lib.InputToString()
	for end := N; end < len(s); end++ {
		bs := []byte(s[end-N : end])
		if len(lib.SetFrom(bs...)) == N {
			fmt.Println(end)
			break
		}
	}
}
