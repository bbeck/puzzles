package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/lib"
	"log"
	"math/big"
	"strings"
)

// This reddit comment was very helpful in deriving the math:
// https://www.reddit.com/r/adventofcode/comments/ee0rqi/comment/fbnkaju
func main() {
	instructions := InputToInstructions()

	// The numbers we're working with along with intermediate values during
	// computations can get quite large and overflow an int64.  Because of this
	// we'll work with arbitrary precision integers.
	N := big.NewInt(119315717514047) // number of cards in the deck
	S := big.NewInt(101741582076661) // number of shuffles

	// By observing the examples we can see that a shuffle is uniquely determined
	// by the value of the first card and a delta value between cards.  The n-th
	// card in the deck can be computed by (first + n*delta) % N.
	first := big.NewInt(0)
	delta := big.NewInt(1)

	for _, instruction := range instructions {
		switch instruction.Kind {
		case DealNewStack:
			// Reverses the deck.  The last card becomes the first, and we change the
			// direction we move in.
			first.Sub(first, delta)
			delta.Neg(delta)

		case Cut:
			// Move arg cards forward.  Only the first card changes.
			var temp big.Int
			temp.Mul(big.NewInt(instruction.Arg), delta)
			first.Add(first, &temp)

		case DealWithIncrement:
			// The first element always remains the same, the second element moves to
			// the arg-th element.  The third element moves to the 2*arg-th element,
			// and so on.  So if we can figure out which element becomes the new
			// second element we can compute the new delta.  Thus, we want to find x
			// in x*arg = 1 (mod N).  Thus x = inverse(arg) (mod N).
			//
			// The value of the x-th element in our list is first + x*delta, and
			// subtracting the first element from that gives us a new delta of
			// x*delta or inverse(arg)*delta.
			var temp big.Int
			temp.ModInverse(big.NewInt(instruction.Arg), N)
			delta.Mul(delta, &temp)
		}

		first.Mod(first, N)
		delta.Mod(delta, N)
	}

	// Now that we've computed one round of the shuffling, use the result to
	// determine the result after many more rounds of shuffling.
	//
	// Viewing the single round of shuffling algorithm we can draw two
	// conclusions about the effects of a round of shuffling:
	//   1. delta is always some integer multiple of the starting delta (dd)
	//   2. first is always incremented by some multiple of delta (df)
	//
	// Extrapolating to represent many shuffles (S) from a factory ordered deck:
	//   delta(s) = dd*delta(s-1)              ; delta(0) = 1
	//   first(s) = first(s-1) + df*delta(s-1) ; first(0) = 0
	//
	// From this we can conclude that deltas are computed via exponentiation, and
	// first is just a geometric series.
	//   delta(s) = (dd)^s
	//   first(s) = df*(1 + dd + (dd)^2 + ... + (dd)^s) = df*(1-(dd)^s)/(1-dd)

	var deltaS, firstS big.Int
	deltaS.Exp(delta, S, N)
	firstS = GeometricSum(first, delta, S, N)

	// Lastly determine the 2020th card in the deck.
	var card big.Int
	card.Mod(card.Add(&firstS, card.Mul(big.NewInt(2020), &deltaS)), N)
	fmt.Println(card.Int64())
}

func GeometricSum(a, r, n, mod *big.Int) big.Int {
	var numerator big.Int
	numerator.Mod(numerator.Sub(big.NewInt(1), numerator.Exp(r, n, mod)), mod)

	var denominator big.Int
	denominator.Mod(denominator.Sub(big.NewInt(1), r), mod)

	var result big.Int
	result.Mod(result.Mul(a, result.Mul(&numerator, denominator.ModInverse(&denominator, mod))), mod)
	return result
}

const (
	DealNewStack      = "new_stack"
	Cut               = "cut"
	DealWithIncrement = "increment"
)

type Instruction struct {
	Kind string
	Arg  int64
}

func InputToInstructions() []Instruction {
	return lib.InputLinesTo(func(line string) Instruction {
		if strings.HasPrefix(line, "deal into new stack") {
			return Instruction{Kind: DealNewStack}
		} else if strings.HasPrefix(line, "cut") {
			arg := lib.ParseInt(strings.Split(line, " ")[1])
			return Instruction{Kind: Cut, Arg: int64(arg)}
		} else if strings.HasPrefix(line, "deal with increment") {
			arg := lib.ParseInt(strings.Split(line, " ")[3])
			return Instruction{Kind: DealWithIncrement, Arg: int64(arg)}
		} else {
			log.Fatalf("unrecognized line: %s", line)
			return Instruction{}
		}
	})
}
