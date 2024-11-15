package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/puz"
	"strings"
)

const N = 5

func main() {
	nums, boards := InputToGame()

	var called puz.Set[int]
	var winner Board
	var last int

outer:
	for _, n := range nums {
		called.Add(n)
		last = n

		for _, board := range boards {
			for _, line := range board.Lines {
				if len(called.Intersect(line)) == N {
					winner = board
					break outer
				}
			}
		}
	}

	var all puz.Set[int]
	for _, line := range winner.Lines {
		all = all.Union(line)
	}

	remaining := all.Difference(called)
	fmt.Println(last * puz.Sum(remaining.Entries()...))
}

type Board struct {
	Lines []puz.Set[int]
}

func InputToGame() ([]int, []Board) {
	lines := puz.InputToLines()

	var nums []int
	for _, s := range strings.Split(lines[0], ",") {
		nums = append(nums, puz.ParseInt(s))
	}

	var boards []Board
	for base := 2; base < len(lines); base += 6 {
		board := Board{Lines: make([]puz.Set[int], 2*N)}
		for y := 0; y < N; y++ {
			for x, s := range strings.Fields(lines[base+y]) {
				n := puz.ParseInt(s)
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
		ns = append(ns, puz.ParseInt(s))
	}
	return ns
}
