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

// Add adds a new value at the location specified by the X and Y coordinate.
func (g *Grid2D[T]) Add(x, y int, value T) {
	g.Cells[y*g.Width+x] = value
}

// AddPoint adds a new value to the grid at the specified location.
func (g *Grid2D[T]) AddPoint(p Point2D, value T) {
	g.Add(p.X, p.Y, value)
}

// Get retrieves the value in the grid at the location specified by the X
// and Y coordinate.  If the location doesn't contain a value then the zero
// value of the underlying grid type will be returned.
func (g *Grid2D[T]) Get(x, y int) T {
	return g.Cells[y*g.Width+x]
}

// GetPoint retrieves the value in the grid at the specified location.  If the
// location doesn't contain a value then the zero value of the underlying grid
// type will be returned.
func (g *Grid2D[T]) GetPoint(p Point2D) T {
	return g.Get(p.X, p.Y)
}

// InBounds determines if the location specified by the X and Y coordinate is
// in bounds of the grid.
func (g *Grid2D[T]) InBounds(x, y int) bool {
	return x >= 0 && x < g.Width && y >= 0 && y < g.Height
}

// InBoundsPoint determines if the location specified is in bounds of the grid.
func (g *Grid2D[T]) InBoundsPoint(p Point2D) bool {
	return g.InBounds(p.X, p.Y)
}

// ForEach invokes a callback for every point in the grid.  The x and y
// coordinates of the point along with the value are passed into the callback.
func (g *Grid2D[T]) ForEach(fn func(int, int, T)) {
	for y := 0; y < g.Height; y++ {
		for x := 0; x < g.Width; x++ {
			fn(x, y, g.Get(x, y))
		}
	}
}

// ForEachPoint invokes a callback for every point in the grid.  The point in
// the grid along with the value is passed into the callback.
func (g *Grid2D[T]) ForEachPoint(fn func(Point2D, T)) {
	g.ForEach(func(x, y int, value T) {
		fn(Point2D{X: x, Y: y}, value)
	})
}

// ForEachNeighbor invokes a callback for all 8 neighbors of a point.  The
// neighboring point along with the value are passed into the callback.
func (g *Grid2D[T]) ForEachNeighbor(x0, y0 int, fn func(int, int, T)) {
	for dy := -1; dy <= 1; dy++ {
		y := y0 + dy

		for dx := -1; dx <= 1; dx++ {
			x := x0 + dx

			if (x != x0 || y != y0) && g.InBounds(x, y) {
				fn(x, y, g.Get(x, y))
			}
		}
	}
}

// ForEachNeighborPoint invokes a callback for all 8 neighbors of a point.  The
// neighboring point along with the value are passed into the callback.
func (g *Grid2D[T]) ForEachNeighborPoint(p Point2D, fn func(Point2D, T)) {
	for _, n := range p.Neighbors() {
		if g.InBoundsPoint(n) {
			fn(n, g.GetPoint(n))
		}
	}
}

// ForEachOrthogonalNeighbor invokes a callback for each orthogonal neighbor
// of a point.  The neighboring point along with the value are passed into the
// callback.
func (g *Grid2D[T]) ForEachOrthogonalNeighbor(p Point2D, fn func(Point2D, T)) {
	for _, n := range p.OrthogonalNeighbors() {
		if g.InBoundsPoint(n) {
			fn(n, g.GetPoint(n))
		}
	}
}
