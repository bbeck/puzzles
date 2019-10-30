package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"strings"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	key := aoc.InputToString(2015, 4)

	var nonce int
	for nonce = 0; ; nonce++ {
		hash := md5.New()
		_, _ = io.WriteString(hash, key)
		_, _ = io.WriteString(hash, fmt.Sprintf("%d", nonce))
		data := hex.EncodeToString(hash.Sum(nil))

		if strings.HasPrefix(data, "00000") {
			break
		}
	}

	fmt.Printf("nonce: %d\n", nonce)
}
