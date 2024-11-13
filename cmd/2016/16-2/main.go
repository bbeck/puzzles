package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/puz"
)

const Size = 35651584

func main() {
	data := puz.InputToBytes(2016, 16)
	for len(data) < Size {
		data = Expand(data)
	}

	checksum := Checksum(data[:Size])
	fmt.Println(string(checksum))
}

func Expand(data []byte) []byte {
	flipped := map[byte]byte{
		'0': '1',
		'1': '0',
	}
	N := len(data)

	next := make([]byte, 2*N+1)
	next[N] = '0'
	for i := 0; i < N; i++ {
		next[i] = data[i]
		next[2*N-i] = flipped[data[i]]
	}
	return next
}

func Checksum(data []byte) []byte {
	checksum := map[bool]byte{
		false: '0',
		true:  '1',
	}

	for len(data)%2 == 0 {
		for i := 0; i < len(data); i += 2 {
			data[i/2] = checksum[data[i] == data[i+1]]
		}
		data = data[:len(data)/2]
	}

	return data
}
