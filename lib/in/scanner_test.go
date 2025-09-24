package in

import (
	"strings"
	"testing"

	"github.com/alecthomas/assert/v2"
	. "github.com/bbeck/puzzles/lib"
)

func TestScannerByte(t *testing.T) {
	type test struct {
		input     []byte
		want      byte
		remaining []byte
	}

	tests := []test{
		{
			input:     []byte("a"),
			want:      'a',
			remaining: []byte{},
		},
		{
			input:     []byte(" "),
			want:      ' ',
			remaining: []byte{},
		},
		{
			input:     []byte("abc"),
			want:      'a',
			remaining: []byte("bc"),
		},
	}
	for _, test := range tests {
		t.Run(string(test.input), func(t *testing.T) {
			scanner := Scanner[any](test.input)
			assert.Equal(t, test.want, scanner.Byte())
			assert.Equal(t, test.remaining, scanner)
		})
	}
}

func TestScannerByteError(t *testing.T) {
	type test struct {
		input []byte
	}

	tests := []test{
		{input: []byte("")},
		{input: nil},
	}
	for _, test := range tests {
		t.Run(string(test.input), func(t *testing.T) {
			scanner := Scanner[any](test.input)
			assert.Panics(t, func() { scanner.Byte() })
		})
	}
}

func TestScannerBytes(t *testing.T) {
	type test struct {
		input []byte
		want  []byte
	}

	tests := []test{
		{
			input: []byte("a"),
			want:  []byte("a"),
		},
		{
			input: []byte(" "),
			want:  []byte(" "),
		},
		{
			input: []byte("abc"),
			want:  []byte("abc"),
		},
	}
	for _, test := range tests {
		t.Run(string(test.input), func(t *testing.T) {
			scanner := Scanner[any](test.input)
			assert.Equal(t, test.want, scanner.Bytes())
		})
	}
}

func TestScannerBytesError(t *testing.T) {
	type test struct {
		input []byte
	}

	tests := []test{
		{input: []byte("")},
		{input: nil},
	}
	for _, test := range tests {
		t.Run(string(test.input), func(t *testing.T) {
			scanner := Scanner[any](test.input)
			assert.Panics(t, func() { scanner.Bytes() })
		})
	}
}

func TestScannerChunk(t *testing.T) {
	type test struct {
		input     []byte
		want      []string
		remaining []byte
	}

	tests := []test{
		{
			input:     []byte("a"),
			want:      []string{"a"},
			remaining: []byte{},
		},
		{
			input:     []byte("a\nb"),
			want:      []string{"a", "b"},
			remaining: []byte{},
		},
		{
			input:     []byte("a\nb\n"),
			want:      []string{"a", "b"},
			remaining: []byte{},
		},
		{
			input:     []byte("a\nb\nc"),
			want:      []string{"a", "b", "c"},
			remaining: []byte{},
		},
		{
			input:     []byte("a\nb\n\nc\nd"),
			want:      []string{"a", "b"},
			remaining: []byte("c\nd"),
		},
		{
			input:     []byte("a\nb\n\n\n\nc\nd"),
			want:      []string{"a", "b"},
			remaining: []byte("c\nd"),
		},
		{
			input:     []byte("a\n"),
			want:      []string{"a"},
			remaining: []byte{},
		},
	}
	for _, test := range tests {
		t.Run("Chunk", func(t *testing.T) {
			t.Run(string(test.input), func(t *testing.T) {
				scanner := Scanner[any](test.input)
				assert.Equal(t, test.want, scanner.Chunk())
				assert.Equal(t, test.remaining, scanner)
			})
		})

		t.Run("ChunkS", func(t *testing.T) {
			t.Run(string(test.input), func(t *testing.T) {
				want := Scanner[any](strings.Join(test.want, "\n"))

				scanner := Scanner[any](test.input)
				assert.Equal(t, want, scanner.ChunkS())
				assert.Equal(t, test.remaining, scanner)
			})
		})
	}
}

func TestScannerChunkError(t *testing.T) {
	type test struct {
		input []byte
	}

	tests := []test{
		{input: nil},
		{input: []byte("")},
	}
	for _, test := range tests {
		t.Run("Chunk", func(t *testing.T) {
			t.Run(string(test.input), func(t *testing.T) {
				scanner := Scanner[any](test.input)
				assert.Panics(t, func() { scanner.Chunk() })
			})
		})

		t.Run("ChunkS", func(t *testing.T) {
			t.Run(string(test.input), func(t *testing.T) {
				scanner := Scanner[any](test.input)
				assert.Panics(t, func() { scanner.ChunkS() })
			})
		})
	}
}

