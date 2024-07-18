package saihon

import (
	"strings"
	"testing"
)

func TestLength(t *testing.T) {
	doc, _ := Parse(strings.NewReader(test_html))
	collection := doc.All()
	expect := 17
	actual := collection.Length()
	if actual != expect {
		t.Errorf("\ngot : %v, want: %v\n", actual, expect)
	}
}

func doNothing(*Element) {}

func BenchmarkFor(b *testing.B) {
	doc, _ := Parse(strings.NewReader(test_html))
	collection := doc.All()
	length := collection.Length()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for i := 0; i < length; i++ {
			doNothing(collection.Get(i))
		}
	}
}

func BenchmarkEnumerator(b *testing.B) {
	doc, _ := Parse(strings.NewReader(test_html))
	collection := doc.All()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for v := range collection.Enumerator() {
			doNothing(v)
		}
	}
}

func BenchmarkForEach(b *testing.B) {
	doc, _ := Parse(strings.NewReader(test_html))
	collection := doc.All()
	fn := func(v *Element, i int, a Collection) {}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		collection.ForEach(fn)
	}
}
