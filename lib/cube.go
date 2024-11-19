package lib

// Cube represents a 3D volume in space with its faces aligned with the axes.
type Cube struct {
	MinX, MaxX int
	MinY, MaxY int
	MinZ, MaxZ int
}

// Volume returns the volume of the cube.
func (c Cube) Volume() int {
	return (c.MaxX - c.MinX) * (c.MaxY - c.MinY) * (c.MaxZ - c.MinZ)
}

// Intersects returns true if two cubes intersect with one another.
func (c Cube) Intersects(d Cube) bool {
	// NOTE: This considers cubes that share a corner as overlapping.
	x := c.MinX <= d.MaxX && c.MaxX >= d.MinX
	y := c.MinY <= d.MaxY && c.MaxY >= d.MinY
	z := c.MinZ <= d.MaxZ && c.MaxZ >= d.MinZ
	return x && y && z
}

// Subtract removes the volume of the provided cube.  This will result in up to
// 6 new cubes being created.
func (c Cube) Subtract(d Cube) []Cube {
	if !c.Intersects(d) {
		return []Cube{c}
	}

	// Clip d to the bounds of c
	d = Cube{
		MinX: Max(c.MinX, d.MinX), MaxX: Min(c.MaxX, d.MaxX),
		MinY: Max(c.MinY, d.MinY), MaxY: Min(c.MaxY, d.MaxY),
		MinZ: Max(c.MinZ, d.MinZ), MaxZ: Min(c.MaxZ, d.MaxZ),
	}

	candidates := []Cube{
		{MinX: c.MinX, MinY: c.MinY, MinZ: c.MinZ, MaxX: c.MaxX, MaxY: c.MaxY, MaxZ: d.MinZ},
		{MinX: c.MinX, MinY: c.MinY, MinZ: d.MaxZ, MaxX: c.MaxX, MaxY: c.MaxY, MaxZ: c.MaxZ},
		{MinX: d.MinX, MinY: c.MinY, MinZ: d.MinZ, MaxX: c.MaxX, MaxY: d.MinY, MaxZ: d.MaxZ},
		{MinX: d.MaxX, MinY: d.MinY, MinZ: d.MinZ, MaxX: c.MaxX, MaxY: c.MaxY, MaxZ: d.MaxZ},
		{MinX: c.MinX, MinY: d.MaxY, MinZ: d.MinZ, MaxX: d.MaxX, MaxY: c.MaxY, MaxZ: d.MaxZ},
		{MinX: c.MinX, MinY: c.MinY, MinZ: d.MinZ, MaxX: d.MinX, MaxY: d.MaxY, MaxZ: d.MaxZ},
	}

	var children []Cube
	for _, c := range candidates {
		if c.Volume() > 0 {
			children = append(children, c)
		}
	}

	return children
}
