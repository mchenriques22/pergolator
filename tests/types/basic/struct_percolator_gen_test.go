package basic

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/antoninferrand/pergolator/modifiers"
	"github.com/antoninferrand/pergolator/tree"
	"github.com/antoninferrand/pergolator/tree/defaultparser"
)

func TestPercolate(t *testing.T) {
	tests := []struct {
		input     string
		modifiers []tree.Modifiers
		document  Struct
		expected  bool
	}{
		{
			input:    "BasicString:my-name",
			document: Struct{BasicString: "my-name"},
			expected: true,
		},
		{
			input:    "BasicString:my-name",
			document: Struct{BasicString: "not-my-name"},
			expected: false,
		},
		{
			input:    "-BasicString:my-name",
			document: Struct{BasicString: "my-name"},
			expected: false,
		},
		{
			input:    "BasicString:my-name AND BasicInt32:0",
			document: Struct{BasicString: "my-name", BasicInt32: 0},
			expected: true,
		},
		{
			input:    "BasicString:my-name AND BasicInt32:1",
			document: Struct{BasicString: "my-name", BasicInt32: 0},
			expected: false,
		},
		{
			input:    "BasicInt32:>=0 AND BasicInt64:<=128",
			document: Struct{BasicInt32: 0, BasicInt64: 127},
			expected: true,
		},
		{
			input:    "BasicInt32:>=0 AND BasicInt64:<=128",
			document: Struct{BasicInt32: 0, BasicInt64: 129},
			expected: false,
		},
		{
			input:    "BasicBool:true",
			document: Struct{BasicBool: true},
			expected: true,
		},
		{
			input:    "BasicBool:false",
			document: Struct{BasicBool: false},
			expected: true,
		},
		{
			input:    "BasicBool:true",
			document: Struct{BasicBool: false},
		},
		{
			input:    "BasicString:a OR BasicString:b",
			document: Struct{BasicString: "a"},
			expected: true,
		},
		{
			input:    "BasicString:a OR BasicString:b",
			document: Struct{BasicString: "b"},
			expected: true,
		},
		{
			input:    "BasicString:a OR BasicString:b",
			document: Struct{BasicString: "c"},
			expected: false,
		},
		{
			input:    "BasicString:a OR BasicInt32:0",
			document: Struct{BasicInt32: 0},
			expected: true,
		},
		{
			input:    "BasicString:my-name AND BasicInt32:0 OR BasicString:else",
			document: Struct{BasicString: "my-name", BasicInt32: 0},
			expected: true,
		},
		{
			input:    "BasicString:my-name AND BasicInt32:0 OR BasicString:else",
			document: Struct{BasicString: "else"},
			expected: true,
		},
		{
			input:     "BasicString:MY-NAME",
			document:  Struct{BasicString: "my-name"},
			modifiers: []tree.Modifiers{modifiers.FormatValues(strings.ToLower)},
			expected:  true,
		},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			percolator, err := NewStructPercolator(defaultparser.Parse, test.input, test.modifiers...)
			require.NoError(t, err)

			assert.Equal(t, test.expected, percolator.Percolate(&test.document))
		})
	}
}

func BenchmarkBasicTypes(b *testing.B) {
	tests := []struct {
		name     string
		query    string
		document Struct
	}{
		{
			name:     "bool",
			query:    "BasicBool:true",
			document: Struct{BasicBool: true},
		},
		{
			name:     "same string",
			query:    "BasicString:my-name",
			document: Struct{BasicString: "my-name"},
		},
		{
			name:     "different string with same length",
			query:    "BasicString:medium_string",
			document: Struct{BasicString: "medium-string"},
		},
		{
			name:     "different string with different length",
			query:    "BasicString:long_string",
			document: Struct{BasicString: "string"},
		},
		{
			name:     "comparison string",
			query:    "BasicString:>=my-name",
			document: Struct{BasicString: "my"},
		},
		{
			name:     "same int32",
			query:    "BasicInt32:32",
			document: Struct{BasicInt32: 32},
		},
		{
			name:     "different int32",
			query:    "BasicInt32:32",
			document: Struct{BasicInt32: 64},
		},
		{
			name:     "comparison int32",
			query:    "BasicInt32:>=32",
			document: Struct{BasicInt32: 35},
		},
		{
			name:     "comparison int32 with multiple ANDs matching",
			query:    "BasicInt32:>=32 AND BasicInt32:>=10 AND BasicInt32:<=50 AND BasicInt32:<=100 AND BasicInt32:>0",
			document: Struct{BasicInt32: 35},
		},
	}

	for _, test := range tests {
		b.Run(test.name, func(b *testing.B) {
			percolator, err := NewStructPercolator(defaultparser.Parse, test.query)
			require.NoError(b, err)

			for b.Loop() {
				percolator.Percolate(&test.document)
			}
			b.ReportAllocs()
		})
	}
}
