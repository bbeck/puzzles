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

	password := make(map[int]string)
	index := 0
	for n := 0; len(password) < 8; n++ {
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

		digit := checksum[5]
		if digit >= '0' && digit <= '7' {
			position := aoc.ParseInt(string(digit))
			if password[position] == "" {
				password[position] = string(checksum[6])
			}
		}

		index++
	}

	fmt.Print("password: ")
	for i := 0; i < 8; i++ {
		fmt.Print(password[i])
	}
	fmt.Println()
}
