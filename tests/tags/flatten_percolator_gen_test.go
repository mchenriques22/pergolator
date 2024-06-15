package tags

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/antoninferrand/pergolator/tree"
	"github.com/antoninferrand/pergolator/tree/defaultparser"
)

func TestPercolate(t *testing.T) {
	tests := []struct {
		input     string
		modifiers []tree.Modifiers
		document  MyStruct
		expected  bool
	}{
		{
			input:    "FieldA.FieldB:my-name",
			document: MyStruct{FieldA: MyStructB{FieldB: "my-name"}},
			expected: false,
		},
		{
			input:    "FieldB:my-name",
			document: MyStruct{FieldA: MyStructB{FieldB: "my-name"}},
			expected: true,
		},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			percolator, err := NewMyStructPercolator(defaultparser.Parse, test.input, test.modifiers...)
			require.NoError(t, err)

			result := percolator.Percolate(&test.document)
			assert.Equal(t, test.expected, result)
		})
	}
}
