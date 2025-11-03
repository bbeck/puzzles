package main

import (
	"fmt"

	. "github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	grid := in.ToGrid2D(func(_, _ int, s string) string { return s })

	var start Point2D
	grid.ForEachPoint(func(p Point2D, s string) {
		if s == "S" {
			start = p
		}
	})

	var dp = Make4D[int](101, grid.Width, grid.Height, 4)
	dp[0][start.X][start.Y][Up] = 1000
	dp[0][start.X][start.Y][Right] = 1000
	dp[0][start.X][start.Y][Down] = 1000
	dp[0][start.X][start.Y][Left] = 1000

	for step := 0; step < 100; step++ {
		for x := 0; x < grid.Width; x++ {
			for y := 0; y < grid.Height; y++ {
				var nx, ny int

				if grid.Get(x, y) == "#" {
					// This is a wall, can't be here.
					continue
				}

				for dir := 0; dir < 4; dir++ {
					current := dp[step][x][y][dir]

					// forward
					nx, ny = x+dx[F[dir]], y+dy[F[dir]]
					if grid.InBounds(nx, ny) {
						dp[step+1][nx][ny][F[dir]] = max(
							dp[step+1][nx][ny][F[dir]],
							current+dz[grid.Get(nx, ny)],
						)
					}

					// left
					nx, ny = x+dx[L[dir]], y+dy[L[dir]]
					if grid.InBounds(nx, ny) {
						dp[step+1][nx][ny][L[dir]] = max(
							dp[step+1][nx][ny][L[dir]],
							current+dz[grid.Get(nx, ny)],
						)
					}

					// right
					nx, ny = x+dx[R[dir]], y+dy[R[dir]]
					if grid.InBounds(nx, ny) {
						dp[step+1][nx][ny][R[dir]] = max(
							dp[step+1][nx][ny][R[dir]],
							current+dz[grid.Get(nx, ny)],
						)
					}
				}
			}
		}
	}

	var best int
	for x := 0; x < grid.Width; x++ {
		for y := 0; y < grid.Height; y++ {
			for dir := 0; dir < 4; dir++ {
				best = max(best, dp[100][x][y][dir])
			}
		}
	}
	fmt.Println(best)
}

var dx = []int{-1, 0, 1, 0}
var dy = []int{0, 1, 0, -1}
var F = []int{0, 1, 2, 3}
var L = []int{3, 0, 1, 2}
var R = []int{1, 2, 3, 0}
var dz = map[string]int{".": -1, "+": 1, "-": -2}

func Make4D[T any](a, b, c, d int) [][][][]T {
	arr := make([][][][]T, a)
	for i := range a {
		arr[i] = Make3D[T](b, c, d)
	}
	return arr
}
