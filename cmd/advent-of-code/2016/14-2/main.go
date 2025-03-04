package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	salt := in.String()

	var keys []int
	for nonce := 0; len(keys) < 64; nonce++ {
		hash := Hash(salt, nonce)
		if c, found := FindTriple(hash); found {
			needed := fmt.Sprintf("%c%c%c%c%c", c, c, c, c, c)
			for i := 1; i <= 1000; i++ {
				if strings.Contains(Hash(salt, nonce+i), needed) {
					keys = append(keys, nonce)
					break
				}
			}
		}
	}

	fmt.Println(keys[len(keys)-1])
}

var memo = make(map[int]string)

func Hash(prefix string, nonce int) string {
	if hash, found := memo[nonce]; found {
		return hash
	}

	input := fmt.Sprintf("%s%d", prefix, nonce)
	for i := 0; i <= 2016; i++ {
		sum := md5.Sum([]byte(input))
		input = hex.EncodeToString(sum[:])
	}

	memo[nonce] = input
	return input
}

func FindTriple(s string) (byte, bool) {
	for i := range len(s) - 2 {
		if s[i] == s[i+1] && s[i] == s[i+2] {
			return s[i], true
		}
	}

	return 0, false
}
