package main

import (
	"fmt"
	"strings"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	var buffer []byte
	for i := 0; i <= 255; i++ {
		buffer = append(buffer, byte(i))
	}

	var current, skip byte
	for _, length := range InputToLengths() {
		Reverse(buffer, current, length)
		current += length + skip
		skip++
	}
	fmt.Println(int(buffer[0]) * int(buffer[1]))
}

func Reverse(buffer []byte, current, length byte) {
	for i := byte(0); i < length/2; i++ {
		buffer[current+i], buffer[current+length-i-1] = buffer[current+length-i-1], buffer[current+i]
	}
}

func InputToLengths() []byte {
	input := aoc.InputToString(2017, 10)

	var lengths []byte
	for _, s := range strings.Split(input, ",") {
		lengths = append(lengths, byte(aoc.ParseInt(s)))
	}

	return lengths
}