func TestScannerCut(t *testing.T) {
	type test struct {
		input     []byte
		separator string
		want      []string
		remaining []byte
	}

	tests := []test{
		{
			input:     []byte("a=b"),
			separator: "=",
			want:      []string{"a", "b"},
			remaining: []byte{},
		},
		{
			input:     []byte("a,b"),
			separator: ",",
			want:      []string{"a", "b"},
			remaining: []byte{},
		},
		{
			input:     []byte("a -> b"),
			separator: " -> ",
			want:      []string{"a", "b"},
			remaining: []byte{},
		},
		{
			input:     []byte("a -> b -> c"),
			separator: " -> ",
			want:      []string{"a", "b -> c"},
			remaining: []byte{},
		},
		{
			input:     []byte("a -> b\nc -> d"),
			separator: " -> ",
			want:      []string{"a", "b"},
			remaining: []byte("c -> d"),
		},
		{
			input:     []byte("abc"),
			separator: ",",
			want:      []string{"abc", ""},
			remaining: []byte{},
		},
	}
	for _, test := range tests {
		t.Run("Cut", func(t *testing.T) {
			t.Run(string(test.input), func(t *testing.T) {
				scanner := Scanner[any](test.input)
				lhs, rhs := scanner.Cut(test.separator)
				assert.Equal(t, test.want[0], lhs)
				assert.Equal(t, test.want[1], rhs)
				assert.Equal(t, test.remaining, scanner)
			})
		})

		t.Run("CutS", func(t *testing.T) {
			t.Run(string(test.input), func(t *testing.T) {
				scanner := Scanner[any](test.input)
				lhs, rhs := scanner.CutS(test.separator)
				assert.Equal(t, []byte(test.want[0]), lhs)
				assert.Equal(t, []byte(test.want[1]), rhs)
				assert.Equal(t, test.remaining, scanner)
			})
		})
	}
}

func TestScannerCutError(t *testing.T) {
	type test struct {
		input     []byte
		separator string
	}

	tests := []test{
		{
			input: nil,
		},
		{
			input:     nil,
			separator: ",",
		},
		{
			input: []byte(""),
		},
		{
			input:     []byte(""),
			separator: ",",
		},
	}
	for _, test := range tests {
		t.Run("Cut", func(t *testing.T) {
			t.Run(string(test.input), func(t *testing.T) {
				scanner := Scanner[any](test.input)
				assert.Panics(t, func() { scanner.Cut(test.separator) })
			})
		})

		t.Run("CutS", func(t *testing.T) {
			t.Run(string(test.input), func(t *testing.T) {
				scanner := Scanner[any](test.input)
				assert.Panics(t, func() { scanner.CutS(test.separator) })
			})
		})
	}
}

func TestScannerExpect(t *testing.T) {
	type test struct {
		input     []byte
		expect    string
		remaining []byte
	}

	tests := []test{
		{
			input:  []byte("a"),
			expect: "a",
		},
		{
			input:     []byte("abc"),
			expect:    "a",
			remaining: []byte("bc"),
		},
		{
			input:     []byte("abc"),
			expect:    "",
			remaining: []byte("abc"),
		},
	}
	for _, test := range tests {
		t.Run(string(test.input), func(t *testing.T) {
			scanner := Scanner[any](test.input)
			scanner.Expect(test.expect)
			assert.Equal(t, test.remaining, []byte(scanner))
		})
	}
}

func TestScannerExpectError(t *testing.T) {
	type test struct {
		input  []byte
		expect string
	}

	tests := []test{
		{
			input:  []byte("a"),
			expect: "b",
		},
		{
			input:  []byte("a"),
			expect: "abc",
		},
		{
			input:  []byte(""),
			expect: "a",
		},
		{
			input:  nil,
			expect: "a",
		},
	}
	for _, test := range tests {
		t.Run(string(test.input), func(t *testing.T) {
			scanner := Scanner[any](test.input)
			assert.Panics(t, func() { scanner.Expect(test.expect) })
		})
	}
}

