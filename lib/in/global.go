package in

import (
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

		scanner = Scanner[any](fin)
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

func Chunk() Scanner[any] {
	initialize()
	return scanner.Chunk()
}

func Cut(sep string) (string, string) {
	initialize()
	return scanner.Cut(sep)
}

func Expect(s string) {
	initialize()
	scanner.Expect(s)
}

func Fields() []string {
	initialize()
	return scanner.Fields()
}

func HasNext() bool {
	initialize()
	return scanner.HasNext()
}

func ToGrid2D[T any](fn func(x, y int, s string) T) Grid2D[T] {
	initialize()
	return as[T]().Grid2D(fn)
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

func LinesTo[T any](fn func(*Scanner[T]) T) []T {
	initialize()
	return as[T]().LinesTo(fn)
}

func OneOf(options ...string) string {
	initialize()
	return scanner.OneOf(options...)
}

func Scanf(format string, a ...interface{}) {
	initialize()
	scanner.Scanf(format, a...)
}
func String() string {
	initialize()
	return scanner.String()
}

// As will reinterpret the scanner with the type T so that we can parse the data
// as/ a new type.  The underlying byte slices are shared between the scanners
// so data will only be read a single time.
func as[T any]() *Scanner[T] {
	var s = Scanner[T](scanner)
	return &s
}
