package main

import (
	"fmt"

	. "github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	instructions := InputToInstructions()
	var deck Deck
	for i := 0; i < 10007; i++ {
		deck.PushBack(i)
	}

	for _, instruction := range instructions {
		switch instruction.Kind {
		case DealNewStack:
			deck.DealIntoNewStack()
		case Cut:
			deck.Cut(instruction.Arg)
		case DealWithIncrement:
			deck.DealWithIncrement(instruction.Arg)
		}
	}

	var position int
	for deck.PeekFront() != 2019 {
		deck.Rotate(-1)
		position++
	}
	fmt.Println(position)
}

type Deck struct {
	Deque[int]
}

func (d *Deck) DealIntoNewStack() {
	var next Deck
	for d.Len() > 0 {
		next.PushBack(d.PopBack())
	}
	*d = next
}

func (d *Deck) Cut(n int) {
	d.Rotate(-n)
}

func (d *Deck) DealWithIncrement(n int) {
	cards := make([]int, d.Len())

	var index int
	for d.Len() > 0 {
		cards[index] = d.PopFront()
		index = (index + n) % len(cards)
	}

	var next Deck
	for _, c := range cards {
		next.PushBack(c)
	}
	*d = next
}

const (
	DealNewStack      = "new_stack"
	Cut               = "cut"
	DealWithIncrement = "increment"
)

type Instruction struct {
	Kind string
	Arg  int
}

func InputToInstructions() []Instruction {
	return in.LinesToS(func(in in.Scanner[Instruction]) Instruction {
		switch {
		case in.HasPrefix("deal into new stack"):
			return Instruction{Kind: DealNewStack}
		case in.HasPrefix("cut"):
			var arg int
			in.Scanf("cut %d", &arg)
			return Instruction{Kind: Cut, Arg: arg}
		case in.HasPrefix("deal with increment"):
			var arg int
			in.Scanf("deal with increment %d", &arg)
			return Instruction{Kind: DealWithIncrement, Arg: arg}
		default:
			panic("unsupported prefix")
		}
	})
}
