//go:generate go run github.com/mchenriques22/pergolator github.com/mchenriques22/pergolator/tests/tags.MyStruct
package tags

type MyStruct struct {
	FieldA MyStructB `json:"FieldA,!flatten"`
}

type MyStructB struct {
	FieldB string
}
