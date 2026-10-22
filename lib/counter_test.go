package lib

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert/v2"
)

func TestFrequencyCounter_ZeroValue(t *testing.T) {
	var fc FrequencyCounter[string]
	assert.Equal(t, 0, fc.GetCount("a"))
	assert.Equal(t, 0, len(fc.Entries()))
}

func TestFrequencyCounter_GetCount(t *testing.T) {
	type test struct {
		adds  []string
		value string
		want  int
	}

	tests := []test{
		{adds: []string{}, value: "a", want: 0},
		{adds: []string{"a"}, value: "a", want: 1},
		{adds: []string{"a"}, value: "b", want: 0},
		{adds: []string{"a", "a", "a"}, value: "a", want: 3},
		{adds: []string{"a", "b", "a", "c", "a"}, value: "a", want: 3},
		{adds: []string{"a", "b", "a", "c", "a"}, value: "b", want: 1},
	}
	for _, test := range tests {
		t.Run(fmt.Sprintf("%v/%s", test.adds, test.value), func(t *testing.T) {
			var fc FrequencyCounter[string]
			for _, value := range test.adds {
				fc.Add(value)
			}
			assert.Equal(t, test.want, fc.GetCount(test.value))
		})
	}
}

func TestFrequencyCounter_AddWithCount(t *testing.T) {
	var fc FrequencyCounter[string]
	fc.AddWithCount("a", 5)
	fc.AddWithCount("b", 3)
	fc.AddWithCount("a", 2)
	fc.Add("b")

	assert.Equal(t, 7, fc.GetCount("a"))
	assert.Equal(t, 4, fc.GetCount("b"))
}

func TestFrequencyCounter_Entries(t *testing.T) {
	type test struct {
		adds []string
		want []FrequencyCount[string]
	}

	tests := []test{
		{
			adds: []string{"a"},
			want: []FrequencyCount[string]{{Value: "a", Count: 1}},
		},
		{
			adds: []string{"a", "b", "b", "c", "c", "c"},
			want: []FrequencyCount[string]{
				{Value: "c", Count: 3},
				{Value: "b", Count: 2},
				{Value: "a", Count: 1},
			},
		},
		{
			adds: []string{"c", "c", "c", "b", "b", "a"},
			want: []FrequencyCount[string]{
				{Value: "c", Count: 3},
				{Value: "b", Count: 2},
				{Value: "a", Count: 1},
			},
		},
	}
	for _, test := range tests {
		t.Run(fmt.Sprintf("%v", test.adds), func(t *testing.T) {
			var fc FrequencyCounter[string]
			for _, value := range test.adds {
				fc.Add(value)
			}
			assertEntries(t, test.want, fc.Entries())
		})
	}
}

func TestFrequencyCounter_EntriesIsRepeatable(t *testing.T) {
	var fc FrequencyCounter[string]
	for _, value := range []string{"a", "b", "b", "c", "c", "c"} {
		fc.Add(value)
	}

	assertEntries(t, fc.Entries(), fc.Entries())
}

// Regression test: Entries pops from a copy of the heap, but the copy shares
// the *FrequencyCount pointers with the original, so popping rewrites their
// heap indices.  A subsequent update then calls heap.Fix with a stale index
// and the heap is left out of order.
func TestFrequencyCounter_AddAfterEntries(t *testing.T) {
	var fc FrequencyCounter[string]
	fc.Add("a")
	fc.AddWithCount("b", 2)
	_ = fc.Entries()
	fc.AddWithCount("a", 5)

	want := []FrequencyCount[string]{
		{Value: "a", Count: 6},
		{Value: "b", Count: 2},
	}
	assertEntries(t, want, fc.Entries())
}

// assertEntries compares just the exported fields of two entry lists.
func assertEntries[T comparable](t *testing.T, want, got []FrequencyCount[T]) {
	t.Helper()

	strip := func(entries []FrequencyCount[T]) []FrequencyCount[T] {
		var stripped []FrequencyCount[T]
		for _, e := range entries {
			stripped = append(stripped, FrequencyCount[T]{Value: e.Value, Count: e.Count})
		}
		return stripped
	}
	assert.Equal(t, strip(want), strip(got))
}
