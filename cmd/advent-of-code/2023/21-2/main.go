package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/lib"
)

func main() {
	grid, start := InputToGridAndStartingLocation()
	S, N := 26501365, grid.Width

	// Observing the input data we see that there's a direct path from the
	// starting location to adjacent grid tiles.  This means that the shortest
	// path from the start point to any grid is going to enter from one of the
	// side midpoints or one of the corners.  This means that in the end the
	// solution will have a rough diamond shape.
	//
	// Additionally, because the width of an individual grid is odd adjacent grids
	// will have an alternating parity of cells that can be reached when filled.
	//
	// The approach to solving this will be as follows.  We will first compute the
	// general dimensions of the diamond in terms of grids.  Then we'll determine
	// how many of these interior, fully filled, grids there are of each parity.
	// Finally, we'll consider the perimeter of the diamond that has partially
	// filled grids.

	// 1. Compute the dimensions of the diamond.
	//
	// We observe that the number of steps we have to take has a remainder of 65
	// after dividing the by the grid size. This means that because we start in
	// the center of a grid with dimension 131, we'll use those 65 steps to get
	// to the very edge of the starting grid and then will need 1 more step to
	// reach the next grid.  From there on we'll be able to count in whole grids.
	GW := (S - N/2 - 1) / N

	// 2. Count how many "odd" and "even" grids are fully contained within the
	// perimeter.
	//
	// Drawing a picture shows that counting the number of odd/even grids can be
	// done by computing a triangular sum parameterized by the number of odd/even
	// grids in the largest row.  For odd grids this will be GW from above, but
	// for even grids this will be (GW+1).
	//
	// odd  = 1+2+...+GW+...+2+1 = GW(GW-1) + GW = GW^2
	// even = 1+2+...+(GW+1)+...+2+1 = GW(GW+1) + (GW+1) = (GW+1)^2
	//
	//        .
	//       .E.
	//      .EOE.
	//     .EOEOE.
	//    .EOEOEOE.
	//     .EOEOE.
	//      .EOE.
	//       .E.
	//        .
	odd, even := GW*GW, (GW+1)*(GW+1)

	// 3. Count how many "small" and "large" grids are on the perimeter.
	//
	// Drawing a picture shows that there are two types of grids on the perimeter.
	// A "small" grid where only a small corner of it will be reached, and a
	// "large" grid where a much larger corner of it will be reached.  Along an
	// edge there (GW+1) fully filled grids.  Each of these is going to touch a
	// "small" and a "large" partially filled grid.  In addition, there is one
	// extra "small" partially filled grid touching the point of the diamond.
	//
	//     -----------------
	//     |   ^   | Small |
	//     | /   \ |       |
	//     |/     \|       |
	//     /       \       |
	//     |       |\      |
	//     |       | \     |
	//     -----------\-----
	//     |       |   \   |
	//     |       |    \  |
	//     |       |     \ |
	//     |       |      \|
	//     |       |       \
	//     |       | Large |
	//     -----------------
	small, large := GW+1, GW

	// 4. Compute the filled area of each type of grid cell.
	var total int

	// 4a. Fully filled grid cells.  For these we'll count the number of reachable
	// cells in an "odd" grid and an "even" grid.  The number of steps we allow
	// ourselves in each cell isn't very important as long as there are enough,
	// and they match the parity of the grid cell.  2N steps is enough to walk the
	// full perimeter of the grid which should also give enough time to fill the
	// interior.
	total += odd * Count(grid, start, 2*N+1)
	total += even * Count(grid, start, 2*N)

	// 4b. Points of the diamond.  These will be reached from the midpoint of the
	// edge closest to the starting grid.  Based on our previous computation we
	// know that when we reach one of these grids we'll have just enough steps
	// remaining to reach the other side.
	total += Count(grid, lib.Point2D{X: start.X, Y: N - 1}, N-1) // Top
	total += Count(grid, lib.Point2D{X: 0, Y: start.Y}, N-1)     // Right
	total += Count(grid, lib.Point2D{X: start.X, Y: 0}, N-1)     // Bottom
	total += Count(grid, lib.Point2D{X: N - 1, Y: start.Y}, N-1) // Left

	// 4c. "Small" partially filled grids.  These will be first reached in the
	// corner that is closest to the adjacent fully filled grid.  When the
	// adjacent grid is reached there will be N-1 steps remaining, but since it is
	// reached at the midpoint of a side, N/2 of those steps will have to be spent
	// to get to the corner of the "small" grid.
	total += small * Count(grid, lib.Point2D{X: N - 1, Y: N - 1}, N/2-1) // Top left
	total += small * Count(grid, lib.Point2D{X: 0, Y: N - 1}, N/2-1)     // Top right
	total += small * Count(grid, lib.Point2D{X: 0, Y: 0}, N/2-1)         // Bottom right
	total += small * Count(grid, lib.Point2D{X: N - 1, Y: 0}, N/2-1)     // Bottom left

	// 4d. "Large" partially filled grids.  These behave similarly to the "small"
	// grids, however they can be first reached by an interior grid.  This means
	// that they will have an additional N steps available to them.
	total += large * Count(grid, lib.Point2D{X: N - 1, Y: N - 1}, N+N/2-1) // Top left
	total += large * Count(grid, lib.Point2D{X: 0, Y: N - 1}, N+N/2-1)     // Top right
	total += large * Count(grid, lib.Point2D{X: 0, Y: 0}, N+N/2-1)         // Bottom right
	total += large * Count(grid, lib.Point2D{X: N - 1, Y: 0}, N+N/2-1)     // Bottom left

	fmt.Println(total)
}

func Count(g lib.Grid2D[string], p lib.Point2D, n int) int {
	// To run more quickly use the parity property to prune the search space.  If
	// a point is visited with an even number of steps remaining then it is part
	// of the final set of points that will be reachable when there are no steps
	// remaining.
	type State struct {
		lib.Point2D
		N int
	}

	var seen lib.Set[lib.Point2D]
	var counted lib.Set[lib.Point2D]

	isGoal := func(s State) bool {
		seen.Add(s.Point2D)
		if s.N%2 == 0 {
			counted.Add(s.Point2D)
		}

		return false
	}

	children := func(s State) []State {
		if s.N == 0 {
			return nil
		}

		var children []State
		g.ForEachOrthogonalNeighbor(s.Point2D, func(q lib.Point2D, ch string) {
			if !seen.Contains(q) && ch != "#" {
				children = append(children, State{q, s.N - 1})
			}
		})

		return children
	}

	lib.BreadthFirstSearch(State{p, n}, children, isGoal)
	return len(counted)
}

func InputToGridAndStartingLocation() (lib.Grid2D[string], lib.Point2D) {
	grid := lib.InputToStringGrid2D()

	var start lib.Point2D
	grid.ForEachPoint(func(p lib.Point2D, s string) {
		if s == "S" {
			start = p
		}
	})

	return grid, start
}
