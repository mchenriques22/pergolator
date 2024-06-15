package slice

import (
	"fmt"
	"testing"

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
			input:    "slice_string:my-name",
			document: Struct{SliceString: []string{"my-name"}},
			expected: true,
		},
		{
			input:    "basic_string:my-name",
			document: Struct{SliceString: []string{"not-my-name"}},
			expected: false,
		},
		{
			input:    "slice_int16:1",
			document: Struct{SliceInt16: []int16{1}},
			expected: true,
		},
		{
			input:    "slice_int16:0",
			document: Struct{SliceInt16: []int16{1}},
			expected: false,
		},
		{
			input:    "slice_nested.value:my-name",
			document: Struct{SliceNested: []NestedStruct{{Value: "my-name"}}},
			expected: true,
		},
		{
			input:    "slice_nested.value:not-my-name",
			document: Struct{SliceNested: []NestedStruct{{Value: "my-name"}}},
			expected: false,
		},
		{
			input:    "slice_nested.nested_value.field:my-name",
			document: Struct{SliceNested: []NestedStruct{{NestedValue: DeeplyNestedStruct{Field: "my-name"}}}},
			expected: true,
		},
		{
			input:    "slice_nested.nested_value.field:not-my-name",
			document: Struct{SliceNested: []NestedStruct{{NestedValue: DeeplyNestedStruct{Field: "my-name"}}}},
			expected: false,
		},
		{
			input:    "slice_pointer_basic:my-name",
			document: Struct{SlicePointerBasic: []*string{toPointer("my-name")}},
			expected: true,
		},
		{
			input:    "slice_pointer_basic:not-my-name",
			document: Struct{SlicePointerBasic: []*string{toPointer("my-name")}},
			expected: false,
		},
		{
			input:    "slice_pointer_basic:ok",
			document: Struct{SlicePointerBasic: []*string{nil}},
			expected: false,
		},
		{
			input:    "slice_pointer_int16:1",
			document: Struct{SlicePointerInt16: []*int16{toInt16(1)}},
			expected: true,
		},
		{
			input:    "slice_pointer_int16:32768",
			document: Struct{SlicePointerInt16: []*int16{toInt16(32767)}},
			expected: false,
		},
		{
			input:    "slice_pointer_nested.nested_value.field:my-name",
			document: Struct{SlicePointerNested: []*NestedStruct{{NestedValue: DeeplyNestedStruct{Field: "my-name"}}}},
			expected: true,
		},
		{
			input:    "slice_pointer_nested.nested_value.field:not-my-name",
			document: Struct{SlicePointerNested: []*NestedStruct{{NestedValue: DeeplyNestedStruct{Field: "my-name"}}}},
			expected: false,
		},
		{
			input:    "slice_bool:true",
			document: Struct{SliceBool: []bool{false, false, false, true}},
			expected: true,
		},
		{
			input:    "slice_bool:false",
			document: Struct{SliceBool: []bool{true, true, true, false}},
			expected: true,
		},
		{
			input:    "slice_bool:true",
			document: Struct{SliceBool: []bool{false, false, false, false}},
			expected: false,
		},
		{
			input:    "slice_bool:false",
			document: Struct{SliceBool: []bool{true, true, true, true}},
			expected: false,
		},
		{
			input:    "slice_int16:>10",
			document: Struct{SliceInt16: []int16{1, 8, 13}},
			expected: true,
		},
		{
			input:    "slice_int16:<10",
			document: Struct{SliceInt16: []int16{10, 11, 13}},
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

func toPointer(s string) *string {
	return &s
}

func toInt16(i int16) *int16 {
	return &i
}

func BenchmarkPercolate(b *testing.B) {
	tests := []int{2, 50, 100, 1000}
	percolator, err := NewStructPercolator(defaultparser.Parse, "slice_string:noop")
	require.NoError(b, err)
	for _, tt := range tests {
		b.Run(fmt.Sprintf("len:%d", tt), func(b *testing.B) {
			doc := Struct{SliceString: getBenchmarkSlice(tt)}
			for b.Loop() {
				percolator.Percolate(&doc)
			}
		})
	}
}

func getBenchmarkSlice(n int) (out []string) {
	for range n {
		out = append(out, "item")
	}
	return out
}
