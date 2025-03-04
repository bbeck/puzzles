package main

import (
	"encoding/hex"
	"fmt"
	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	lengths := append([]byte(in.Line()), []byte{17, 31, 73, 47, 23}...)

	var buffer []byte
	for i := 0; i <= 255; i++ {
		buffer = append(buffer, byte(i))
	}

	var current, skip byte
	for round := 0; round < 64; round++ {
		for _, length := range lengths {
			Reverse(buffer, current, length)
			current += length + skip
			skip++
		}
	}

	fmt.Println(Hash(buffer))
}

func Reverse(buffer []byte, current, length byte) {
	for i := byte(0); i < length/2; i++ {
		buffer[current+i], buffer[current+length-i-1] = buffer[current+length-i-1], buffer[current+i]
	}
}

func Hash(buffer []byte) string {
	hash := make([]byte, len(buffer)/16)
	for chunk := 0; chunk < len(buffer)/16; chunk++ {
		for i := 0; i < 16; i++ {
			hash[chunk] ^= buffer[16*chunk+i]
		}
	}

	return hex.EncodeToString(hash)
}
