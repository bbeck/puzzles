package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	buffer := make(chan int, 1024)
	input := func(addr int) int {
		return <-buffer
	}

	cpu := &CPU{
		memory: InputToMemory(2019, 25),
		input:  input,
		output: func(value int) {
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
