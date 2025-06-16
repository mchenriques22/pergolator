package boxing

import (
	"testing"

	"github.com/mchenriques22/pergolator/tree/defaultparser"
)

func BenchmarkPercolateWithMultipleBoxingLevel(b *testing.B) {
	b.Run("0 boxing", func(b *testing.B) {
		percolator, _ := NewRootPercolator(defaultparser.Parse, "value:1")
		document := Root{Value: 1}
		if !percolator.Percolate(&document) {
			b.Fatalf("query is supposed to match the document")
		}

		for b.Loop() {
			_ = percolator.Percolate(&document)
		}
	})

	b.Run("1 boxing", func(b *testing.B) {
		percolator, _ := NewRootPercolator(defaultparser.Parse, "a.value:1")
		document := Root{A: A{Value: 1}}
		if !percolator.Percolate(&document) {
			b.Fatalf("query is supposed to match the document")
		}

		for b.Loop() {
			_ = percolator.Percolate(&document)
		}
	})

	b.Run("4 boxing", func(b *testing.B) {
		percolator, _ := NewRootPercolator(defaultparser.Parse, "a.b.c.d.value:1")
		document := Root{A: A{B: B{C: C{D: D{Value: 1}}}}}
		if !percolator.Percolate(&document) {
			b.Fatalf("query is supposed to match the document")
		}

		for b.Loop() {
			_ = percolator.Percolate(&document)
		}
	})
}
