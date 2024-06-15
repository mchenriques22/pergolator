package benchmark

import (
	"reflect"
	"testing"
)

// String of length 8 bytes is different from a string of 4 bytes (different length)
func BenchmarkStringsDifferentLength(b *testing.B) {
	x := getString(8, -1)
	y := getString(4, -1)

	for b.Loop() {
		equal(x, y)
	}
}

// String of 64 bytes is different from other 64 bytes strings (1st byte is different, 2nd byte is different, 9th byte is different, 65th byte is different)
func BenchmarkDifferent128To1st(b *testing.B) {
	x := getString(128, -1)
	y := getString(128, 0)

	for b.Loop() {
		equal(x, y)
	}
}

func BenchmarkDifferent128To2nd(b *testing.B) {
	x := getString(128, -1)
	y := getString(128, 1)

	for b.Loop() {
		equal(x, y)
	}
}

func BenchmarkDifferent128To8th(b *testing.B) {
	x := getString(128, -1)
	y := getString(128, 7)

	for b.Loop() {
		equal(x, y)
	}
}

func BenchmarkDifferent128To16th(b *testing.B) {
	x := getString(128, -1)
	y := getString(128, 15)

	for b.Loop() {
		equal(x, y)
	}
}

func BenchmarkDifferent128To64th(b *testing.B) {
	x := getString(128, -1)
	y := getString(128, 63)

	for b.Loop() {
		equal(x, y)
	}
}

func BenchmarkDifferent128To128th(b *testing.B) {
	x := getString(128, -1)
	y := getString(128, 127)

	for b.Loop() {
		equal(x, y)
	}
}

func BenchmarkEqual4Bytes(b *testing.B) {
	x := getString(4, -1)
	y := getString(4, -1)

	for b.Loop() {
		equal(x, y)
	}
}

func BenchmarkEqual8Bytes(b *testing.B) {
	x := getString(8, -1)
	y := getString(8, -1)

	for b.Loop() {
		equal(x, y)
	}
}

func BenchmarkEqual32Bytes(b *testing.B) {
	x := getString(32, -1)
	y := getString(32, -1)

	for b.Loop() {
		equal(x, y)
	}
}

func BenchmarkEqual64Bytes(b *testing.B) {
	x := getString(64, -1)
	y := getString(64, -1)

	for b.Loop() {
		equal(x, y)
	}
}

func BenchmarkEqual128Bytes(b *testing.B) {
	x := getString(128, -1)
	y := getString(128, -1)

	for b.Loop() {
		equal(x, y)
	}
}

func getString(n int, flipIndex int) string {
	var s []rune
	for range n {
		s = append(s, 'a')
	}

	if flipIndex >= 0 && flipIndex < n {
		s[flipIndex] = 'b'
	}

	return string(s)
}

//go:noinline
func equal(a, b string) bool {
	return a == b
}

func BenchmarkFunction(b *testing.B) {
	wrapperFn := func() func(a string, b string) bool {
		return equal
	}

	for b.Loop() {
		if reflect.ValueOf(wrapperFn()).Pointer() != reflect.ValueOf(equal).Pointer() {
			b.Error("unexpected")
		}
	}
}
