package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/bbeck/advent-of-code/puz"
	"strings"
)

func main() {
	salt := puz.InputToString()

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
	sum := md5.Sum([]byte(input))
	hash := hex.EncodeToString(sum[:])
	memo[nonce] = hash
	return hash
}

func FindTriple(s string) (byte, bool) {
	for i := 0; i < len(s)-2; i++ {
		if s[i] == s[i+1] && s[i] == s[i+2] {
			return s[i], true
		}
	}

	return 0, false
}
