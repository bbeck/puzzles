package main

import (
	"fmt"

	. "github.com/bbeck/puzzles/lib"
)

func main() {
	locks, keys := InputToLocksAndKeys()

	var count int
	for _, lock := range locks {
		for _, key := range keys {
			if Fits(key, lock) {
				count++
			}
		}
	}
	fmt.Println(count)
}

func Fits(key, lock []int) bool {
	for i := 0; i < len(key); i++ {
		if lock[i]+key[i] >= 6 {
			return false
		}
	}
	return true
}

func InputToLocksAndKeys() ([][]int, [][]int) {
	var locks, keys [][]int

	for _, chunk := range Chunk(InputToLines(), 8) {
		var current []int
		for x := 0; x < 5; x++ {
			var count int
			for y := 0; y < 7; y++ {
				if chunk[y][x] == '#' {
					count++
				}
			}
			current = append(current, count-1)
		}

		if chunk[0] == "#####" {
			locks = append(locks, current)
		} else {
			keys = append(keys, current)
		}
	}

	return locks, keys
}
