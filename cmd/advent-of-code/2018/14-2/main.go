package main

import (
	"bytes"
	"fmt"
	"github.com/bbeck/advent-of-code/puz"
)

func main() {
	digits := InputToDigits()
	L := len(digits) + 1 // How many entries from the end of the recipes list to check for our digits
	recipes := []byte{3, 7}

	index := -1
	elf1, elf2 := 0, 1
	for index == -1 {
		sum := recipes[elf1] + recipes[elf2]
		recipes = append(recipes, puz.Digits(sum)...)
		elf1 = (elf1 + int(recipes[elf1]) + 1) % len(recipes)
		elf2 = (elf2 + int(recipes[elf2]) + 1) % len(recipes)

		end := puz.Max(0, len(recipes)-L)
		index = bytes.Index(recipes[end:], digits)
	}

	fmt.Println(len(recipes) - L + index)
}

func InputToDigits() []byte {
	var digits []byte
	for _, c := range puz.InputToString() {
		digits = append(digits, byte(puz.ParseInt(string(c))))
	}
	return digits
}
