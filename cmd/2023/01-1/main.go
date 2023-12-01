package main

import (
  "fmt"

  "github.com/bbeck/advent-of-code/aoc"
)

func main() {
  digits := map[rune]int{
    '0': 0,
    '1': 1,
    '2': 2,
    '3': 3,
    '4': 4,
    '5': 5,
    '6': 6,
    '7': 7,
    '8': 8,
    '9': 9,
  }

  nums := aoc.InputLinesTo[int](2023, 1, func(line string) (int, error) {
    var nums []int
    for _, c := range line {
      if n, ok := digits[c]; ok {
        nums = append(nums, n)
      }
    }

    return nums[0]*10 + nums[len(nums)-1], nil
  })

  fmt.Println(aoc.Sum(nums...))
}
