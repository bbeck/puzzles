package aoc

// Grid2D is a struct that organizes data into a two-dimensional grid of cells
// indexed by location.
type Grid2D[T any] struct {
	Cells  []T
	Height int
	Width  int
}

// NewGrid2D instantiates a new two-dimensional grid with the specified
// dimensions.
func NewGrid2D[T any](width, height int) Grid2D[T] {
	return Grid2D[T]{
		Cells:  make([]T, width*height),
		Width:  width,
		Height: height,
	}
}

// Add adds a new value to the grid at the specified location.
func (g *Grid2D[T]) Add(p Point2D, value T) {
	g.AddXY(p.X, p.Y, value)
}

// AddXY adds a new value at the location specified by the X and Y coordinate.
func (g *Grid2D[T]) AddXY(x, y int, value T) {
	g.Cells[y*g.Width+x] = value
}

// Get retrieves the value in the grid at the specified location.  If the
// location doesn't contain a value then the zero value of the underlying grid
// type will be returned.
func (g Grid2D[T]) Get(p Point2D) T {
	return g.GetXY(p.X, p.Y)
}

// GetXY retrieves the value in the grid at the location specified by the X
// and Y coordinate.  If the location doesn't contain a value then the zero
// value of the underlying grid type will be returned.
func (g Grid2D[T]) GetXY(x, y int) T {
	return g.Cells[y*g.Width+x]
}

// InBounds determines if the location specified is in bounds of the grid.
func (g Grid2D[T]) InBounds(p Point2D) bool {
	return g.InBoundsXY(p.X, p.Y)
}

// InBoundsXY determines if the location specified by the X and Y coordinate is
// in bounds of the grid.
func (g Grid2D[T]) InBoundsXY(x, y int) bool {
	return x >= 0 && x < g.Width && y >= 0 && y < g.Height
}

func (g Grid2D[T]) ForEach(fn func(Point2D, T)) {
	g.ForEachXY(func(x, y int, value T) {
		fn(Point2D{X: x, Y: y}, value)
	})
}

func (g Grid2D[T]) ForEachXY(fn func(int, int, T)) {
	for y := 0; y < g.Height; y++ {
		for x := 0; x < g.Width; x++ {
			fn(x, y, g.GetXY(x, y))
		}
	}
}
