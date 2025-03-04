package in

import (
	"bytes"
	"fmt"
	. "github.com/bbeck/puzzles/lib"
	"regexp"
	"strings"
)

// Scanner is a wrapper around a slice of bytes that provides a convenient way
// to read and parse input data.
//
// The methods of a Scanner come in two flavors.  By default, methods receive
// and return strings since this is generally what the user desires.  However,
// where it makes sense methods also include a version with an S suffix that
// returns Scanner(s) instead of strings.  These versions are useful when the
// input needs to be parsed further.
//
// NOTE: The Scanner type unfortunately requires the use of an unused type
// parameters to allow some methods to have arguments that utilize a type
// parameter.
type Scanner[T any] []byte

// Byte returns the next byte from the scanner.
func (bs *Scanner[T]) Byte() byte {
	if len(*bs) == 0 {
		panic("no more bytes")
	}

	var b byte
	b, *bs = (*bs)[0], (*bs)[1:]
	return b
}

// Bytes returns all remaining bytes from the scanner.
func (bs *Scanner[T]) Bytes() []byte {
	if len(*bs) == 0 {
		panic("no more bytes")
	}

	var cs = []byte(*bs)
	*bs = []byte{}
	return cs
}

// Chunk returns the next sequence of lines in the Scanner up until one or more
// blank lines.
func (bs *Scanner[T]) Chunk() []string {
	if len(*bs) == 0 {
		panic("no more bytes")
	}

	var lines []string

	// Consume any leading newline characters
	for bs.HasPrefix("\n") {
		*bs = (*bs)[1:]
	}

	// Collect lines until we run into an empty line or the end of the bytes
	for len(*bs) > 0 {
		line := bs.Line()
		if line == "" {
			break
		}
		lines = append(lines, line)
	}

	// Consume any remaining newline characters
	for bs.HasPrefix("\n") {
		*bs = (*bs)[1:]
	}

	return lines
}

// ChunkS returns a scanner for the next sequence of lines in the Scanner up
// until one or more blank lines.
func (bs *Scanner[T]) ChunkS() Scanner[T] {
	return Scanner[T](strings.Join(bs.Chunk(), "\n"))
}

// Cut splits the current line into two parts at the first occurrence of sep.
// If sep is not found, the second part is empty.
func (bs *Scanner[T]) Cut(sep string) (string, string) {
	if sep == "" {
		panic("empty separator")
	}

	lhs, rhs, _ := strings.Cut(bs.Line(), sep)
	return lhs, rhs
}

// CutS splits the current line into two parts at the first occurrence of sep.
// If sep is not found, the second part is empty.
func (bs *Scanner[T]) CutS(sep string) (Scanner[T], Scanner[T]) {
	var lhs, rhs = bs.Cut(sep)
	return Scanner[T](lhs), Scanner[T](rhs)
}

// Expect ensures that the next string from the scanner is equal to s.
func (bs *Scanner[T]) Expect(s string) {
	if !bytes.HasPrefix(*bs, []byte(s)) {
		panic("unable to expect")
	}

	*bs = (*bs)[len(s):]
}

// Fields splits the current line into fields around one or more consecutive
// whitespace characters.
func (bs *Scanner[T]) Fields() []string {
	return strings.Fields(bs.Line())
}

// FieldsS splits the current line into fields around one or more consecutive
// whitespace characters.
func (bs *Scanner[T]) FieldsS() []Scanner[T] {
	var scanners []Scanner[T]
	for _, s := range bs.Fields() {
		scanners = append(scanners, Scanner[T](s))
	}
	return scanners
}

// Grid2D builds a Grid2D instance from the input using the provided function
// to determine the value of each cell.
func (bs *Scanner[T]) Grid2D(fn func(int, int, string) T) Grid2D[T] {
	var lines []string
	for bs.HasNext() {
		lines = append(lines, bs.Line())
	}

	var grid = NewGrid2D[T](len(lines[0]), len(lines))
	for y, line := range lines {
		for x, ch := range line {
			grid.Set(x, y, fn(x, y, string(ch)))
		}
	}
	return grid
}

// HasNext returns true if there are more non-whitespace bytes to read.
func (bs *Scanner[T]) HasNext() bool {
	for _, b := range *bs {
		if !isWhitespace(b) {
			return true
		}
	}

	return false
}

// HasPrefix returns true if the scanner starts with the given prefix.
func (bs *Scanner[T]) HasPrefix(prefix string) bool {
	return bytes.HasPrefix(*bs, []byte(prefix))
}

// Int returns the next integer from the scanner.
func (bs *Scanner[T]) Int() int {
	bs.skipUntilDigitCharacter()

	if len(*bs) == 0 {
		panic("no more bytes")
	}

	var isNegative bool
	if (*bs)[0] == '-' {
		isNegative = true
		*bs = (*bs)[1:]
	}

	var n int
	for len(*bs) > 0 {
		if (*bs)[0] < '0' || (*bs)[0] > '9' {
			break
		}

		n = n*10 + int((*bs)[0]-'0')
		*bs = (*bs)[1:]
	}

	if isNegative {
		n = -n
	}

	return n
}

// Ints returns all integers remaining in the scanner.
func (bs *Scanner[T]) Ints() []int {
	var ints []int
	for bs.HasNext() {
		bs.skipUntilDigitCharacter()

		if len(*bs) == 0 {
			break
		}

		ints = append(ints, bs.Int())
	}
	return ints
}

