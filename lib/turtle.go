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
	var move func(d Point2D) Point2D
	switch t.Heading {
	case Up:
		move = func(p Point2D) Point2D { return p.Up() }
	case Right:
		move = func(p Point2D) Point2D { return p.Right() }
	case Down:
		move = func(p Point2D) Point2D { return p.Down() }
	case Left:
		move = func(p Point2D) Point2D { return p.Left() }
	default:
		return
	}

	for ; n > 0; n-- {
		t.Location = move(t.Location)
	}
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
