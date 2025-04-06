package main

import (
	"bytes"
	"fmt"
	. "github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/in"
	"math"
)

func main() {
	matches := make(map[byte]byte)
	for i := range byte(26) {
		matches['a'+i] = 'A' + i
		matches['A'+i] = 'a' + i
	}

	var input = in.Bytes()

	var best = math.MaxInt
	for i := range byte(26) {
		bs := bytes.ReplaceAll(input, []byte{'a' + i}, nil)
		bs = bytes.ReplaceAll(bs, []byte{'A' + i}, nil)
		best = Min(best, Collapse(bs, matches))
	}

	fmt.Println(best)
}

func Collapse(s []byte, matches map[byte]byte) int {
	var stack Stack[byte]
	for _, c := range s {
		if stack.Peek() == matches[c] {
			stack.Pop()
		} else {
			stack.Push(c)
		}
	}

	return stack.Len()
}
