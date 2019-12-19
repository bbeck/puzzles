package main

import (
	"fmt"
)

func main() {
	var sum int
	for y := 0; y < 50; y++ {
		for x := 0; x < 50; x++ {
			if Affected(x, y) {
				sum++
			}
		}
	}

	fmt.Println("num affected:", sum)
}

func Affected(x, y int) bool {
	inputs := make(chan int, 2)
	inputs <- x
	inputs <- y
	close(inputs)

	var output int
	cpu := &CPU{
		memory: InputToMemory(2019, 19),
		input:  func(addr int) int { return <-inputs },
		output: func(value int) { output = value },
	}
	cpu.Execute()

	return output == 1
}
