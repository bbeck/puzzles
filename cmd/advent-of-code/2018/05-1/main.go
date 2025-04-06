package main

import (
	"fmt"
	. "github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	matches := make(map[byte]byte)
	for i := range byte(26) {
		matches['a'+i] = 'A' + i
		matches['A'+i] = 'a' + i
	}

	var stack Stack[byte]
	for in.HasNext() {
		c := in.Byte()
		if stack.Peek() == matches[c] {
			stack.Pop()
		} else {
			stack.Push(c)
		}
	}

	fmt.Println(stack.Len())
}
