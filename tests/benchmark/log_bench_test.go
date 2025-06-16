package benchmark

import (
	"math/rand/v2"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/mchenriques22/pergolator/tree/defaultparser"
)

var (
	err1 = "failed to connect to the database"
	err2 = "timeout while waiting for the response"
)

func generateLogs() []*Log {
	var logs []*Log
	for i := 0; i < 1000; i++ {
		var level string
		switch i % 3 {
		case 0:
			level = "info"
		case 1:
			level = "warn"
		case 2:
			level = "error"
		}

		var message string
		switch i % 2 {
		case 0:
			message = err1
		case 1:
			message = err2
		}

		logs = append(logs, &Log{
			TraceID:   rand.Uint64(),
			SpanID:    rand.Uint64(),
			Timestamp: rand.Int64(),
			Service:   "my-service",
			Level:     level,
			Message:   message,
			Error:     "id",
			Tags: map[string]string{
				"env":     "production",
				"region":  "us-west-1",
				"version": "1.0.0",
			},
		})
	}
	return logs
}

func BenchmarkMatchAll(b *testing.B) {
	logs := generateLogs()
	percolator, err := NewLogPercolator(defaultparser.Parse, "Service:my-service AND Level:info AND Tags.env:production")
	require.NoError(b, err)

	var counter int
	for b.Loop() {
		for _, log := range logs {
			if percolator.Percolate(log) {
				counter++
			}
		}
	}
	b.ReportAllocs()
	assert.Greater(b, counter, 0, "No logs matched the query")
}
