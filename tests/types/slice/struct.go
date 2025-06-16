//go:generate go run github.com/mchenriques22/pergolator github.com/mchenriques22/pergolator/tests/types/slice.Struct
package slice

type Struct struct {
	SliceString        []string        `json:"slice_string"`         // slice of strings
	SliceInt16         []int16         `json:"slice_int16"`          // slice of int16
	SliceBool          []bool          `json:"slice_bool"`           // slice of bool
	SliceNested        []NestedStruct  `json:"slice_nested"`         // slice of nested struct
	SlicePointerBool   []*bool         `json:"slice_pointer_bool"`   // slice of pointer to bool
	SlicePointerBasic  []*string       `json:"slice_pointer_basic"`  // slice of pointer to string
	SlicePointerInt16  []*int16        `json:"slice_pointer_int16"`  // slice of pointer to int16
	SlicePointerNested []*NestedStruct `json:"slice_pointer_nested"` // slice of pointer to nested struct
}

type DeeplyNestedStruct struct {
	Field string `json:"field"`
}

type NestedStruct struct {
	Value       string             `json:"value"`
	NestedValue DeeplyNestedStruct `json:"nested_value"`
}
