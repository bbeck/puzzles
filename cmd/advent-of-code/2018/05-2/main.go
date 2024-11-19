package main

import (
	"bytes"
	"fmt"
	"github.com/bbeck/puzzles/lib"
	"math"
)

func main() {
	matches := make(map[byte]byte)
	for i := byte(0); i < 26; i++ {
		matches['a'+i] = 'A' + i
		matches['A'+i] = 'a' + i
	}

	input := lib.InputToBytes()

	var best = math.MaxInt
	for i := byte(0); i < 26; i++ {
		bs := bytes.ReplaceAll(input, []byte{'a' + i}, nil)
		bs = bytes.ReplaceAll(bs, []byte{'A' + i}, nil)
		best = lib.Min(best, Collapse(bs, matches))
	}

	fmt.Println(best)
}

func Collapse(s []byte, matches map[byte]byte) int {
	var stack lib.Stack[byte]
	for _, c := range s {
		if stack.Peek() == matches[c] {
			stack.Pop()
		} else {
			stack.Push(c)
		}
	}

	return stack.Len()
}
