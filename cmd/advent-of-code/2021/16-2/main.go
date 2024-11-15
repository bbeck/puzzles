package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/puz"
)

func main() {
	packet := ReadPacket(InputToBitStream())
	fmt.Println(packet.Eval())
}

type Packet struct {
	Version, Type int
	Value         int       // Literal packets
	Children      []*Packet // Operator packets
}

func (p *Packet) Eval() int {
	var children []int
	for _, child := range p.Children {
		children = append(children, child.Eval())
	}

	switch p.Type {
	case 0: // sum
		return puz.Sum(children...)
	case 1: // product
		return puz.Product(children...)
	case 2: // minimum
		return puz.Min(children...)
	case 3: // maximum
		return puz.Max(children...)
	case 4: // literal
		return p.Value
	case 5: // greater than
		if children[0] > children[1] {
			return 1
		}
	case 6: // less than
		if children[0] < children[1] {
			return 1
		}
	case 7: // equal to
		if children[0] == children[1] {
			return 1
		}
	}

	return 0
}

func ReadPacket(bits *Bits) *Packet {
	// We need at least 10 bits to form a packet, anything less is padding
	if len(*bits) < 10 {
		return nil
	}

	version := bits.Read(3)
	typeID := bits.Read(3)

	if typeID == 4 { // Literal packet
		return &Packet{
			Version: version,
			Type:    typeID,
			Value:   ReadVarInt(bits),
		}
	}

	if typeID != 4 { // Operator packet
		var children []*Packet

		if bits.Read(1) == 0 { // next 15 bits are the length of the children
			length := bits.Read(15)
			sub := bits.Skip(length)

			for {
				child := ReadPacket(sub)
				if child == nil {
					break
				}

				children = append(children, child)
			}
		} else { // next 11 bits are the count of the children
			count := bits.Read(11)
			for i := 0; i < count; i++ {
				children = append(children, ReadPacket(bits))
			}
		}

		return &Packet{
			Version:  version,
			Type:     typeID,
			Children: children,
		}
	}

	return nil
}

func ReadVarInt(bits *Bits) int {
	var n int
	for {
		more := bits.Read(1)
		n = (n << 4) + bits.Read(4)

		if more == 0 {
			break
		}
	}
	return n
}

type Bits []bool

func (b *Bits) Read(n int) int {
	var value int
	for i := 0; i < n; i++ {
		value = value << 1
		if (*b)[i] {
			value |= 1
		}
	}
	*b = (*b)[n:]
	return value
}

func (b *Bits) Skip(n int) *Bits {
	var skipped Bits
	skipped = (*b)[:n]
	*b = (*b)[n:]
	return &skipped
}

func InputToBitStream() *Bits {
	var bits Bits
	for _, c := range puz.InputToString(2021, 16) {
		n := puz.ParseIntWithBase(string(c), 16)
		for mask := 0b1000; mask > 0; mask >>= 1 {
			bits = append(bits, n&mask == mask)
		}
	}

	return &bits
}