func TestScannerFields(t *testing.T) {
	type test struct {
		input     []byte
		want      []string
		remaining []byte
	}

	tests := []test{
		{
			input:     []byte("a"),
			want:      []string{"a"},
			remaining: []byte{},
		},
		{
			input:     []byte("a b"),
			want:      []string{"a", "b"},
			remaining: []byte{},
		},
		{
			input:     []byte("a b c"),
			want:      []string{"a", "b", "c"},
			remaining: []byte{},
		},
		{
			input:     []byte("abc def"),
			want:      []string{"abc", "def"},
			remaining: []byte{},
		},
		{
			input:     []byte("abc\tdef"),
			want:      []string{"abc", "def"},
			remaining: []byte{},
		},
		{
			input:     []byte("a b\nc d"),
			want:      []string{"a", "b"},
			remaining: []byte("c d"),
		},
		{
			input:     []byte(" "),
			want:      []string{},
			remaining: []byte{},
		},
	}
	for _, test := range tests {
		t.Run("Fields", func(t *testing.T) {
			t.Run(string(test.input), func(t *testing.T) {
				scanner := Scanner[any](test.input)
				fields := scanner.Fields()
				assert.Equal(t, test.want, fields)
				assert.Equal(t, test.remaining, scanner)
			})
		})

		t.Run("FieldsS", func(t *testing.T) {
			t.Run(string(test.input), func(t *testing.T) {
				var want []Scanner[any]
				for _, w := range test.want {
					want = append(want, Scanner[any](w))
				}

				scanner := Scanner[any](test.input)
				fields := scanner.FieldsS()
				assert.Equal(t, want, fields)
				assert.Equal(t, test.remaining, scanner)
			})
		})
	}
}

func TestScannerFieldsError(t *testing.T) {
	type test struct {
		input []byte
	}

	tests := []test{
		{input: nil},
		{input: []byte("")},
	}
	for _, test := range tests {
		t.Run("Fields", func(t *testing.T) {
			t.Run(string(test.input), func(t *testing.T) {
				scanner := Scanner[any](test.input)
				assert.Panics(t, func() { scanner.Fields() })
			})
		})

		t.Run("FieldsS", func(t *testing.T) {
			t.Run(string(test.input), func(t *testing.T) {
				scanner := Scanner[any](test.input)
				assert.Panics(t, func() { scanner.FieldsS() })
			})
		})
	}
}

func TestScannerGrid2D(t *testing.T) {
	type test struct {
		input []byte
		fn    func(int, int, string) any
		want  Grid2D[any]
	}

	tests := []test{
		{
			input: []byte("ab\ncd"),
			fn:    func(x, y int, s string) any { return s },
			want:  Grid2D[any]{Cells: []any{"a", "b", "c", "d"}, Width: 2, Height: 2},
		},
		{
			input: []byte("123\n456"),
			fn:    func(x, y int, s string) any { return ParseInt(s) },
			want:  Grid2D[any]{Cells: []any{1, 2, 3, 4, 5, 6}, Width: 3, Height: 2},
		},
	}
	for _, test := range tests {
		t.Run(string(test.input), func(t *testing.T) {
			scanner := Scanner[any](test.input)
			assert.Equal(t, test.want, scanner.Grid2D(test.fn))
		})
	}
}

func TestScannerHasNext(t *testing.T) {
	type test struct {
		input []byte
		want  bool
	}

	tests := []test{
		{input: []byte("a"), want: true},
		{input: []byte("abc"), want: true},
		{input: []byte(""), want: false},
		{input: []byte(" "), want: false},
		{input: []byte(" \r\n"), want: false},
		{input: []byte(" \r\na"), want: true},
		{input: nil, want: false},
	}
	for _, test := range tests {
		t.Run(string(test.input), func(t *testing.T) {
			scanner := Scanner[any](test.input)
			assert.Equal(t, test.want, scanner.HasNext())
		})
	}
}

func TestScannerHasNextLine(t *testing.T) {
	type test struct {
		input []byte
		want  bool
	}

	tests := []test{
		{input: nil, want: false},
		{input: []byte(""), want: false},
		{input: []byte("a"), want: false},
		{input: []byte("abc"), want: false},
		{input: []byte(""), want: false},
		{input: []byte(" "), want: false},
		{input: []byte(" \r\n"), want: false},
		{input: []byte(" \r\na"), want: true},
		{input: nil, want: false},
	}
	for _, test := range tests {
		t.Run(string(test.input), func(t *testing.T) {
			scanner := Scanner[any](test.input)
			assert.Equal(t, test.want, scanner.HasNextLine())
		})
	}
}