// Line returns the next line from the scanner.  The line is delimited by a
// newline except the last line which isn't required to have a newline.
func (bs *Scanner[T]) Line() string {
	if len(*bs) == 0 {
		panic("no more bytes")
	}

	var b byte
	var sb strings.Builder
	for len(*bs) > 0 {
		b, *bs = (*bs)[0], (*bs)[1:]
		if b == '\n' {
			break
		}
		sb.WriteByte(b)
	}

	return sb.String()
}

// LinesTo transforms each line in the scanner to an arbitrary type.  The
// arbitrary type is determined by the fn argument's return value.
func (bs *Scanner[T]) LinesTo(fn func(string) T) []T {
	var ts []T
	for bs.HasNext() {
		ts = append(ts, fn(bs.Line()))
	}
	return ts
}

// LinesToS transforms each line in the scanner to an arbirary type via a new
// scanner with just the line's contents.
func (bs *Scanner[T]) LinesToS(fn func(Scanner[T]) T) []T {
	return bs.LinesTo(func(s string) T {
		return fn(Scanner[T](s))
	})
}

// OneOf returns the next string from the scanner that matches one of the given
// options.
func (bs *Scanner[T]) OneOf(options ...string) string {
	// Consume any leading whitespace.
	for len(*bs) > 0 && isWhitespace((*bs)[0]) {
		*bs = (*bs)[1:]
	}
	var opts = SetFrom(options...)
	var sb strings.Builder

	for len(*bs) > 0 {
		sb.WriteByte((*bs)[0])
		*bs = (*bs)[1:]

		if opts.Contains(sb.String()) {
			return sb.String()
		}
	}

	panic("no matching option")
}

// Remove removes all occurrences of the given string(s) from the scanner.
func (bs *Scanner[T]) Remove(s ...string) {
	for _, remove := range s {
		*bs = []byte(strings.ReplaceAll(string(*bs), remove, ""))
	}
}

var scanfMemo = make(map[string]*regexp.Regexp)

// Scanf reads the next line from the scanner and parses it according to the
// given format string and arguments.  The format string may contain the
// following verbs:
//
//	%d: match a decimal integer
//	%s: match a string
//	%w: match a word
//	%%: match a literal '%'
//
// NOTE: This function uses a custom scanner to parse the input line to work
// around the limitation of the standard library's scan functions that tokens
// must be separated by whitespace.
func (bs *Scanner[T]) Scanf(format string, args ...any) {
	var regex = scanfMemo[format]
	if regex == nil {
		var sb strings.Builder
		sb.WriteString(`^`)

		for i := 0; i < len(format); i++ {
			switch format[i] {
			case '%':
				switch format[i+1] {
				case 'd':
					sb.WriteString(`(-?\d+)`)
					i++
				case 's':
					if len(format) > i+2 && format[i+2] == '?' {
						sb.WriteString(`(.*)`)
					} else {
						sb.WriteString(`(.+)`)
					}
					i++
				case 'w':
					sb.WriteString(`(\w+)`)
					i++
				case '%':
					sb.WriteString(`%`)
					i++
				default:
					panicf("unrecognized scan verb: %%%c", format[i+1])
				}

			default:
				sb.WriteString(regexp.QuoteMeta(string(format[i])))
			}
		}

		regex = regexp.MustCompile(sb.String())
		scanfMemo[format] = regex
	}

	line := bs.Line()
	matches := regex.FindStringSubmatch(line)
	if matches == nil {
		panicf("no match for line: %s", line)
	}

	// Drop the full match and keep only the captured groups.
	matches = matches[1:]
	if len(args) != len(matches) {
		panicf("mismatched number of arguments: %d != %d", len(args), len(matches))
	}

	for i := range args {
		switch v := args[i].(type) {
		case *int:
			*v = ParseInt(matches[i])
		case *string:
			*v = matches[i]
		default:
			panicf("unsupported type: %T", v)
		}
	}
}

// String returns the next string from the scanner.  The string is delimited by
// whitespace.
func (bs *Scanner[T]) String() string {
	// Consume any leading whitespace.
	for len(*bs) > 0 && isWhitespace((*bs)[0]) {
		*bs = (*bs)[1:]
	}

	if len(*bs) == 0 {
		panic("no more bytes")
	}

	// Consume the string up until the next whitespace.
	var sb strings.Builder
	for len(*bs) > 0 && !isWhitespace((*bs)[0]) {
		sb.WriteByte((*bs)[0])
		*bs = (*bs)[1:]
	}

	// Consume any trailing whitespace.
	for len(*bs) > 0 && isWhitespace((*bs)[0]) {
		*bs = (*bs)[1:]
	}

	return sb.String()
}

func (bs *Scanner[T]) skipUntilDigitCharacter() {
	for len(*bs) > 0 {
		if '0' <= (*bs)[0] && (*bs)[0] <= '9' {
			break
		}
		if (*bs)[0] == '-' && len(*bs) > 1 && '0' <= (*bs)[1] && (*bs)[1] <= '9' {
			break
		}
		*bs = (*bs)[1:]
	}
}

func isWhitespace(b byte) bool {
	return b == ' ' || b == '\f' || b == '\n' || b == '\r' || b == '\t' || b == '\v'
}

func panicf(format string, args ...interface{}) {
	panic(fmt.Sprintf(format, args...))
}
