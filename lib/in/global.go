package in

import (
	"bytes"
	"os"

	. "github.com/bbeck/puzzles/lib"
)

var scanner Scanner[any]

// initialize will lazily initialize the scanner.  We do this so that when we
// try to infer the path to the input filename we're being called from a
// puzzle's `main.go` file.  This will allow us to use the call stack to infer
// the site, year, day and part.
func initialize() {
	if scanner == nil {
		fin, err := os.ReadFile(Filename())
		if err != nil {
			panicf("unable to read input.txt: %+v", err)
		}

		// Remove any trailing newline characters, but leave any other whitespace
		// intact.
		scanner = bytes.TrimRight(fin, "\n")
	}
}

func Byte() byte {
	initialize()
	return scanner.Byte()
}

func Bytes() []byte {
	initialize()
	return scanner.Bytes()
}

func Chunk() []string {
	initialize()
	return scanner.Chunk()
}

func ChunkS() Scanner[any] {
	initialize()
	return scanner.ChunkS()
}

func Cut(sep string) (string, string) {
	initialize()
	return scanner.Cut(sep)
}

func CutS[T any](sep string) (Scanner[T], Scanner[T]) {
	initialize()
	return as[any, T](&scanner).CutS(sep)
}

func Expect(s string) {
	initialize()
	scanner.Expect(s)
}

func Fields() []string {
	initialize()
	return scanner.Fields()
}

func FieldsS[T any]() []Scanner[T] {
	initialize()
	return as[any, T](&scanner).FieldsS()
}

func HasNext() bool {
	initialize()
	return scanner.HasNext()
}

func HasNextLine() bool {
	initialize()
	return scanner.HasNextLine()
}

func HasPrefix(prefix string) bool {
	initialize()
	return scanner.HasPrefix(prefix)
}

func Int() int {
	initialize()
	return scanner.Int()
}

func Ints() []int {
	initialize()
	return scanner.Ints()
}

func Line() string {
	initialize()
	return scanner.Line()
}

func Lines() []string {
	initialize()
	return scanner.Lines()
}

func LinesTo[T any](fn func(string) T) []T {
	initialize()
	return as[any, T](&scanner).LinesTo(fn)
}

func LinesToS[T any](fn func(Scanner[T]) T) []T {
	initialize()
	return as[any, T](&scanner).LinesToS(fn)
}

func OneOf(options ...string) string {
	initialize()
	return scanner.OneOf(options...)
}

func Remove(s ...string) {
	initialize()
	scanner.Remove(s...)
}

func Scanf(format string, a ...interface{}) {
	initialize()
	scanner.Scanf(format, a...)
}

func Split(sep string) []string {
	initialize()
	return scanner.Split(sep)
}

func SplitS[T any](sep string) []Scanner[T] {
	initialize()
	return as[any, T](&scanner).SplitS(sep)
}

func String() string {
	initialize()
	return scanner.String()
}

func ToGrid2D[T any](fn func(x, y int, s string) T) Grid2D[T] {
	initialize()
	return as[any, T](&scanner).Grid2D(fn)
}

func as[T, U any](s *Scanner[T]) *Scanner[U] {
	return (*Scanner[U])(s)
}
