package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"strings"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	salt := aoc.InputToString(2016, 14)
	q := NewQueue(salt, 1000)

	var keys []string
	for i := 0; ; i++ {
		hash, offset := q.Pop()
		if c, found := triple(hash); found {
			needle := fmt.Sprintf("%c%c%c%c%c", c, c, c, c, c)
			for j := 1; j <= 1000; j++ {
				if strings.Contains(q.Peek(j), needle) {
					keys = append(keys, q.Peek(j))
					break
				}
			}
		}

		if len(keys) == 64 {
			fmt.Printf("key %d found on offset: %d\n", len(keys), offset)
			break
		}
	}
}

type Queue struct {
	salt string

	// window contains the current hash we're working on at the 0th index and the
	// next 1000 hashes at indices 1 through 1001.
	window []string

	// offset is the index of the offset of the next hash that will be popped.
	offset int
}

func NewQueue(salt string, size int) *Queue {
	var window []string
	for i := 0; i <= size; i++ {
		window = append(window, hash(salt, i))
	}

	return &Queue{
		salt:   salt,
		window: window,
	}
}

func (q *Queue) Pop() (string, int) {
	q.offset++
	q.window = append(q.window[1:], hash(q.salt, q.offset+1000))

	return q.window[0], q.offset
}

func (q *Queue) Peek(n int) string {
	if n > 1000 {
		log.Fatalf("peeking too far ahead: %d", n)
	}

	return q.window[n]
}

func hash(salt string, offset int) string {
	s := fmt.Sprintf("%s%d", salt, offset)
	for round := 0; round <= 2016; round++ {
		hash := md5.New()
		_, _ = io.WriteString(hash, s)
		s = hex.EncodeToString(hash.Sum(nil))
	}

	return s
}

func triple(s string) (byte, bool) {
	for i := 0; i < len(s)-2; i++ {
		if s[i] == s[i+1] && s[i] == s[i+2] {
			return s[i], true
		}
	}

	return 0, false
}
