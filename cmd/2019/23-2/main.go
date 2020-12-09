package main

import (
	"fmt"
	"time"

	aoc "github.com/bbeck/advent-of-code/aoc/cpus"
)

func main() {
	var cpus []*aoc.IntcodeCPU
	inputs := make(map[int]chan int)
	for i := 0; i < 50; i++ {
		input := make(chan int, 1024)
		input <- i // Seed with the network address
		inputs[i] = input

		var buffer []int

		cpu := aoc.IntcodeCPU{
			Memory: aoc.InputToIntcodeMemory(2019, 23),
			Input: func(addr int) int {
				select {
				case value := <-input:
					return value
				default:
					return -1
				}
			},
			Output: func(value int) {
				buffer = append(buffer, value)
				if len(buffer) == 3 {
					destination, x, y := buffer[0], buffer[1], buffer[2]
					buffer = nil

					inputs[destination] <- x
					inputs[destination] <- y
				}
			},
		}
		cpus = append(cpus, &cpu)
	}

	inputs[255] = make(chan int)

	for i := 0; i < len(cpus); i++ {
		go cpus[i].Execute()
	}

	var natBuffer []int
	var lastY int

outer:
	for {
		select {
		case value := <-inputs[255]:
			natBuffer = append(natBuffer, value)
			if len(natBuffer) == 4 {
				natBuffer = natBuffer[2:]
			}
		case <-time.After(200 * time.Millisecond):
			if len(natBuffer) >= 2 {
				x, y := natBuffer[0], natBuffer[1]
				natBuffer = natBuffer[2:]

				inputs[0] <- x
				inputs[0] <- y

				if lastY == y {
					fmt.Printf("double y: %d\n", y)
					break outer
				}
				lastY = y
			}
		}
	}
}
