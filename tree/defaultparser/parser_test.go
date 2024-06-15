package defaultparser

import (
	"errors"
	"testing"

	"github.com/antoninferrand/pergolator/tree/defaultparser/lexer"
)

func TestParser_Parse(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			input:    "a:a",
			expected: "a:a",
		},
		{
			input:    "a:a AND b:b",
			expected: "a:a AND b:b",
		},
		{
			input:    "a:a AND b:b AND c:c OR d:d",
			expected: "a:a AND b:b AND c:c OR d:d",
		},
		{
			input:    "a:a AND b:b OR c:c AND d:d",
			expected: "a:a AND b:b OR c:c AND d:d",
		},
		{
			input:    "a:a OR b:b AND c:c OR d:d",
			expected: "a:a OR b:b AND c:c OR d:d",
		},
		{
			input:    "a:a AND b:b OR c:c OR d:d",
			expected: "a:a AND b:b OR c:c OR d:d",
		},
		{
			input:    "a:a AND b:b AND c:c OR d:d OR e:e AND f:f",
			expected: "a:a AND b:b AND c:c OR d:d OR e:e AND f:f",
		},
		{
			input:    "(a:a)",
			expected: "a:a",
		},
		{
			input:    "(a:a AND b:b)",
			expected: "a:a AND b:b",
		},
		{
			input:    "(a:a OR b:b) AND c:c",
			expected: "(a:a OR b:b) AND c:c",
		},
		{
			input:    "a:a AND (b:b OR c:c)",
			expected: "a:a AND (b:b OR c:c)",
		},
		{
			input:    "a:a OR (b:b AND (c:c OR d:d))",
			expected: "a:a OR b:b AND (c:c OR d:d)",
		},
		{
			input:    "-a:a",
			expected: "NOT(a:a)",
		},
		{
			input:    "NOT a:a",
			expected: "NOT(a:a)",
		},
		{
			input:    "-a:a AND b:b",
			expected: "(NOT(a:a)) AND b:b",
		},
		{
			input:    "-a:a OR b:b",
			expected: "(NOT(a:a)) OR b:b",
		},
		{
			input:    "-(a:a AND b:b)",
			expected: "NOT(a:a AND b:b)",
		},
		{
			input:    "key:>0",
			expected: "key:>0",
		},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			p := newParser(lexer.New(tt.input))
			expr, err := p.parse()
			if err != nil {
				t.Fatalf("got an unexpected error: %v", err)
			}

			if tt.expected != expr.String() {
				t.Errorf("expected %v, got %v", tt.expected, expr.String())
			}
		})
	}
}

func TestParser_ParseErrors(t *testing.T) {
	tests := []struct {
		input    string
		expected error
	}{
		{
			input:    "a:a AND",
			expected: ErrUnexpectedToken,
		},
		{
			input:    "a:",
			expected: ErrExpectedValue,
		},
		{
			input:    ":a",
			expected: ErrUnexpectedToken,
		},
		{
			input:    "AND OR",
			expected: ErrUnexpectedToken,
		},
		{
			input:    "a AND b:b",
			expected: ErrExpectedSign,
		},
		{
			input:    "a*a",
			expected: ErrExpectedSign,
		},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			p := newParser(lexer.New(tt.input))
			_, err := p.parse()
			if err == nil {
				t.Fatalf("expected an error")
			}

			if !errors.Is(err, tt.expected) {
				t.Errorf("expected %v, got %v", tt.expected, err)
			}
		})
	}
}

func BenchmarkParse(b *testing.B) {
	runes := []rune("a:a AND b:b AND c:c OR d:d OR e:e AND f:f")
	for b.Loop() {
		p := newParser(lexer.NewFromRunes(runes))
		_, _ = p.parse()
	}
}
