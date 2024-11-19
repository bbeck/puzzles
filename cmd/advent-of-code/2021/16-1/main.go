package main

import (
	"fmt"
	"github.com/bbeck/puzzles/lib"
)

func main() {
	packet := ReadPacket(InputToBitStream())

	var walk func(*Packet, func(*Packet))
	walk = func(p *Packet, fn func(*Packet)) {
		fn(p)
		for _, child := range p.Children {
			walk(child, fn)
		}
	}

	var sum int
	walk(packet, func(p *Packet) {
		sum += p.Version
	})
	fmt.Println(sum)
}

type Packet struct {
	Version, Type int
	Value         int       // Literal packets
	Children      []*Packet // Operator packets
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
	for _, c := range lib.InputToString() {
		n := lib.ParseIntWithBase(string(c), 16)
		for mask := 0b1000; mask > 0; mask >>= 1 {
			bits = append(bits, n&mask == mask)
		}
	}

	return &bits
}
