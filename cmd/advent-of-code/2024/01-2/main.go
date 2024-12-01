package main

import (
	"fmt"
	"strings"

	. "github.com/bbeck/puzzles/lib"
)

func main() {
	as, bs := InputToNums()

	var sum int
	for a := range as {
		sum += a * as[a] * bs[a]
	}
	fmt.Println(sum)
}

func InputToNums() (map[int]int, map[int]int) {
	as, bs := make(map[int]int), make(map[int]int)
	for _, line := range InputToLines() {
		fields := strings.Fields(line)
		as[ParseInt(fields[0])]++
		bs[ParseInt(fields[1])]++
	}
	return as, bs
}
