package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	prefix := in.String()

	password := make([]uint8, 8)
	for nonce, count := 0, 0; count < 8; nonce++ {
		hash := Hash(prefix, nonce)
		if strings.HasPrefix(hash, "00000") {
			position := hash[5] - '0'
			if position < 8 && password[position] == 0 {
				password[position] = hash[6]
				count++
			}
		}
	}

	fmt.Println(string(password))
}

func Hash(prefix string, nonce int) string {
	input := fmt.Sprintf("%s%d", prefix, nonce)
	hash := md5.Sum([]byte(input))
	return hex.EncodeToString(hash[:])
}
