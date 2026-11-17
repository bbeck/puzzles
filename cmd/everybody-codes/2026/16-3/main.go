package main

import (
	"bytes"
	"fmt"

	"github.com/bbeck/puzzles/lib/in"
)

var REVERSE = map[string]string{"B": "F", "F": "B"}

func main() {
	bs := bytes.Repeat(in.Bytes(), 100)

	// The critical observation is that the bytes have three distinct segments
	// in them, the prefix, the middle, and the suffix.  The prefix is just a run
	// of whatever character is being worked on.  The middle is the segment that
	// ends in the character that is being worked on and is constantly being
	// reversed.  The suffix contains the characters that have been sorted through
	// and are not the prefix character.
	//
	// Instead of reversing the middle segment we'll instead change the direction
	// we traverse it in.  This has the same effect as reversing but avoids all
	// the swapping.

	var start, end int
	var dir string

	// As we process the bytes in the suffix we have to remember them since their
	// order changes with the reversing.
	var suffix []byte
	var suffixIdx int

	// Few helpers to read/pop from both ends of the middle.

	front := func() byte {
		if dir == "F" {
			return bs[start]
		}
		return bs[end]
	}

	popFront := func() {
		if dir == "F" {
			start++
		} else {
			end--
		}
	}

	back := func() byte {
		if dir == "F" {
			return bs[end]
		}
		return bs[start]
	}

	popBack := func() {
		suffixIdx--
		if dir == "F" {
			suffix[suffixIdx] = bs[end]
			end--
		} else {
			suffix[suffixIdx] = bs[start]
			start++
		}
	}

	var count int
	for len(bs) > 0 {
		start, end, dir = 1, len(bs)-1, "F"
		suffix, suffixIdx = make([]byte, len(bs)), len(bs)

		var ch = bs[0]
		for start <= end {
			for start <= end && front() == ch {
				popFront()
			}
			for start <= end && back() != ch {
				popBack()
			}

			if start > end {
				break
			}

			count += end - start + 1
			dir = REVERSE[dir]
		}

		bs = suffix[suffixIdx:]
	}

	fmt.Println(count)
}
