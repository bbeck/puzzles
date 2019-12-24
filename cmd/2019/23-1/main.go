package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	var cpus []*CPU
	var inputs []chan int
	for i := 0; i < 50; i++ {
		input := make(chan int, 1024)
		input <- i // Seed with the network address
		inputs = append(inputs, input)

		var buffer []int

		cpu := &CPU{
			memory: InputToMemory(2019, 23),
			input: func(addr int) int {
				select {
				case value := <-input:
					return value
				default:
					return -1
				}
			},
			output: func(value int) {
				buffer = append(buffer, value)
				if len(buffer) == 3 {
					destination, x, y := buffer[0], buffer[1], buffer[2]
					buffer = nil

					if destination == 255 {
						fmt.Printf("writing to 255: x=%d, y=%d\n", x, y)
						close(c)
						return
					}

					inputs[destination] <- x
					inputs[destination] <- y
				}
			},
		}
		cpus = append(cpus, cpu)
	}
	for i := 0; i < 50; i++ {
		go cpus[i].Execute()
	}

	<-c
}
