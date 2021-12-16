package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/aoc"
	"log"
	"math"
	"strings"
)

func main() {
	packet := InputToPacket()
	fmt.Println(packet.Eval())
}

type Packet struct {
	kind string
	V, T int

	// data for literal packets
	literal int

	// data for operator packets
	children []Packet
}

func (p Packet) Eval() int {
	if p.kind == "literal" {
		return p.literal
	}

	if p.T == 0 {
		var sum int
		for _, c := range p.children {
			sum += c.Eval()
		}
		return sum
	}

	if p.T == 1 {
		var prod = 1
		for _, c := range p.children {
			prod = prod * c.Eval()
		}
		return prod
	}

	if p.T == 2 {
		var min = math.MaxInt
		for _, c := range p.children {
			min = aoc.MinInt(min, c.Eval())
		}
		return min
	}

	if p.T == 3 {
		var max = 0
		for _, c := range p.children {
			max = aoc.MaxInt(max, c.Eval())
		}
		return max
	}

	if p.T == 5 {
		if p.children[0].Eval() > p.children[1].Eval() {
			return 1
		}
		return 0
	}

	if p.T == 6 {
		if p.children[0].Eval() < p.children[1].Eval() {
			return 1
		}
		return 0
	}

	if p.T == 7 {
		if p.children[0].Eval() == p.children[1].Eval() {
			return 1
		}
		return 0
	}

	log.Fatalf("unexpected packet type T: %d", p.T)
	return 0
}

func InputToPacket() Packet {
	var bits []int
	for _, c := range strings.TrimSpace(aoc.InputToString(2021, 16)) {
		n := aoc.ParseIntWithBase(string(c), 16)
		bs := ToBits(n)
		bits = append(bits, bs...)
	}

	bits, p := ReadPacket(bits)
	return p
}

func ReadPacket(bits []int) ([]int, Packet) {
	var V, T int
	bits, V = ReadNBits(bits, 3)
	bits, T = ReadNBits(bits, 3)

	if T == 4 {
		return ReadLiteralPacket(bits, V)
	}

	return ReadOperatorPacket(bits, T, V)
}

func ReadLiteralPacket(bits []int, V int) ([]int, Packet) {
	var n int
	bits, n = ReadVarInt(bits)
	return bits, Packet{
		kind:    "literal",
		V:       V,
		literal: n,
	}
}

func ReadVarInt(bits []int) ([]int, int) {
	var N int
	var more, n int
	for {
		bits, more = ReadNBits(bits, 1)
		bits, n = ReadNBits(bits, 4)
		N = N<<4 | n

		if more == 0 {
			break
		}
	}

	return bits, N
}

func ReadNBits(bits []int, N int) ([]int, int) {
	var n int
	for i := 0; i < N; i++ {
		n = n<<1 | bits[0]
		bits = bits[1:]
	}
	return bits, n
}

func ReadOperatorPacket(bits []int, T int, V int) ([]int, Packet) {
	var tid int
	var children []Packet

	bits, tid = ReadNBits(bits, 1)
	if tid == 0 {
		var length int
		bits, length = ReadNBits(bits, 15)

		sbits := bits[:length]
		bits = bits[length:]

		for len(sbits) > 6 {
			var p Packet
			sbits, p = ReadPacket(sbits)
			children = append(children, p)
		}
	} else {
		var num int
		bits, num = ReadNBits(bits, 11)

		var p Packet
		for i := 0; i < num; i++ {
			bits, p = ReadPacket(bits)
			children = append(children, p)
		}
	}

	packet := Packet{
		kind:     "operator",
		V:        V,
		T:        T,
		children: children,
	}

	return bits, packet
}

func LiteralValue(data []int) int {
	var n int
	for _, d := range data {
		n = (n << 4) | d
	}
	return n
}

func ToBits(n int) []int {
	var bits []int
	for i := 7; i >= 0; i-- {
		if n&(1<<i) == 0 {
			bits = append(bits, 0)
		} else {
			bits = append(bits, 1)
		}
	}
	return bits[4:]
}