func TestScannerHasPrefix(t *testing.T) {
	type test struct {
		input  []byte
		prefix string
		want   bool
	}

	tests := []test{
		{input: []byte("abc"), prefix: "a", want: true},
		{input: []byte("abc"), prefix: "ab", want: true},
		{input: []byte("abc"), prefix: "abc", want: true},
		{input: []byte("abc"), prefix: "abcd", want: false},
		{input: []byte("abc"), prefix: "b", want: false},
		{input: []byte("abc"), prefix: "c", want: false},
		{input: []byte("abc"), prefix: "", want: true},
		{input: []byte(""), prefix: "", want: true},
		{input: []byte(""), prefix: "a", want: false},
	}
	for _, test := range tests {
		t.Run(string(test.input), func(t *testing.T) {
			scanner := Scanner[any](test.input)
			assert.Equal(t, test.want, scanner.HasPrefix(test.prefix))
		})
	}
}

func TestScannerInt(t *testing.T) {
	type test struct {
		input     []byte
		want      int
		remaining []byte
	}

	tests := []test{
		{input: []byte("0"), want: 0},
		{input: []byte("1"), want: 1},
		{input: []byte("123"), want: 123},
		{input: []byte("-1"), want: -1},
		{input: []byte("-123"), want: -123},
		{input: []byte(" 123"), want: 123},
		{input: []byte("abc123"), want: 123},
		{input: []byte("123a"), want: 123, remaining: []byte("a")},
		{input: []byte("-123a"), want: -123, remaining: []byte("a")},
		{input: []byte("-123-456"), want: -123, remaining: []byte("-456")},
		{input: []byte("--123"), want: -123, remaining: []byte("")},
	}
	for _, test := range tests {
		t.Run(string(test.input), func(t *testing.T) {
			scanner := Scanner[any](test.input)
			assert.Equal(t, test.want, scanner.Int())
			assert.Equal(t, test.remaining, []byte(scanner))
		})
	}
}

func TestScannerIntError(t *testing.T) {
	type test struct {
		input []byte
	}

	tests := []test{
		{input: []byte("-")},
		{input: []byte("-a")},
		{input: []byte("abc")},
		{input: []byte("")},
		{input: nil},
	}
	for _, test := range tests {
		t.Run(string(test.input), func(t *testing.T) {
			scanner := Scanner[any](test.input)
			assert.Panics(t, func() { scanner.Int() })
		})
	}
}

func TestScannerInts(t *testing.T) {
	type test struct {
		input     []byte
		want      []int
		remaining []byte
	}

	tests := []test{
		{input: []byte("")},
		{input: []byte("123"), want: []int{123}, remaining: []byte{}},
		{input: []byte("123a"), want: []int{123}, remaining: []byte("a")},
		{input: []byte("-123a"), want: []int{-123}, remaining: []byte("a")},
		{input: []byte("-123-456"), want: []int{-123, -456}},
		{input: []byte("-123,456"), want: []int{-123, 456}},
		{input: []byte("-123 456"), want: []int{-123, 456}},
		{input: []byte("--123"), want: []int{-123}},
	}
	for _, test := range tests {
		t.Run(string(test.input), func(t *testing.T) {
			scanner := Scanner[any](test.input)
			assert.Equal(t, test.want, scanner.Ints())
		})
	}
}

func TestScannerLine(t *testing.T) {
	type test struct {
		input     []byte
		want      string
		remaining []byte
	}

	tests := []test{
		{input: []byte("a"), want: "a"},
		{input: []byte("\na"), want: "", remaining: []byte("a")},
		{input: []byte("a "), want: "a "},
		{input: []byte("a b c"), want: "a b c"},
		{input: []byte("a\nb\nc"), want: "a", remaining: []byte("b\nc")},
		{input: []byte("\n"), want: ""},
	}
	for _, test := range tests {
		t.Run(string(test.input), func(t *testing.T) {
			scanner := Scanner[any](test.input)
			assert.Equal(t, test.want, scanner.Line())
			assert.Equal(t, test.remaining, []byte(scanner))
		})
	}
}

