package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	key := aoc.InputToString(2017, 14)

	grid := aoc.NewGrid2D[bool](128, 128)
	for row := 0; row < grid.Height; row++ {
		hash := KnotHash(fmt.Sprintf("%s-%d", key, row))

		// Convert the slice of bytes into binary, keeping in mind that we're
		// reading the hash from MSB to LSB.
		for i, c := range hash {
			for bit := 0; bit < 8; bit++ {
				col := 8*i + (8 - bit - 1)
				grid.AddXY(col, row, c&(1<<bit) > 0)
			}
		}
	}

	var count int
	for row := 0; row < grid.Height; row++ {
		for col := 0; col < grid.Width; col++ {
			if grid.GetXY(col, row) {
				count++
			}
		}
	}
	fmt.Println(count)
}

func KnotHash(s string) []byte {
	bs := []byte(s)
	bs = append(bs, []byte{17, 31, 73, 47, 23}...)

	var buffer []byte
	for i := 0; i <= 255; i++ {
		buffer = append(buffer, byte(i))
	}

	var current, skip byte
	for round := 0; round < 64; round++ {
		for _, length := range bs {
			Reverse(buffer, current, length)
			current += length + skip
			skip++
		}
	}

	hash := make([]byte, len(buffer)/16)
	for chunk := 0; chunk < len(buffer)/16; chunk++ {
		for i := 0; i < 16; i++ {
			hash[chunk] ^= buffer[16*chunk+i]
		}
	}

	return hash
}

func Reverse[T any](buffer []T, current, length byte) {
	for i := byte(0); i < length/2; i++ {
		buffer[current+i], buffer[current+length-i-1] = buffer[current+length-i-1], buffer[current+i]
	}
}
