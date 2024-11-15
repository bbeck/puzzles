package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/puz"
	"github.com/bbeck/advent-of-code/puz/cpus"
	"sync"
)

const N = 50

func main() {
	nat := NewNAT()

	var computers [N]cpus.IntcodeCPU
	for addr := 0; addr < N; addr++ {
		nat.Send(addr, addr) // initialize with address of each computer
		computers[addr] = *NewComputer(addr, nat)
	}

	for addr := 0; addr < N; addr++ {
		go computers[addr].Execute()
	}
	fmt.Println(nat.Wait())

	for addr := 0; addr < N; addr++ {
		computers[addr].Stop()
	}
}

type NAT struct {
	Mutex          sync.Mutex
	Buffers        [N]puz.Deque[int]
	Idle           puz.BitSet
	ToForward      []int // The last received packet that can be forwarded
	LastForwardedY int
	Duplicates     chan int
}

func NewNAT() *NAT {
	var nat NAT
	nat.Duplicates = make(chan int)

	return &nat
}

func (n *NAT) Send(addr int, data ...int) {
	n.Mutex.Lock()
	defer n.Mutex.Unlock()

	if addr == 255 {
		n.ToForward = data
		return
	}

	for _, d := range data {
		n.Buffers[addr].PushBack(d)
	}
	n.Idle = n.Idle.Remove(addr)
}

func (n *NAT) Receive(addr int) int {
	n.Mutex.Lock()
	defer n.Mutex.Unlock()

	if addr == 0 && n.Buffers[0].Len() == 0 && n.Idle.Size() == N {
		x, y := n.ToForward[0], n.ToForward[1]
		n.Buffers[0].PushBack(x)
		n.Buffers[0].PushBack(y)

		if y == n.LastForwardedY {
			n.Duplicates <- y
		}
		n.LastForwardedY = y
	}

	if n.Buffers[addr].Len() == 0 {
		n.Idle = n.Idle.Add(addr)
		return -1
	}

	n.Idle = n.Idle.Remove(addr)
	return n.Buffers[addr].PopFront()
}

func (n *NAT) Wait() int {
	return <-n.Duplicates
}

func NewComputer(id int, nat *NAT) *cpus.IntcodeCPU {
	var buffer []int

	return &cpus.IntcodeCPU{
		Memory: cpus.InputToIntcodeMemory(),
		Input:  func() int { return nat.Receive(id) },
		Output: func(value int) {
			buffer = append(buffer, value)
			if len(buffer) == 3 {
				addr, x, y := buffer[0], buffer[1], buffer[2]
				buffer = nil

				nat.Send(addr, x, y)
			}
		},
	}
}
