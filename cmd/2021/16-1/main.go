package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/aoc"
	"strings"
)

func main() {
	bs := InputToBitStream()
	packet := ReadPacket(bs)

	var walk func(p Packet, fn func(Packet))
	walk = func(p Packet, fn func(Packet)) {
		fn(p)
		for _, c := range p.Children {
			walk(c, fn)
		}
	}

	var sum int
	walk(*packet, func(p Packet) {
		sum += p.V
	})
	fmt.Println(sum)
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

func (p Packet) String() string {
	indent := func(n int) string {
		var sb strings.Builder
		for i := 0; i < n; i++ {
			sb.WriteRune(' ')
		}
		return sb.String()
	}

	var helper func(ind int, p Packet) string
	helper = func(ind int, p Packet) string {
		if p.Kind == "literal" {
			return indent(ind) + fmt.Sprintf("Literal: %d", p.Value)
		}

		var sb strings.Builder
		sb.WriteString(indent(ind) + "Packet{\n")
		sb.WriteString(indent(ind) + fmt.Sprintf("  Kind: %s\n", p.Kind))
		sb.WriteString(indent(ind) + fmt.Sprintf("  V: %d\n", p.V))
		sb.WriteString(indent(ind) + fmt.Sprintf("  T: %d\n", p.T))

		sb.WriteString(indent(ind) + fmt.Sprintf("  Mode: %d\n", p.Mode))
		sb.WriteString(indent(ind) + "  Children: [\n")
		for _, c := range p.Children {
			sb.WriteString(helper(ind+4, c) + "\n")
		}
		sb.WriteString(indent(ind) + "  ]\n")
		sb.WriteString(indent(ind) + "}\n")
		return sb.String()
	}

	return helper(0, p)
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

func (bs BitStream) String() string {
	var sb strings.Builder

	sb.WriteString("data: [")
	for i, n := range bs.data {
		sb.WriteString(fmt.Sprintf("%b", n))
		if i < len(bs.data)-1 {
			sb.WriteString(" ")
		}
	}
	sb.WriteString("], ")
	sb.WriteString(fmt.Sprintf("dp: %d, bp: %d", bs.dp, bs.bp))

	return sb.String()
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
