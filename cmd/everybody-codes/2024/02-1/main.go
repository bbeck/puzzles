package main

import (
  "fmt"

  "github.com/bbeck/puzzles/lib"
)

func main() {
  for _, line := range lib.InputToLines() {
    fmt.Println(line)
  }
}
