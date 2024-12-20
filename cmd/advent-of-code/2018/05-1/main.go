package main

import (
	"fmt"
	"github.com/bbeck/puzzles/lib"
)

func main() {
	matches := make(map[rune]rune)
	for i := 0; i < 26; i++ {
		matches[rune('a'+i)] = rune('A' + i)
		matches[rune('A'+i)] = rune('a' + i)
	}

	var stack lib.Stack[rune]
	for _, c := range lib.InputToString() {
		if stack.Peek() == matches[c] {
			stack.Pop()
		} else {
			stack.Push(c)
		}
	}

	fmt.Println(stack.Len())
}
