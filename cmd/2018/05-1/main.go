package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/puz"
)

func main() {
	matches := make(map[rune]rune)
	for i := 0; i < 26; i++ {
		matches[rune('a'+i)] = rune('A' + i)
		matches[rune('A'+i)] = rune('a' + i)
	}

	var stack puz.Stack[rune]
	for _, c := range puz.InputToString(2018, 5) {
		if stack.Peek() == matches[c] {
			stack.Pop()
		} else {
			stack.Push(c)
		}
	}

	fmt.Println(stack.Len())
}
