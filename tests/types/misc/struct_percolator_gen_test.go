package misc

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/antoninferrand/pergolator/tree/defaultparser"
)

func TestPercolate(t *testing.T) {
	tests := []struct {
		input    string
		document Struct
		expected bool
	}{
		{
			input:    "ignored:my-name",
			document: Struct{Ignore: "my-name"},
			expected: false,
		},
		{
			input:    "nested.value:my-name",
			document: Struct{Nested: NestedStruct{Value: "my-name"}},
			expected: true,
		},
		{
			input:    "nested.value:not-my-name",
			document: Struct{Nested: NestedStruct{Value: "my-name"}},
			expected: false,
		},
		{
			input:    "nested.nested_value.field:my-name",
			document: Struct{Nested: NestedStruct{NestedValue: DeeplyNestedStruct{Field: "my-name"}}},
			expected: true,
		},
		{
			input:    "nested.nested_value.field:not-my-name",
			document: Struct{Nested: NestedStruct{NestedValue: DeeplyNestedStruct{Field: "my-name"}}},
			expected: false,
		},
		{
			input:    "inline_string:my-name",
			document: Struct{InlineString: InlineString("my-name")},
			expected: true,
		},
		{
			input:    "inline_string:not-my-name",
			document: Struct{InlineString: InlineString("my-name")},
			expected: false,
		},
		{
			input:    "pointer_string:my-name",
			document: Struct{PointerString: toPointer("my-name")},
			expected: true,
		},
		{
			input:    "pointer_string:not-my-name",
			document: Struct{PointerString: toPointer("my-name")},
			expected: false,
		},
		{
			input:    "pointer_nested.value:my-name",
			document: Struct{PointerNested: &NestedStruct{Value: "my-name"}},
			expected: true,
		},
		{
			input:    "pointer_nested.value:not-my-name",
			document: Struct{PointerNested: &NestedStruct{Value: "my-name"}},
			expected: false,
		},
		{
			input:    "pointer_bool:true",
			document: Struct{PointerBool: toPointer(true)},
			expected: true,
		},
		{
			input:    "pointer_bool:true",
			document: Struct{PointerBool: toPointer(false)},
			expected: false,
		},
		{
			input:    "pointer_bool:invalid",
			document: Struct{PointerBool: toPointer(true)},
			expected: false,
		},
		{
			input:    "pointer_bool:false",
			document: Struct{},
			expected: false,
		},
		{
			input:    "inline_uint64:1",
			document: Struct{InlineUint64: 1},
			expected: true,
		},
		{
			input:    "flatten_nested.value:not-found",
			document: Struct{FlattenNested: NestedStruct{Value: "not-found"}},
			expected: false,
		},
		{
			input:    "value:found",
			document: Struct{FlattenNested: NestedStruct{Value: "found"}},
			expected: true,
		},
		{
			input:    "another_field:found",
			document: Struct{FlattenNested: NestedStruct{FlattenRecStruct: FlattenRecStruct{AnotherField: "found"}}},
			expected: true,
		},
		{
			input:    "time:<\"2021-01-02T00:00:00Z\"",
			document: Struct{Time: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)},
			expected: true,
		},
		{
			input:    "time:>\"2020-12-30T00:00:00Z\"",
			document: Struct{Time: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)},
			expected: true,
		},
		{
			input:    "time:\"2021-01-01T00:00:00Z\"",
			document: Struct{Time: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)},
			expected: true,
		},
		{
			input:    "time:\"2021-01-02T00:00:00Z\"",
			document: Struct{Time: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)},
			expected: false,
		},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			percolator, err := NewStructPercolator(defaultparser.Parse, test.input)
			require.NoError(t, err)

			assert.Equal(t, test.expected, percolator.Percolate(&test.document))
		})
	}
}

func toPointer[T any](t T) *T {
	return &t
}
