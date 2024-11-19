package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/lib"
	"log"
	"strings"
)

func main() {
	var sum int
	for _, row := range InputToRows() {
		n, d := FindEvenlyDivisible(row)
		sum += n / d
	}
	fmt.Println(sum)
}

func FindEvenlyDivisible(row []int) (int, int) {
	for i := 0; i < len(row); i++ {
		for j := 0; j < len(row); j++ {
			if i == j {
				continue
			}

			if row[i]%row[j] == 0 {
				return row[i], row[j]
			}
		}
	}

	log.Fatalf("unable to find evenly divisible pair: %v", row)
	return 0, 0
}

func InputToRows() [][]int {
	return lib.InputLinesTo(func(line string) []int {
		var row []int
		for _, s := range strings.Fields(line) {
			row = append(row, lib.ParseInt(s))
		}

		return row
	})
}
