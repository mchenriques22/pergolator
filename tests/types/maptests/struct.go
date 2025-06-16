//go:generate go run github.com/mchenriques22/pergolator github.com/mchenriques22/pergolator/tests/types/maptests.Struct
package maptests

type Struct struct {
	MapStringString  map[string]string  `json:"map_string_string"`  // map of string to string
	MapStringFloat32 map[string]float32 `json:"map_string_float32"` // map of string to float32
	MapStringNested  map[string]Nested  `json:"map_string_nested"`  // map of string to nested struct
}

type Nested struct {
	Value string `json:"value"`
}