func TestScannerLineError(t *testing.T) {
	type test struct {
		input []byte
	}

	tests := []test{
		{input: []byte("")},
		{input: nil},
	}
	for _, test := range tests {
		t.Run(string(test.input), func(t *testing.T) {
			scanner := Scanner[any](test.input)
			assert.Panics(t, func() { scanner.Line() })
		})
	}
}

func TestScannerLinesTo(t *testing.T) {
	type test struct {
		input []byte
		fn    func(Scanner[any]) any
		want  []any
	}

	tests := []test{
		{
			input: []byte(""),
			fn:    func(s Scanner[any]) any { return s.String() },
		},
		{
			input: []byte("a"),
			fn:    func(s Scanner[any]) any { return s.String() },
			want:  []any{"a"},
		},
		{
			input: []byte("a\nb\nc"),
			fn:    func(s Scanner[any]) any { return s.String() },
			want:  []any{"a", "b", "c"},
		},
		{
			input: []byte("1\n2\n3"),
			fn:    func(s Scanner[any]) any { return s.Int() },
			want:  []any{1, 2, 3},
		},
		{
			input: []byte("1,2\n3,4"),
			fn:    func(s Scanner[any]) any { return Point2D{X: s.Int(), Y: s.Int()} },
			want:  []any{Point2D{X: 1, Y: 2}, Point2D{X: 3, Y: 4}},
		},
	}
	for _, test := range tests {
		t.Run("LinesTo", func(t *testing.T) {
			t.Run(string(test.input), func(t *testing.T) {
				var fn = func(s string) any { return test.fn(Scanner[any](s)) }

				scanner := Scanner[any](test.input)
				assert.Equal(t, scanner.LinesTo(fn), test.want)
			})
		})

		t.Run("LinesToS", func(t *testing.T) {
			t.Run(string(test.input), func(t *testing.T) {
				scanner := Scanner[any](test.input)
				assert.Equal(t, scanner.LinesToS(test.fn), test.want)
			})
		})
	}
}

func TestScannerOneOf(t *testing.T) {
	type test struct {
		input     []byte
		options   []string
		want      string
		remaining []byte
	}

	tests := []test{
		{
			input:     []byte("a"),
			options:   []string{"a", "b", "c"},
			want:      "a",
			remaining: []byte{},
		},
		{
			input:     []byte("b"),
			options:   []string{"a", "b", "c"},
			want:      "b",
			remaining: []byte{},
		},
		{
			input:     []byte("c"),
			options:   []string{"a", "b", "c"},
			want:      "c",
			remaining: []byte{},
		},
		{
			input:     []byte("a b c"),
			options:   []string{"a"},
			want:      "a",
			remaining: []byte(" b c"),
		},
		{
			input:     []byte(" a b c"),
			options:   []string{"a"},
			want:      "a",
			remaining: []byte(" b c"),
		},
		{
			input:     []byte("one word"),
			options:   []string{"one", "two words"},
			want:      "one",
			remaining: []byte(" word"),
		},
		{
			input:     []byte("two words"),
			options:   []string{"one", "two words"},
			want:      "two words",
			remaining: []byte{},
		},
	}
	for _, test := range tests {
		t.Run(string(test.input), func(t *testing.T) {
			scanner := Scanner[any](test.input)
			assert.Equal(t, test.want, scanner.OneOf(test.options...))
			assert.Equal(t, test.remaining, scanner)
		})
	}
}

func TestScannerOneOfError(t *testing.T) {
	type test struct {
		input   []byte
		options []string
	}

	tests := []test{
		{
			input: []byte(""),
		},
		{
			input: nil,
		},
		{
			input:   []byte("a b c"),
			options: []string{"x", "y", "z"},
		},
	}
	for _, test := range tests {
		t.Run(string(test.input), func(t *testing.T) {
			scanner := Scanner[any](test.input)
			assert.Panics(t, func() { scanner.OneOf(test.options...) })
		})
	}
}

