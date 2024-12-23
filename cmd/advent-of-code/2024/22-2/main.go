package main

import (
	"fmt"
	. "github.com/bbeck/puzzles/lib"
)

func main() {
	var profits []map[Sequence]int
	for _, secret := range InputToInts() {
		profits = append(profits, Profits(secret))
	}

	var best int
	totals := make(map[Sequence]int)
	for _, ps := range profits {
		for sequence, profit := range ps {
			totals[sequence] += profit
			best = Max(best, totals[sequence])
		}
	}
	fmt.Println(best)
}

type Sequence [4]int

func Profits(secret int) map[Sequence]int {
	var prices, changes []int

	prices = append(prices, secret%10)
	for i := range 2000 {
		secret = Next(secret)
		prices = append(prices, secret%10)
		changes = append(changes, secret%10-prices[i])
	}

	var profits = make(map[Sequence]int)
	for i := 4; i < len(changes); i++ {
		sequence := Sequence(changes[i-4 : i])

		if _, found := profits[sequence]; !found {
			profits[sequence] = prices[i]
		}
	}

	return profits
}

func Next(secret int) int {
	secret = (secret ^ (secret * 64)) % 16777216
	secret = (secret ^ (secret / 32)) % 16777216
	secret = (secret ^ (secret * 2048)) % 16777216
	return secret
}
