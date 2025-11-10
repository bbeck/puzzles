package in

import (
	"testing"

	"github.com/alecthomas/assert/v2"
)

func TestAsSharesByteSlices(t *testing.T) {
	var s Scanner[any] = []byte("hello:123 goodbye:456\nhello:456 goodbye:789")

	type KV map[string]int
	kvs := as[any, KV](&s).LinesToS(
		func(in Scanner[KV]) KV {
			kvs := make(KV)
			for _, field := range in.FieldsS() {
				lhs, rhs := field.CutS(":")
				kvs[lhs.String()] = rhs.Int()
			}
			return kvs
		},
	)

	assert.Equal(t, 123, kvs[0]["hello"])
	assert.Equal(t, 456, kvs[0]["goodbye"])
	assert.Equal(t, 456, kvs[1]["hello"])
	assert.Equal(t, 789, kvs[1]["goodbye"])
	assert.Equal(t, []byte{}, s)
}