func TestScannerRemove(t *testing.T) {
	type test struct {
		input     []byte
		remove    []string
		remaining []byte
	}

	tests := []test{
		{input: []byte("abc"), remove: []string{"a"}, remaining: []byte("bc")},
		{input: []byte("abc"), remove: []string{"b"}, remaining: []byte("ac")},
		{input: []byte("abc"), remove: []string{"c"}, remaining: []byte("ab")},
		{input: []byte("abc"), remove: []string{"d"}, remaining: []byte("abc")},
		{input: []byte("abc"), remove: []string{"ab"}, remaining: []byte("c")},
		{input: []byte("abc"), remove: []string{"abc"}, remaining: []byte("")},
		{input: []byte("abc"), remove: []string{"a", "c"}, remaining: []byte("b")},
	}
	for _, test := range tests {
		t.Run(string(test.input), func(t *testing.T) {
			scanner := Scanner[any](test.input)
			scanner.Remove(test.remove...)
			assert.Equal(t, test.remaining, []byte(scanner))
		})
	}
}

func TestScannerScanf(t *testing.T) {
	type test struct {
		input  []byte
		format string
		want   []any
	}

	tests := []test{
		{
			input:  []byte("abc"),
			format: "abc",
		},
		{
			input:  []byte("0"),
			format: "%d",
			want:   []any{0},
		},
		{
			input:  []byte("123"),
			format: "%d",
			want:   []any{123},
		},
		{
			input:  []byte("-123"),
			format: "%d",
			want:   []any{-123},
		},
		{
			input:  []byte("abc"),
			format: "%s",
			want:   []any{"abc"},
		},
		{
			input:  []byte("a b c"),
			format: "%s",
			want:   []any{"a b c"},
		},
		{
			input:  []byte("abc"),
			format: "%w",
			want:   []any{"abc"},
		},
		{
			input:  []byte("abc = 123"),
			format: "%s = %d",
			want:   []any{"abc", 123},
		},
		{
			input:  []byte("abc=123"),
			format: "%s=%d",
			want:   []any{"abc", 123},
		},
		{
			input:  []byte("abc def"),
			format: "%s %s",
			want:   []any{"abc", "def"},
		},
		{
			input:  []byte("abc def"),
			format: "%w %w",
			want:   []any{"abc", "def"},
		},
		{
			input:  []byte("10%"),
			format: "%d%%",
			want:   []any{10},
		},
		{
			input:  []byte("[2] .=* (abc.def)"),
			format: "[%d] .=* (%s.%w)",
			want:   []any{2, "abc", "def"},
		},
		{
			input:  []byte("abc 123"),
			format: "%s",
			want:   []any{Scanner[any]("abc 123")},
		},
	}
	for _, test := range tests {
		t.Run(string(test.input), func(t *testing.T) {
			scanner := Scanner[any](test.input)

			var args []any
			for _, expected := range test.want {
				switch expected.(type) {
				case int:
					var arg int
					args = append(args, &arg)
				case string:
					var arg string
					args = append(args, &arg)
				case Scanner[any]:
					var arg Scanner[any]
					args = append(args, &arg)
				default:
					panic("unsupported type")
				}
			}

			scanner.Scanf(test.format, args...)

			for i := range test.want {
				switch expected := test.want[i].(type) {
				case int:
					assert.Equal(t, expected, *args[i].(*int))
				case string:
					assert.Equal(t, expected, *args[i].(*string))
				case Scanner[any]:
					assert.Equal(t, expected, *args[i].(*Scanner[any]))
				}
			}
		})
	}
}

func TestScannerScanfError(t *testing.T) {
	type test struct {
		input  []byte
		format string
		want   []any
	}

	tests := []test{
		// Bad scan verb
		{
			format: "%v",
		},

		// Missing args
		{
			input:  []byte("abc"),
			format: "%s",
		},

		// Mismatched number of args
		{
			input:  []byte("abc"),
			format: "%s",
			want:   []any{"abc", "def"},
		},

		// No match
		{
			input:  []byte("abc"),
			format: "def",
		},

		// Unsupported argument type
		{
			input:  []byte("123"),
			format: "%d",
			want:   []any{int64(123)},
		},
	}
	for _, test := range tests {
		t.Run(string(test.input), func(t *testing.T) {
			scanner := Scanner[any](test.input)

			var args []any
			for _, expected := range test.want {
				switch expected.(type) {
				case int:
					var arg int
					args = append(args, &arg)
				case int64: // Allow this only to test an unsupported type
					var arg int64
					args = append(args, &arg)
				case string:
					var arg string
					args = append(args, &arg)
				default:
					panic("unsupported type")
				}
			}

			assert.Panics(t, func() { scanner.Scanf(test.format, args...) })
		})
	}
}

