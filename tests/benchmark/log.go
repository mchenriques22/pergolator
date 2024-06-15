//go:generate go run github.com/antoninferrand/pergolator github.com/antoninferrand/pergolator/tests/benchmark.Log
package benchmark

type Log struct {
	TraceID   uint64
	SpanID    uint64
	Timestamp int64
	Service   string
	Level     string
	Message   string
	Error     string
	Tags      map[string]string
}
