package lib

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert/v2"
)

func TestAbs(t *testing.T) {
	type test struct {
		input int
		want  int
	}

	tests := []test{
		{input: 1, want: 1},
		{input: -1, want: 1},
		{input: 0, want: 0},
		{input: 123, want: 123},
		{input: -123, want: 123},
	}
	for _, test := range tests {
		t.Run(fmt.Sprintf("%d", test.input), func(t *testing.T) {
			assert.Equal(t, test.want, Abs(test.input))
		})
	}
}

func TestSign(t *testing.T) {
	type test struct {
		input int
		want  int
	}

	tests := []test{
		{input: 1, want: 1},
		{input: -1, want: -1},
		{input: 0, want: 0},
		{input: 123, want: 1},
		{input: -123, want: -1},
	}
	for _, test := range tests {
		t.Run(fmt.Sprintf("%d", test.input), func(t *testing.T) {
			assert.Equal(t, test.want, Sign(test.input))
		})
	}
}

func TestSum(t *testing.T) {
	type test struct {
		input []int
		want  int
	}

	tests := []test{
		{input: []int{1}, want: 1},
		{input: []int{1, 2}, want: 3},
		{input: []int{}, want: 0},
		{input: []int{-1, 1, 2, 3, 4}, want: 9},
	}
	for _, test := range tests {
		t.Run(fmt.Sprintf("%d", test.input), func(t *testing.T) {
			assert.Equal(t, test.want, Sum(test.input...))
		})
	}
}

func TestProduct(t *testing.T) {
	type test struct {
		input []int
		want  int
	}

	tests := []test{
		{input: []int{1}, want: 1},
		{input: []int{1, 2}, want: 2},
		{input: []int{}, want: 1},
		{input: []int{-1, 1, 2, 3, 4}, want: -24},
	}
	for _, test := range tests {
		t.Run(fmt.Sprintf("%d", test.input), func(t *testing.T) {
			assert.Equal(t, test.want, Product(test.input...))
		})
	}
}

func TestMin(t *testing.T) {
	type test struct {
		input []int
		want  int
	}

	tests := []test{
		{input: []int{1}, want: 1},
		{input: []int{3, 1, 2}, want: 1},
		{input: []int{99}, want: 99},
		{input: []int{1, -1, 2, 3, 4}, want: -1},
	}
	for _, test := range tests {
		t.Run(fmt.Sprintf("%d", test.input), func(t *testing.T) {
			assert.Equal(t, test.want, Min(test.input...))
		})
	}
}

func TestMax(t *testing.T) {
	type test struct {
		input []int
		want  int
	}

	tests := []test{
		{input: []int{1}, want: 1},
		{input: []int{3, 1, 2}, want: 3},
		{input: []int{99}, want: 99},
		{input: []int{1, -1, 2, 4, 3}, want: 4},
	}
	for _, test := range tests {
		t.Run(fmt.Sprintf("%d", test.input), func(t *testing.T) {
			assert.Equal(t, test.want, Max(test.input...))
		})
	}
}

func TestClamp(t *testing.T) {
	type test struct {
		input, min, max int
		want            int
	}

	tests := []test{
		{input: 3, min: 0, max: 5, want: 3},
		{input: 0, min: 0, max: 5, want: 0},
		{input: 5, min: 0, max: 5, want: 5},
		{input: -10, min: 0, max: 5, want: 0},
		{input: 10, min: 0, max: 5, want: 5},
	}
	for _, test := range tests {
		t.Run(fmt.Sprintf("%d", test.input), func(t *testing.T) {
			assert.Equal(t, test.want, Clamp(test.input, test.min, test.max))
		})
	}
}

func TestModulo(t *testing.T) {
	type test struct {
		value, mod int
		want       int
	}

	tests := []test{
		{value: 3, mod: 2, want: 1},
		{value: 4, mod: 2, want: 0},
		{value: 4, mod: 3, want: 1},
		{value: 99, mod: 12, want: 3},
	}
	for _, test := range tests {
		t.Run(fmt.Sprintf("%d %% %d", test.value, test.mod), func(t *testing.T) {
			assert.Equal(t, test.want, Modulo(test.value, test.mod))
		})
	}
}

func TestPow(t *testing.T) {
	type test struct {
		value int
		pow   uint
		want  int
	}

	tests := []test{
		{value: 3, pow: 0, want: 1},
		{value: 3, pow: 2, want: 9},
		{value: 4, pow: 2, want: 16},
		{value: 4, pow: 3, want: 64},
		{value: 5, pow: 12, want: 244140625},
	}
	for _, test := range tests {
		t.Run(fmt.Sprintf("%d ^ %d", test.value, test.pow), func(t *testing.T) {
			assert.Equal(t, test.want, Pow(test.value, test.pow))
		})
	}
}

func TestModPow(t *testing.T) {
	type test struct {
		value, pow, mod int
		want            int
	}

	tests := []test{
		{value: 3, pow: 0, mod: 2, want: 1},
		{value: 3, pow: 200, mod: 2, want: 1},
		{value: 4, pow: 123, mod: 97, want: 64},
		{value: 129, pow: 3, mod: 64, want: 1},
		{value: 5, pow: 12, mod: 8429, want: 3069},
	}
	for _, test := range tests {
		t.Run(fmt.Sprintf("(%d ^ %d) %% %d", test.value, test.pow, test.mod), func(t *testing.T) {
			assert.Equal(t, test.want, ModPow(test.value, test.pow, test.mod))
		})
	}
}

func TestDigits(t *testing.T) {
	type test struct {
		value int
		want  []int
	}

	tests := []test{
		{value: 0, want: []int{0}},
		{value: 1, want: []int{1}},
		{value: 2, want: []int{2}},
		{value: 9, want: []int{9}},
		{value: 15, want: []int{1, 5}},
		{value: 125, want: []int{1, 2, 5}},
		{value: 1539284828, want: []int{1, 5, 3, 9, 2, 8, 4, 8, 2, 8}},
	}
	for _, test := range tests {
		t.Run(fmt.Sprintf("%d", test.value), func(t *testing.T) {
			assert.Equal(t, test.want, Digits(test.value))
		})
	}
}

func TestJoinDigits(t *testing.T) {
	type test struct {
		value []int
		want  int
	}

	tests := []test{
		{value: []int{0}, want: 0},
		{value: []int{1}, want: 1},
		{value: []int{2}, want: 2},
		{value: []int{9}, want: 9},
		{value: []int{1, 5}, want: 15},
		{value: []int{1, 2, 5}, want: 125},
		{value: []int{1, 5, 3, 9, 2, 8, 4, 8, 2, 8}, want: 1539284828},
	}
	for _, test := range tests {
		t.Run(fmt.Sprintf("%d", test.value), func(t *testing.T) {
			assert.Equal(t, test.want, JoinDigits(test.value))
		})
	}
}

func TestChineseRemainderTheorm(t *testing.T) {
	type test struct {
		as, ns []int
		want   int
	}

	tests := []test{
		{as: []int{2, 3, 2}, ns: []int{3, 5, 7}, want: 23},
		{as: []int{3, 5}, ns: []int{5, 7}, want: 33},
		{as: []int{2, 3, 10}, ns: []int{5, 7, 11}, want: 87},
	}
	for _, test := range tests {
		t.Run(fmt.Sprintf("CRT %v (%v)", test.as, test.ns), func(t *testing.T) {
			assert.Equal(t, test.want, ChineseRemainderTheorem(test.as, test.ns))
		})
	}
}
