package lib

import (
	"testing"

	"github.com/alecthomas/assert/v2"
)

func TestPose2D_TurnLeft(t *testing.T) {
	type test struct {
		input Pose2D
		want  Pose2D
	}

	tests := []test{
		{input: Pose2D{Point: Point2D{X: 1, Y: 2}, Heading: Up}, want: Pose2D{Point: Point2D{X: 1, Y: 2}, Heading: Left}},
		{input: Pose2D{Point: Point2D{X: 1, Y: 2}, Heading: Right}, want: Pose2D{Point: Point2D{X: 1, Y: 2}, Heading: Up}},
		{input: Pose2D{Point: Point2D{X: 1, Y: 2}, Heading: Down}, want: Pose2D{Point: Point2D{X: 1, Y: 2}, Heading: Right}},
		{input: Pose2D{Point: Point2D{X: 1, Y: 2}, Heading: Left}, want: Pose2D{Point: Point2D{X: 1, Y: 2}, Heading: Down}},
	}
	for _, test := range tests {
		t.Run(test.input.String(), func(t *testing.T) {
			assert.Equal(t, test.want, test.input.TurnLeft())
		})
	}
}

func TestPose2D_TurnRight(t *testing.T) {
	type test struct {
		input Pose2D
		want  Pose2D
	}

	tests := []test{
		{input: Pose2D{Point: Point2D{X: 1, Y: 2}, Heading: Up}, want: Pose2D{Point: Point2D{X: 1, Y: 2}, Heading: Right}},
		{input: Pose2D{Point: Point2D{X: 1, Y: 2}, Heading: Right}, want: Pose2D{Point: Point2D{X: 1, Y: 2}, Heading: Down}},
		{input: Pose2D{Point: Point2D{X: 1, Y: 2}, Heading: Down}, want: Pose2D{Point: Point2D{X: 1, Y: 2}, Heading: Left}},
		{input: Pose2D{Point: Point2D{X: 1, Y: 2}, Heading: Left}, want: Pose2D{Point: Point2D{X: 1, Y: 2}, Heading: Up}},
	}
	for _, test := range tests {
		t.Run(test.input.String(), func(t *testing.T) {
			assert.Equal(t, test.want, test.input.TurnRight())
		})
	}
}

func TestPose2D_TurnAround(t *testing.T) {
	type test struct {
		input Pose2D
		want  Pose2D
	}

	tests := []test{
		{input: Pose2D{Point: Point2D{X: 1, Y: 2}, Heading: Up}, want: Pose2D{Point: Point2D{X: 1, Y: 2}, Heading: Down}},
		{input: Pose2D{Point: Point2D{X: 1, Y: 2}, Heading: Right}, want: Pose2D{Point: Point2D{X: 1, Y: 2}, Heading: Left}},
		{input: Pose2D{Point: Point2D{X: 1, Y: 2}, Heading: Down}, want: Pose2D{Point: Point2D{X: 1, Y: 2}, Heading: Up}},
		{input: Pose2D{Point: Point2D{X: 1, Y: 2}, Heading: Left}, want: Pose2D{Point: Point2D{X: 1, Y: 2}, Heading: Right}},
	}
	for _, test := range tests {
		t.Run(test.input.String(), func(t *testing.T) {
			assert.Equal(t, test.want, test.input.TurnAround())
		})
	}
}

func TestPose2D_Forward(t *testing.T) {
	type test struct {
		input Pose2D
		want  Pose2D
	}

	tests := []test{
		{input: Pose2D{Point: Point2D{X: 1, Y: 2}, Heading: Up}, want: Pose2D{Point: Point2D{X: 1, Y: 1}, Heading: Up}},
		{input: Pose2D{Point: Point2D{X: 1, Y: 2}, Heading: Right}, want: Pose2D{Point: Point2D{X: 2, Y: 2}, Heading: Right}},
		{input: Pose2D{Point: Point2D{X: 1, Y: 2}, Heading: Down}, want: Pose2D{Point: Point2D{X: 1, Y: 3}, Heading: Down}},
		{input: Pose2D{Point: Point2D{X: 1, Y: 2}, Heading: Left}, want: Pose2D{Point: Point2D{X: 0, Y: 2}, Heading: Left}},
		{input: Pose2D{Point: Point2D{X: 0, Y: 0}, Heading: Up}, want: Pose2D{Point: Point2D{X: 0, Y: -1}, Heading: Up}},
	}
	for _, test := range tests {
		t.Run(test.input.String(), func(t *testing.T) {
			assert.Equal(t, test.want, test.input.Forward())
		})
	}
}
