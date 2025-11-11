package main

import (
	"fmt"

	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	var mentor, sum int
	for in.HasNext() {
		switch in.Byte() {
		case 'A':
			mentor++
		case 'a':
			sum += mentor
		}
	}
	fmt.Println(sum)
}
