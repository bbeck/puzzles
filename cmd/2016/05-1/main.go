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
	prefix := aoc.InputToString(2016, 5)

	password := ""
	index := 0
	for n := 0; n < 8; n++ {
		var checksum string

		for {
			hash := md5.New()
			_, _ = io.WriteString(hash, prefix)
			_, _ = io.WriteString(hash, fmt.Sprintf("%d", index))
			checksum = hex.EncodeToString(hash.Sum(nil))

			if strings.HasPrefix(checksum, "00000") {
				break
			}

			index++
		}

		password = password + string(checksum[5])
		index++
	}

	fmt.Printf("password: %s\n", password)
}
