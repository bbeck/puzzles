package lib

type Heading int

const (
	Up Heading = iota
	Right
	Down
	Left
)

func (h Heading) String() string {
	switch h {
	case Up:
		return "U"
	case Right:
		return "R"
	case Down:
		return "D"
	case Left:
		return "L"
	}
	return "?"
}

type Turtle struct {
	Location Point2D
	Heading  Heading
}

func (t *Turtle) Forward(n int) {
	var dx, dy int
	switch t.Heading {
	case Up:
		dy = -1
	case Right:
		dx = 1
	case Down:
		dy = 1
	case Left:
		dx = -1
	}

	t.Location.X += n * dx
	t.Location.Y += n * dy
}

func (t *Turtle) TurnLeft() {
	switch t.Heading {
	case Up:
		t.Heading = Left
	case Right:
		t.Heading = Up
	case Down:
		t.Heading = Right
	case Left:
		t.Heading = Down
	}
}

func (t *Turtle) TurnRight() {
	switch t.Heading {
	case Up:
		t.Heading = Right
	case Right:
		t.Heading = Down
	case Down:
		t.Heading = Left
	case Left:
		t.Heading = Up
	}
}
