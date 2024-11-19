package main

import (
	"fmt"
	"github.com/bbeck/puzzles/lib"
)

const N = 4

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
