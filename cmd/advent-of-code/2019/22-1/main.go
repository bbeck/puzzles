package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/lib"
	"log"
	"strings"
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
	lib.Deque[int]
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
	return lib.InputLinesTo(func(line string) Instruction {
		if strings.HasPrefix(line, "deal into new stack") {
			return Instruction{Kind: DealNewStack}
		} else if strings.HasPrefix(line, "cut") {
			arg := lib.ParseInt(strings.Split(line, " ")[1])
			return Instruction{Kind: Cut, Arg: arg}
		} else if strings.HasPrefix(line, "deal with increment") {
			arg := lib.ParseInt(strings.Split(line, " ")[3])
			return Instruction{Kind: DealWithIncrement, Arg: arg}
		} else {
			log.Fatalf("unrecognized line: %s", line)
			return Instruction{}
		}
	})
}
