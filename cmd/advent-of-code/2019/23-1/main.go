package main

import (
	"fmt"
	"github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/cpus"
	"runtime"
)

func main() {
	inputs := map[int]chan lib.Point2D{255: make(chan lib.Point2D)}
	var computers []*cpus.IntcodeCPU

	for n := 0; n < 50; n++ {
		inputs[n] = make(chan lib.Point2D)
		computers = append(computers, NewComputer(n, inputs))
	}

	for _, computer := range computers {
		go computer.Execute()
	}

	p := <-inputs[255]
	for _, computer := range computers {
		computer.Stop()
	}

	fmt.Println(p.Y)
}

func NewComputer(id int, inputs map[int]chan lib.Point2D) *cpus.IntcodeCPU {
	var buffer []int

	var hasSentId bool
	var point *lib.Point2D

	return &cpus.IntcodeCPU{
		Memory: cpus.InputToIntcodeMemory(),
		Input: func() int {
			if !hasSentId {
				hasSentId = true
				return id
			}

			if point != nil {
				y := point.Y
				point = nil
				return y
			}

			select {
			case value := <-inputs[id]:
				point = &value
				return point.X
			default:
				// yield the cpu when there's no data available so that another
				// goroutine can hopefully generate data for this CPU.  While not
				// strictly necessary, this drastically reduces wasted cycles and makes
				// things run faster overall.
				runtime.Gosched()
				return -1
			}
		},
		Output: func(value int) {
			buffer = append(buffer, value)
			if len(buffer) == 3 {
				addr, x, y := buffer[0], buffer[1], buffer[2]
				buffer = nil

				inputs[addr] <- lib.Point2D{X: x, Y: y}
			}
		},
	}
}
