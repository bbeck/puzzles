package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/aoc"
	"log"
	"math"
	"strings"
)

func main() {
	bs := InputToBitStream()
	packet := ReadPacket(bs)
	fmt.Println(packet.Eval())
}

func InputToBitStream() *BitStream {
	input := aoc.InputToString(2021, 16)

	// There's a newline at the end of the file that needs to be stripped
	input = strings.TrimSpace(input)

	// Convert the input into an array of integers.  This will handle the
	// conversion of the data from strings of hexadecimal digits to numbers.  In
	// order to produce data in 8-bit chunks we'll read two digits at a time.
	var data []uint8
	for i := 0; i < len(input); i += 2 {
		n := aoc.ParseIntWithBase(input[i:i+2], 16)
		data = append(data, uint8(n))
	}

	return NewBitStream(data)
}

type Packet struct {
	Kind string
	V, T int

	// data for literal packets
	Value int

	// data for operator packets
	Mode     int
	Children []Packet
}

func (p Packet) Eval() int {
	if p.Kind == "literal" {
		return p.Value
	}

	if p.T == 0 {
		var sum int
		for _, c := range p.Children {
			sum += c.Eval()
		}
		return sum
	}

	if p.T == 1 {
		var prod = 1
		for _, c := range p.Children {
			prod = prod * c.Eval()
		}
		return prod
	}

	if p.T == 2 {
		var min = math.MaxInt
		for _, c := range p.Children {
			min = aoc.MinInt(min, c.Eval())
		}
		return min
	}

	if p.T == 3 {
		var max = 0
		for _, c := range p.Children {
			max = aoc.MaxInt(max, c.Eval())
		}
		return max
	}

	if p.T == 5 {
		if p.Children[0].Eval() > p.Children[1].Eval() {
			return 1
		}
		return 0
	}

	if p.T == 6 {
		if p.Children[0].Eval() < p.Children[1].Eval() {
			return 1
		}
		return 0
	}

	if p.T == 7 {
		if p.Children[0].Eval() == p.Children[1].Eval() {
			return 1
		}
		return 0
	}

	log.Fatalf("unexpected packet type T: %d", p.T)
	return 0
}

func ReadPacket(bs *BitStream) *Packet {
	if bs.Size() < 10 {
		return nil
	}

	V := bs.ReadNBits(3)
	T := bs.ReadNBits(3)

	if T == 4 {
		// Literal packets contain just an immediate value
		return &Packet{
			Kind:  "literal",
			V:     V,
			T:     T,
			Value: bs.ReadVarInt(),
		}
	} else {
		// Operator packets contain an operator and sub-packets
		var subs []Packet

		mode := bs.ReadBit()
		if mode == 0 {
			// mode 0: the next 15 bits represent the total length (in bits) of the sub-packets
			nbits := bs.ReadNBits(15)

			// Setup a sub-bit-stream that contains the next nbits of data
			sbs := bs.Skip(nbits)

			for {
				p := ReadPacket(sbs)
				if p == nil {
					break
				}

				subs = append(subs, *p)
			}
		} else {
			// mode 1: the next 11 bits represent the number of sub-packets
			num := bs.ReadNBits(11)

			for i := 0; i < num; i++ {
				subs = append(subs, *ReadPacket(bs))
			}
		}

		return &Packet{
			Kind:     "operator",
			V:        V,
			T:        T,
			Mode:     mode,
			Children: subs,
		}
	}
}

type BitStream struct {
	data []uint8
	dp   int // which integer in data we're on
	bp   int // which bit in the current integer we're on
}

func NewBitStream(data []uint8) *BitStream {
	return &BitStream{
		data: data,
		dp:   0,
		bp:   7,
	}
}

// Size returns the number of bits of data remaining in the stream.
func (bs BitStream) Size() int {
	// Number of full bytes of data remaining
	bytes := len(bs.data) - bs.dp - 1

	// Number of bits remaining in the current byte
	bits := bs.bp + 1

	return 8*bytes + bits
}

func (bs *BitStream) ReadBit() int {
	bit := bs.data[bs.dp] & (1 << bs.bp) >> bs.bp
	bs.bp--
	if bs.bp < 0 {
		bs.bp = 7
		bs.dp++
	}

	return int(bit)
}

func (bs *BitStream) ReadNBits(N int) int {
	var n int
	for i := 0; i < N; i++ {
		n = n<<1 | bs.ReadBit()
	}

	return n
}

func (bs *BitStream) ReadVarInt() int {
	var n int
	for {
		more := bs.ReadBit() != 0
		n = n<<4 | bs.ReadNBits(4)

		if !more {
			break
		}
	}

	return n
}

// Skip will skip over the leading n bits from the bit stream and return the skipped bits
// as a new bit stream.
func (bs *BitStream) Skip(size int) *BitStream {
	var data []uint8
	for size >= 8 {
		n := bs.ReadNBits(8)
		data = append(data, uint8(n))
		size -= 8
	}

	if size > 0 {
		n := bs.ReadNBits(size) << (8 - size)
		data = append(data, uint8(n))
	}

	return NewBitStream(data)
}
