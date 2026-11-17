package main

import (
	"bytes"
	"fmt"

	. "github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	bs := in.Bytes()

	var count int
	for len(bs) > 0 {
		var start int
		for start < len(bs)-1 && bs[start+1] == bs[start] {
			start++
		}

		end := bytes.LastIndexByte(bs, bs[start])
		if start == end {
			bs = bs[start+1:]
			continue
		}

		Reverse(bs[start+1 : end+1])
		count++
	}
	fmt.Println(count)
}
