package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	prefix := aoc.InputToString(2016, 5)

	var password string
	for nonce := 0; len(password) < 8; nonce++ {
		hash := Hash(prefix, nonce)
		if strings.HasPrefix(hash, "00000") {
			password += string(hash[5])
		}
	}

	fmt.Println(password)
}

func Hash(prefix string, nonce int) string {
	input := fmt.Sprintf("%s%d", prefix, nonce)
	hash := md5.Sum([]byte(input))
	return hex.EncodeToString(hash[:])
}
