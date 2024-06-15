//go:generate go run github.com/antoninferrand/pergolator github.com/antoninferrand/pergolator/tests/benchmark/boxing.Root --max-depth 5
package boxing

type Root struct {
	A A

	Value int64
}

type A struct {
	B B

	Value int64
}

type B struct {
	C C

	Value int64
}

type C struct {
	D D

	Value int64
}

type D struct {
	Value int64
}
