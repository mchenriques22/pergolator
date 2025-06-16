package modifiers

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/mchenriques22/pergolator/tree"
	"github.com/mchenriques22/pergolator/tree/defaultparser"
	"github.com/mchenriques22/pergolator/modifiers"
)

func TestModifiers(t *testing.T) {
	tests := []struct {
		input     string
		document  Struct
		modifiers []tree.Modifiers
	}{
		{
			input:     "field_a:value",
			document:  Struct{FieldA: "value"},
			modifiers: []tree.Modifiers{modifiers.FormatKeysToCamelCase},
		},
		{
			input:     "FIELD_B:VALUE",
			document:  Struct{FieldB: "value"},
			modifiers: []tree.Modifiers{modifiers.FormatKeysToCamelCase, modifiers.FormatValues(strings.ToLower)},
		},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			// If modifiers are not applied, the percolator should not match
			percolator, err := NewStructPercolator(defaultparser.Parse, test.input)
			require.NoError(t, err)

			result := percolator.Percolate(&test.document)
			require.False(t, result)

			// If modifiers are applied, the percolator should match
			percolator, err = NewStructPercolator(defaultparser.Parse, test.input, test.modifiers...)
			require.NoError(t, err)

			result = percolator.Percolate(&test.document)
			require.True(t, result)
		})
	}
}
