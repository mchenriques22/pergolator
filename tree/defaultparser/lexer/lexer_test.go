package lexer

import (
	"testing"

	"github.com/antoninferrand/pergolator/tree/defaultparser/token"
)

func TestLexer_NextToken(t *testing.T) {
	tests := []struct {
		input    string
		expected []token.Token
	}{
		{
			input:    "",
			expected: []token.Token{},
		},
		{
			input: "strings:my-strings",
			expected: []token.Token{
				{Type: token.String, Literal: "strings", Start: 0, End: 7},
				{Type: token.EQ, Literal: ":", Start: 7, End: 8},
				{Type: token.String, Literal: "my-strings", Start: 8, End: 18},
			},
		},
		{
			input: "numeric:>0",
			expected: []token.Token{
				{Type: token.String, Literal: "numeric", Start: 0, End: 7},
				{Type: token.GT, Literal: ":>", Start: 7, End: 9},
				{Type: token.String, Literal: "0", Start: 9, End: 10},
			},
		},
		{
			input: "numeric:<=123",
			expected: []token.Token{
				{Type: token.String, Literal: "numeric", Start: 0, End: 7},
				{Type: token.LTE, Literal: ":<=", Start: 7, End: 10},
				{Type: token.String, Literal: "123", Start: 10, End: 13},
			},
		},
		{
			input: "strings:my-strings AND mynumber:1234",
			expected: []token.Token{
				{Type: token.String, Literal: "strings", Start: 0, End: 7},
				{Type: token.EQ, Literal: ":", Start: 7, End: 8},
				{Type: token.String, Literal: "my-strings", Start: 8, End: 18},
				{Type: token.AND, Literal: "AND", Start: 19, End: 22},
				{Type: token.String, Literal: "mynumber", Start: 23, End: 31},
				{Type: token.EQ, Literal: ":", Start: 31, End: 32},
				{Type: token.String, Literal: "1234", Start: 32, End: 36},
			},
		},
		{
			input: "strings:my-strings AND (mynumber:1234 OR meta.match:true)",
			expected: []token.Token{
				{Type: token.String, Literal: "strings", Start: 0, End: 7},
				{Type: token.EQ, Literal: ":", Start: 7, End: 8},
				{Type: token.String, Literal: "my-strings", Start: 8, End: 18},
				{Type: token.AND, Literal: "AND", Start: 19, End: 22},
				{Type: token.LeftParenthesis, Literal: "(", Start: 23, End: 24},
				{Type: token.String, Literal: "mynumber", Start: 24, End: 32},
				{Type: token.EQ, Literal: ":", Start: 32, End: 33},
				{Type: token.String, Literal: "1234", Start: 33, End: 37},
				{Type: token.OR, Literal: "OR", Start: 38, End: 40},
				{Type: token.String, Literal: "meta.match", Start: 41, End: 51},
				{Type: token.EQ, Literal: ":", Start: 51, End: 52},
				{Type: token.String, Literal: "true", Start: 52, End: 56},
				{Type: token.RightParenthesis, Literal: ")", Start: 56, End: 57},
			},
		},
		{
			input: "strings:my-strings AND (mynumber:1234 OR foo:bar AND hello:world)",
			expected: []token.Token{
				{Type: token.String, Literal: "strings", Start: 0, End: 7},
				{Type: token.EQ, Literal: ":", Start: 7, End: 8},
				{Type: token.String, Literal: "my-strings", Start: 8, End: 18},
				{Type: token.AND, Literal: "AND", Start: 19, End: 22},
				{Type: token.LeftParenthesis, Literal: "(", Start: 23, End: 24},
				{Type: token.String, Literal: "mynumber", Start: 24, End: 32},
				{Type: token.EQ, Literal: ":", Start: 32, End: 33},
				{Type: token.String, Literal: "1234", Start: 33, End: 37},
				{Type: token.OR, Literal: "OR", Start: 38, End: 40},
				{Type: token.String, Literal: "foo", Start: 41, End: 44},
				{Type: token.EQ, Literal: ":", Start: 44, End: 45},
				{Type: token.String, Literal: "bar", Start: 45, End: 48},
				{Type: token.AND, Literal: "AND", Start: 49, End: 52},
				{Type: token.String, Literal: "hello", Start: 53, End: 58},
				{Type: token.EQ, Literal: ":", Start: 58, End: 59},
				{Type: token.String, Literal: "world", Start: 59, End: 64},
				{Type: token.RightParenthesis, Literal: ")", Start: 64, End: 65},
			},
		},
		{
			input: "strings:my-strings AND ((mynumber:1234 OR foo:bar) AND hello:world) AND meta.match:true",
			expected: []token.Token{
				{Type: token.String, Literal: "strings", Start: 0, End: 7},
				{Type: token.EQ, Literal: ":", Start: 7, End: 8},
				{Type: token.String, Literal: "my-strings", Start: 8, End: 18},
				{Type: token.AND, Literal: "AND", Start: 19, End: 22},
				{Type: token.LeftParenthesis, Literal: "(", Start: 23, End: 24},
				{Type: token.LeftParenthesis, Literal: "(", Start: 24, End: 25},
				{Type: token.String, Literal: "mynumber", Start: 25, End: 33},
				{Type: token.EQ, Literal: ":", Start: 33, End: 34},
				{Type: token.String, Literal: "1234", Start: 34, End: 38},
				{Type: token.OR, Literal: "OR", Start: 39, End: 41},
				{Type: token.String, Literal: "foo", Start: 42, End: 45},
				{Type: token.EQ, Literal: ":", Start: 45, End: 46},
				{Type: token.String, Literal: "bar", Start: 46, End: 49},
				{Type: token.RightParenthesis, Literal: ")", Start: 49, End: 50},
				{Type: token.AND, Literal: "AND", Start: 51, End: 54},
				{Type: token.String, Literal: "hello", Start: 55, End: 60},
				{Type: token.EQ, Literal: ":", Start: 60, End: 61},
				{Type: token.String, Literal: "world", Start: 61, End: 66},
				{Type: token.RightParenthesis, Literal: ")", Start: 66, End: 67},
				{Type: token.AND, Literal: "AND", Start: 68, End: 71},
				{Type: token.String, Literal: "meta.match", Start: 72, End: 82},
				{Type: token.EQ, Literal: ":", Start: 82, End: 83},
				{Type: token.String, Literal: "true", Start: 83, End: 87},
			},
		},
		{
			input: "-strings:mystrings",
			expected: []token.Token{
				{Type: token.NOT, Literal: "-", Start: 0, End: 1},
				{Type: token.String, Literal: "strings", Start: 1, End: 8},
				{Type: token.EQ, Literal: ":", Start: 8, End: 9},
				{Type: token.String, Literal: "mystrings", Start: 9, End: 18},
			},
		},
		{
			input: "strings-with-dash:value",
			expected: []token.Token{
				{Type: token.String, Literal: "strings-with-dash", Start: 0, End: 17},
				{Type: token.EQ, Literal: ":", Start: 17, End: 18},
				{Type: token.String, Literal: "value", Start: 18, End: 23},
			},
		},
		{
			input: "strings:value-with-dash",
			expected: []token.Token{
				{Type: token.String, Literal: "strings", Start: 0, End: 7},
				{Type: token.EQ, Literal: ":", Start: 7, End: 8},
				{Type: token.String, Literal: "value-with-dash", Start: 8, End: 23},
			},
		},
		{
			input: `strings:"-a:a AND \n"`,
			expected: []token.Token{
				{Type: token.String, Literal: "strings", Start: 0, End: 7},
				{Type: token.EQ, Literal: ":", Start: 7, End: 8},
				{Type: token.String, Literal: `-a:a AND \n`, Start: 8, End: 20},
			},
		},
		{
			input: "-strings:value",
			expected: []token.Token{
				{Type: token.NOT, Literal: "-", Start: 0, End: 1},
				{Type: token.String, Literal: "strings", Start: 1, End: 8},
				{Type: token.EQ, Literal: ":", Start: 8, End: 9},
				{Type: token.String, Literal: "value", Start: 9, End: 14},
			},
		},
		{
			input: "NOT strings:value",
			expected: []token.Token{
				{Type: token.NOT, Literal: "NOT", Start: 0, End: 3},
				{Type: token.String, Literal: "strings", Start: 4, End: 11},
				{Type: token.EQ, Literal: ":", Start: 11, End: 12},
				{Type: token.String, Literal: "value", Start: 12, End: 17},
			},
		},
		{
			input: "key:>=100",
			expected: []token.Token{
				{Type: token.String, Literal: "key", Start: 0, End: 3},
				{Type: token.GTE, Literal: ":>=", Start: 3, End: 6},
				{Type: token.String, Literal: "100", Start: 6, End: 9},
			},
		},
		{
			input: "key:<10",
			expected: []token.Token{
				{Type: token.String, Literal: "key", Start: 0, End: 3},
				{Type: token.LT, Literal: ":<", Start: 3, End: 5},
				{Type: token.String, Literal: "10", Start: 5, End: 7},
			},
		},
		{
			input: "key:\"tag:value\"",
			expected: []token.Token{
				{Type: token.String, Literal: "key", Start: 0, End: 3},
				{Type: token.EQ, Literal: ":", Start: 3, End: 4},
				{Type: token.String, Literal: "tag:value", Start: 4, End: 14},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			l := New(tt.input)

			for _, expected := range tt.expected {
				tok := l.NextToken()

				assert(t, expected.Type, tok.Type)
				assert(t, expected.Literal, tok.Literal)
				assert(t, expected.Start, tok.Start)
				assert(t, expected.End, tok.End)
			}

			assert(t, token.EOF, l.NextToken().Type)
		})
	}
}

func assert[T comparable](t *testing.T, expected, got T) {
	if expected != got {
		t.Errorf("expected %v, got %v", expected, got)
	}
}
