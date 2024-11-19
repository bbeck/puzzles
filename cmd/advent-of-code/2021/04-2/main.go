package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/lib"
	"strings"
)

const N = 5

func main() {
	nums, boards := InputToGame()

	var called lib.Set[int]
	var winners lib.Set[int]
	var winner Board
	var last int

outer:
	for _, n := range nums {
		called.Add(n)
		last = n

		for id, board := range boards {
			if winners.Contains(id) {
				continue
			}
			for _, line := range board.Lines {
				if len(called.Intersect(line)) == N {
					winner = board
					winners.Add(id)
					if len(winners) == len(boards) {
						break outer
					}
				}
			}
		}
	}

	var all lib.Set[int]
	for _, line := range winner.Lines {
		all = all.Union(line)
	}

	remaining := all.Difference(called)
	fmt.Println(last * lib.Sum(remaining.Entries()...))
}

type Board struct {
	Lines []lib.Set[int]
}

func InputToGame() ([]int, []Board) {
	lines := lib.InputToLines()

	var nums []int
	for _, s := range strings.Split(lines[0], ",") {
		nums = append(nums, lib.ParseInt(s))
	}

	var boards []Board
	for base := 2; base < len(lines); base += 6 {
		board := Board{Lines: make([]lib.Set[int], 2*N)}
		for y := 0; y < N; y++ {
			for x, s := range strings.Fields(lines[base+y]) {
				n := lib.ParseInt(s)
				board.Lines[y].Add(n)
				board.Lines[x+N].Add(n)
			}
		}

		boards = append(boards, board)
	}
	return nums, boards
}

func ParseLine(line string) []int {
	var ns []int
	for _, s := range strings.Fields(line) {
		ns = append(ns, lib.ParseInt(s))
	}
	return ns
}
