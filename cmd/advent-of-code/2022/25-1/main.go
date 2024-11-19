package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/lib"
	"strings"
)

func main() {
	var sum int
	for _, line := range lib.InputToLines() {
		sum += Parse(line)
	}

	fmt.Println(Encode(sum))
}

var DigitToValue = map[rune]int{
	'0': 0,
	'1': 1,
	'2': 2,
	'-': -1,
	'=': -2,
}

func Parse(s string) int {
	N := len(s)

	var n int
	for i, c := range s {
		pow := lib.Pow(5, uint(N-i-1))
		n += DigitToValue[c] * pow
	}
	return n
}

var ValueToDigit = map[int]rune{
	-2: '=',
	-1: '-',
	0:  '0',
	1:  '1',
	2:  '2',
}

func Encode(n int) string {
	var s lib.Stack[rune]

	var mod int
	for n > 0 {
		n, mod = (n+2)/5, (n+2)%5
		s.Push(ValueToDigit[mod-2])
	}

	var sb strings.Builder
	for !s.Empty() {
		sb.WriteRune(s.Pop())
	}
	return sb.String()
}
