package main

import (
	"fmt"

	"github.com/bbeck/advent-of-code/aoc"
)

type Configuration struct {
	numRows int
	numCols int
	lights  []bool
}

func (c *Configuration) Get(row, col int) bool {
	return c.lights[row*c.numCols+col]
}

func (c *Configuration) Set(row, col int, state bool) {
	c.lights[row*c.numCols+col] = state
}

func (c *Configuration) Next() *Configuration {
	next := &Configuration{
		numRows: c.numRows,
		numCols: c.numCols,
		lights:  make([]bool, len(c.lights)),
	}

	on := func(row, col int) int {
		if row >= 0 && row < c.numRows && col >= 0 && col < c.numCols {
			if c.Get(row, col) {
				return 1
			}
		}

		return 0
	}

	for row := 0; row < c.numRows; row++ {
		for col := 0; col < c.numCols; col++ {
			count := on(row-1, col-1) + on(row-1, col) + on(row-1, col+1) +
				on(row, col-1) + on(row, col+1) +
				on(row+1, col-1) + on(row+1, col) + on(row+1, col+1)

			if c.Get(row, col) {
				next.Set(row, col, count == 2 || count == 3)
			} else {
				next.Set(row, col, count == 3)
			}
		}
	}

	return next
}

func (c *Configuration) Count() int {
	var count int
	for row := 0; row < c.numRows; row++ {
		for col := 0; col < c.numCols; col++ {
			if c.Get(row, col) {
				count++
			}
		}
	}

	return count
}

func (c *Configuration) Print() {
	for row := 0; row < c.numRows; row++ {
		for col := 0; col < c.numCols; col++ {
			if c.Get(row, col) {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func main() {
	config := InputToLightConfiguration(2015, 18)

	for i := 1; i <= 100; i++ {
		config = config.Next()
	}

	fmt.Printf("num lights on: %d\n", config.Count())
}

func InputToLightConfiguration(year, day int) *Configuration {
	lines := make([]string, 0)
	for _, line := range aoc.InputToLines(year, day) {
		lines = append(lines, line)
	}

	N := len(lines)

	lights := make([]bool, N*N)
	for row := 0; row < N; row++ {
		for col := 0; col < N; col++ {
			lights[row*N+col] = lines[row][col] == '#'
		}
	}

	return &Configuration{
		numRows: N,
		numCols: N,
		lights:  lights,
	}
}
