package main

import (
	"fmt"
	. "github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/in"
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
		for j := i + 1; j < len(row); j++ {
			a, b := Max(row[i], row[j]), Min(row[i], row[j])
			if a%b == 0 {
				return a, b
			}
		}
	}

	panic("unable to find evenly divisible pair")
}

func InputToRows() [][]int {
	return in.LinesToS(func(in in.Scanner[[]int]) []int {
		return in.Ints()
	})
}
