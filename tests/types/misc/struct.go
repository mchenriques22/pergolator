//go:generate go run github.com/mchenriques22/pergolator github.com/mchenriques22/pergolator/tests/types/misc.Struct
package misc

import "time"

type Struct struct {
	Ignore        string        `json:"-"`              // ignored field
	Nested        NestedStruct  `json:"nested"`         // nested struct
	InlineString  InlineString  `json:"inline_string"`  // inline_string value
	InlineUint64  InlineUint64  `json:"inline_uint64"`  // inline_uint64 value
	PointerString *string       `json:"pointer_string"` // pointer to string
	PointerNested *NestedStruct `json:"pointer_nested"` // pointer to nested struct
	PointerBool   *bool         `json:"pointer_bool"`   // pointer to bool

	FlattenNested       NestedStruct       `json:"flatten_nested,!flatten"`
	FlattenDeeplyNested DeeplyNestedStruct `json:"flatten_deeply_nested,!flatten"`

	Time time.Time `json:"time"`
}

type DeeplyNestedStruct struct {
	Field string `json:"field"`
}

type NestedStruct struct {
	Value            string             `json:"value"`
	NestedValue      DeeplyNestedStruct `json:"nested_value"`
	FlattenRecStruct FlattenRecStruct   `json:"flatten_rec_struct,!flatten"`
}

type InlineString string

type InlineUint64 uint64

type FlattenRecStruct struct {
	AnotherField string `json:"another_field"`
}
