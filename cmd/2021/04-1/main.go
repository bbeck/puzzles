package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/bbeck/advent-of-code/aoc"
)

const N = 5

func main() {
	numbers, cards := InputToBingo()

outer:
	for _, n := range numbers {
		for _, c := range cards {
			if c.Check(n) {
				// Compute the score, sum of all unmarked numbers times number just called
				var score int
				for y := 0; y < N; y++ {
					for x := 0; x < N; x++ {
						if !c.covered[y][x] {
							score += c.numbers[y][x]
						}
					}
				}
				score *= n
				fmt.Println(score)
				break outer
			}
		}
	}

}

type Card struct {
	numbers [][]int
	covered [][]bool
}

func (c *Card) Check(n int) bool {
	var covered bool
	for y := 0; y < N; y++ {
		for x := 0; x < N; x++ {
			if c.numbers[y][x] == n {
				c.covered[y][x] = true
				covered = true
			}
		}
	}

	if !covered {
		return false
	}

	// We covered a number, check for a victory
	for i := 0; i < N; i++ {
		if c.covered[i][0] && c.covered[i][1] && c.covered[i][2] && c.covered[i][3] && c.covered[i][4] {
			return true
		}
		if c.covered[0][i] && c.covered[1][i] && c.covered[2][i] && c.covered[3][i] && c.covered[4][i] {
			return true
		}
	}

	return false
}

func InputToBingo() ([]int, []Card) {
	lines := aoc.InputToLines(2021, 4)

	var numbers []int
	for _, s := range strings.Split(lines[0], ",") {
		numbers = append(numbers, aoc.ParseInt(s))
	}

	var cards []Card
	for i := 2; i < len(lines); i += N + 1 {
		cards = append(cards, Card{
			numbers: [][]int{
				Row(lines[i]),
				Row(lines[i+1]),
				Row(lines[i+2]),
				Row(lines[i+3]),
				Row(lines[i+4]),
			},
			covered: [][]bool{
				make([]bool, N),
				make([]bool, N),
				make([]bool, N),
				make([]bool, N),
				make([]bool, N),
			},
		})
	}

	return numbers, cards
}

func Row(s string) []int {
	var n1, n2, n3, n4, n5 int
	if _, err := fmt.Sscanf(s, "%d %d %d %d %d", &n1, &n2, &n3, &n4, &n5); err != nil {
		log.Fatal(err)
	}

	return []int{n1, n2, n3, n4, n5}
}