func TestScannerString(t *testing.T) {
	type test struct {
		input     []byte
		want      string
		remaining []byte
	}

	tests := []test{
		{
			input:     []byte("a"),
			want:      "a",
			remaining: []byte{},
		},
		{
			input:     []byte("abc"),
			want:      "abc",
			remaining: []byte{},
		},
		{
			input:     []byte("hello world"),
			want:      "hello",
			remaining: []byte("world"),
		},
		{
			input:     []byte("a b c"),
			want:      "a",
			remaining: []byte("b c"),
		},
		{
			input:     []byte("a\nb\nc"),
			want:      "a",
			remaining: []byte("b\nc"),
		},
	}
	for _, test := range tests {
		t.Run(string(test.input), func(t *testing.T) {
			scanner := Scanner[any](test.input)
			assert.Equal(t, test.want, scanner.String())
			assert.Equal(t, test.remaining, scanner)
		})
	}
}

func TestScannerStringError(t *testing.T) {
	type test struct {
		input []byte
	}

	tests := []test{
		{input: []byte("")},
		{input: []byte(" ")},
		{input: nil},
	}
	for _, test := range tests {
		t.Run(string(test.input), func(t *testing.T) {
			scanner := Scanner[any](test.input)
			assert.Panics(t, func() { scanner.String() })
		})
	}
}

//
// The following tests verify that parsing some real world examples works as expected.
//

func TestInputParsingYear2015Day04(t *testing.T) {
	in := Scanner[any]("line\n")

	line := in.Line()

	assert.Equal(t, "line", line)
}

func TestInputParsingYear2015Day05(t *testing.T) {
	in := Scanner[any]("line1\nline2\nline3")

	var lines []string
	for in.HasNext() {
		lines = append(lines, in.Line())
	}

	assert.Equal(t, 3, len(lines))
	assert.Equal(t, "line1", lines[0])
	assert.Equal(t, "line2", lines[1])
	assert.Equal(t, "line3", lines[2])
}

func TestInputParsingYear2015Day02(t *testing.T) {
	in := Scanner[any]("4x23x21\n22x29x19\n11x4x11")

	type box struct{ L, W, H int }
	var boxes []box
	for in.HasNext() {
		boxes = append(boxes, box{L: in.Int(), W: in.Int(), H: in.Int()})
	}

	assert.Equal(t, box{4, 23, 21}, boxes[0])
	assert.Equal(t, box{22, 29, 19}, boxes[1])
	assert.Equal(t, box{11, 4, 11}, boxes[2])
}

func TestInputParsingYear2015Day06(t *testing.T) {
	in := Scanner[any]("toggle 1,2 through 3,4\nturn on 5,6 through 7,8\nturn off 9,10 through 11,12")

	type instruction struct {
		op         string
		start, end Point2D
	}
	var instructions []instruction
	for in.HasNext() {
		instructions = append(instructions, instruction{
			op:    in.OneOf("toggle", "turn on", "turn off"),
			start: Point2D{X: in.Int(), Y: in.Int()},
			end:   Point2D{X: in.Int(), Y: in.Int()},
		})
	}

	assert.Equal(t, instruction{"toggle", Point2D{X: 1, Y: 2}, Point2D{X: 3, Y: 4}}, instructions[0])
	assert.Equal(t, instruction{"turn on", Point2D{X: 5, Y: 6}, Point2D{X: 7, Y: 8}}, instructions[1])
	assert.Equal(t, instruction{"turn off", Point2D{X: 9, Y: 10}, Point2D{X: 11, Y: 12}}, instructions[2])
}

func TestInputParsingYear2015Day19(t *testing.T) {
	in := Scanner[any]("A => B\nB => C\n\nABC")

	var replacements = make(map[string]string)

	var chunk = in.ChunkS()
	for chunk.HasNext() {
		lhs, rhs := chunk.Cut(" => ")
		replacements[lhs] = rhs
	}

	var target = in.String()

	assert.Equal(t, 2, len(replacements))
	assert.Equal(t, replacements["A"], "B")
	assert.Equal(t, replacements["B"], "C")
	assert.Equal(t, target, "ABC")
}
