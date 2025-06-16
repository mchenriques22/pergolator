//go:generate go run github.com/mchenriques22/pergolator github.com/mchenriques22/pergolator/tests/benchmark/boxing.Root --max-depth 5
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
