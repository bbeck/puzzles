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

	var nonce int
	for nonce = 0; ; nonce++ {
		hash := Hash(prefix, nonce)
		if strings.HasPrefix(hash, "000000") {
			break
		}
	}

	fmt.Println(nonce)
}

func Hash(prefix string, nonce int) string {
	input := fmt.Sprintf("%s%d", prefix, nonce)
	hash := md5.Sum([]byte(input))
	return hex.EncodeToString(hash[:])
}
