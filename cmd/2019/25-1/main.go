package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/bbeck/advent-of-code/aoc/cpus"
)

func main() {
	buffer := make(chan int, 1024)

	cpu := cpus.IntcodeCPU{
		Memory: cpus.InputToIntcodeMemory(2019, 25),
		Input: func(addr int) int {
			return <-buffer
		},
		Output: func(value int) {
			fmt.Printf("%c", value)
		},
	}
	go cpu.Execute()

	reader := bufio.NewReader(os.Stdin)
	for {
		input, _ := reader.ReadString('\n')
		for _, c := range input {
			buffer <- int(c)
		}
	}
}
