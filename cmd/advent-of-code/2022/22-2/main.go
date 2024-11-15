package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/puz"
	"strings"
)

func main() {
	board, directions := InputToBoard()
	var t puz.Turtle

	// Helper to take one step forward
	step := func(t puz.Turtle) puz.Turtle {
		before := t

		t.Forward(1)

		if !board.InBoundsPoint(t.Location) || board.GetPoint(t.Location) == 0 {
			// We're out of bounds, wrap around the faces of the cube.
			t = MoveToNextFace(before)
		}

		if board.GetPoint(t.Location) == '#' {
			// We hit a wall, go back to our previous location.
			return before
		}

		return t
	}

	// Position the turtle at the starting location.
	t.Heading = puz.Right
	for {
		if board.GetPoint(t.Location) == '.' {
			break
		}
		t.Forward(1)
	}

	for _, dir := range directions {
		switch dir {
		case "L":
			t.TurnLeft()

		case "R":
			t.TurnRight()

		default:
			for i := 0; i < puz.ParseInt(dir); i++ {
				t = step(t)
			}
		}
	}

	x, y := t.Location.X+1, t.Location.Y+1
	var dir int
	switch t.Heading {
	case puz.Right:
		dir = 0
	case puz.Down:
		dir = 1
	case puz.Left:
		dir = 2
	case puz.Up:
		dir = 3
	}
	fmt.Println(1000*y + 4*x + dir)
}

func MoveToNextFace(t puz.Turtle) puz.Turtle {
	// This method assumes a cube with side length of 50 that is unfolded like:
	//    12
	//    3
	//   45
	//   6
	var face int
	switch {
	case 50 <= t.Location.X && t.Location.X < 100 && 0 <= t.Location.Y && t.Location.Y < 50:
		face = 1
	case 100 <= t.Location.X && t.Location.X < 150 && 0 <= t.Location.Y && t.Location.Y < 50:
		face = 2
	case 50 <= t.Location.X && t.Location.X < 100 && 50 <= t.Location.Y && t.Location.Y < 100:
		face = 3
	case 0 <= t.Location.X && t.Location.X < 50 && 100 <= t.Location.Y && t.Location.Y < 150:
		face = 4
	case 50 <= t.Location.X && t.Location.X < 100 && 100 <= t.Location.Y && t.Location.Y < 150:
		face = 5
	case 0 <= t.Location.X && t.Location.X < 50 && 150 <= t.Location.Y && t.Location.Y < 200:
		face = 6
	}

	switch {
	case face == 1 && t.Heading == puz.Up:
		// left of 6, facing right  (50, 0) -> (0, 150) and (99, 0) -> (0, 199)
		t.Location = puz.Point2D{X: 0, Y: t.Location.X + 100}
		t.Heading = puz.Right

	case face == 1 && t.Heading == puz.Left:
		// left of 4, facing right  (50, 0) -> (0, 149) and (50, 49) -> (0, 100)
		t.Location = puz.Point2D{X: 0, Y: 149 - t.Location.Y}
		t.Heading = puz.Right

	case face == 2 && t.Heading == puz.Up:
		// bottom of 6, facing up  (100, 0) -> (0, 199) and (149, 0) -> (49, 199)
		t.Location = puz.Point2D{X: t.Location.X - 100, Y: 199}
		t.Heading = puz.Up

	case face == 2 && t.Heading == puz.Right:
		// right of 5, facing left (149, 0) -> (99, 149) and (149, 49) -> (99, 100)
		t.Location = puz.Point2D{X: 99, Y: 149 - t.Location.Y}
		t.Heading = puz.Left

	case face == 2 && t.Heading == puz.Down:
		// right of 3, facing left (100, 49) -> (99, 50) and (149, 49) -> (99, 99)
		t.Location = puz.Point2D{X: 99, Y: t.Location.X - 50}
		t.Heading = puz.Left

	case face == 3 && t.Heading == puz.Left:
		// top of 4, facing down (50, 50) -> (0, 100) and (50, 99) -> (49, 100)
		t.Location = puz.Point2D{X: t.Location.Y - 50, Y: 100}
		t.Heading = puz.Down

	case face == 3 && t.Heading == puz.Right:
		// bottom of 2, facing up (99, 50) -> (100, 49) and (99, 99) -> (149, 49)
		t.Location = puz.Point2D{X: t.Location.Y + 50, Y: 49}
		t.Heading = puz.Up

	case face == 4 && t.Heading == puz.Up:
		// left of 3, facing right (0, 100) -> (50, 50) and (49, 100) -> (50, 99)
		t.Location = puz.Point2D{X: 50, Y: t.Location.X + 50}
		t.Heading = puz.Right

	case face == 4 && t.Heading == puz.Left:
		// left of 1, facing right (0, 100) -> (50, 49) and (0, 149) -> (50, 0)
		t.Location = puz.Point2D{X: 50, Y: 149 - t.Location.Y}
		t.Heading = puz.Right

	case face == 5 && t.Heading == puz.Right:
		// right of 2, facing left (99, 100) -> (149, 49) and (99, 149) -> (149, 0)
		t.Location = puz.Point2D{X: 149, Y: 149 - t.Location.Y}
		t.Heading = puz.Left

	case face == 5 && t.Heading == puz.Down:
		// right of 6, facing left (50, 149) -> (49, 150) and (99, 149) -> (49, 199)
		t.Location = puz.Point2D{X: 49, Y: t.Location.X + 100}
		t.Heading = puz.Left

	case face == 6 && t.Heading == puz.Left:
		// top of 1, facing down (0, 150) -> (50, 0) and (0, 199) -> (99, 0)
		t.Location = puz.Point2D{X: t.Location.Y - 100, Y: 0}
		t.Heading = puz.Down

	case face == 6 && t.Heading == puz.Right:
		// bottom of 5, facing up (49, 150) -> (50, 149) and (49, 199) -> (99, 149)
		t.Location = puz.Point2D{X: t.Location.Y - 100, Y: 149}
		t.Heading = puz.Up

	case face == 6 && t.Heading == puz.Down:
		// top of 2, facing down (0, 199) -> (100, 0) and (49, 199) -> (149, 0)
		t.Location = puz.Point2D{X: t.Location.X + 100, Y: 0}
		t.Heading = puz.Down
	}

	return t
}

func InputToBoard() (puz.Grid2D[rune], []string) {
	lines := puz.InputToLines()

	var W, H int
	for y := 0; y < len(lines)-2; y++ {
		W = puz.Max(W, len(lines[y]))
		H = y + 1
	}

	board := puz.NewGrid2D[rune](W, H)
	for y := 0; y < len(lines)-2; y++ {
		for x, c := range lines[y] {
			if c != ' ' {
				board.Set(x, y, c)
			}
		}
	}

	directions := lines[len(lines)-1]
	directions = strings.ReplaceAll(directions, "L", " L ")
	directions = strings.ReplaceAll(directions, "R", " R ")

	return board, strings.Fields(directions)
}
