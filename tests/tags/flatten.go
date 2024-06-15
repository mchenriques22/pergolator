//go:generate go run github.com/antoninferrand/pergolator github.com/antoninferrand/pergolator/tests/tags.MyStruct
package tags

type MyStruct struct {
	FieldA MyStructB `json:"FieldA,!flatten"`
}

type MyStructB struct {
	FieldB string
}
