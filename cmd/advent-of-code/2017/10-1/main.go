package main

import (
	"fmt"
	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	lengths := in.Ints()

	var buffer []byte
	for i := 0; i <= 255; i++ {
		buffer = append(buffer, byte(i))
	}

	var current, skip byte
	for _, length := range lengths {
		Reverse(buffer, current, byte(length))
		current += byte(length) + skip
		skip++
	}
	fmt.Println(int(buffer[0]) * int(buffer[1]))
}

func Reverse(buffer []byte, current, length byte) {
	for i := byte(0); i < length/2; i++ {
		buffer[current+i], buffer[current+length-i-1] = buffer[current+length-i-1], buffer[current+i]
	}
}
