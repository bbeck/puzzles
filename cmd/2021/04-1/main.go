package main

import (
	"fmt"
	"strings"

	"github.com/bbeck/advent-of-code/aoc"
)

const N = 5

func main() {
	numbers, cards := InputToBingoGame()
	finished := make([]bool, len(cards))

	var scores []int
	for _, n := range numbers {
		for i, c := range cards {
			if !finished[i] && c.Cover(n) {
				scores = append(scores, c.Score()*n)
				finished[i] = true
			}
		}
	}

	fmt.Println(scores[0])
}

type Card struct {
	numbers [N][N]int
	covered [N][N]bool

	// counts of how many numbers are covered in each row/column.
	rows [N]int
	cols [N]int
}

func (c *Card) Cover(n int) bool {
	for row := 0; row < N; row++ {
		for col := 0; col < N; col++ {
			if c.numbers[row][col] == n {
				c.covered[row][col] = true
				c.rows[row]++
				c.cols[col]++

				return c.rows[row] == N || c.cols[col] == N
			}
		}
	}

	return false
}

func (c *Card) Score() int {
	var score int
	for row := 0; row < N; row++ {
		for col := 0; col < N; col++ {
			if !c.covered[row][col] {
				score += c.numbers[row][col]
			}
		}
	}

	return score
}

func InputToBingoGame() ([]int, []*Card) {
	lines := aoc.InputToLines(2021, 4)

	// First line is a comma separated list of the numbers that are going to be called.
	var numbers []int
	for _, s := range strings.Split(lines[0], ",") {
		numbers = append(numbers, aoc.ParseInt(s))
	}

	// Remaining lines are the cards, each card is 5 lines followed by a blank line.
	var cards []*Card
	for i := 2; i < len(lines); i += N + 1 {
		cards = append(cards, &Card{
			numbers: [N][N]int{
				ParseRow(lines[i]),
				ParseRow(lines[i+1]),
				ParseRow(lines[i+2]),
				ParseRow(lines[i+3]),
				ParseRow(lines[i+4]),
			},
		})
	}

	return numbers, cards
}

func ParseRow(s string) [5]int {
	// Single digit numbers are padded with a space
	s = strings.ReplaceAll(strings.TrimSpace(s), "  ", " ")

	var ns [5]int
	for i, s := range strings.Split(s, " ") {
		ns[i] = aoc.ParseInt(s)
	}
	return ns
}
