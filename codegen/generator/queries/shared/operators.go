package shared

import (
	"go/types"

	. "github.com/dave/jennifer/jen"

	"github.com/mchenriques22/pergolator/codegen/generator/queries/shared/jen"
)

func GenerateOperatorCases(typeName *types.TypeName, fieldName string, isPointer bool, inner func(*types.TypeName, string, string, bool) *Statement) *Statement {
	return Add(
		Case(Id("tree").Dot("Eq")).Block(inner(typeName, fieldName, "==", isPointer)),
		Line().Case(Id("tree").Dot("Gte")).Block(inner(typeName, fieldName, ">=", isPointer)),
		Line().Case(Id("tree").Dot("Gt")).Block(inner(typeName, fieldName, ">", isPointer)),
		Line().Case(Id("tree").Dot("Lte")).Block(inner(typeName, fieldName, "<=", isPointer)),
		Line().Case(Id("tree").Dot("Lt")).Block(inner(typeName, fieldName, "<", isPointer)),
	)
}

//	return func(document *Struct) bool {
//		return document.BasicInt64 <= parsed
//	}
func GenerateBasicCmpStatement(typeName *types.TypeName, fieldName string, sign string, isPointer bool) *Statement {
	op := Empty()
	if isPointer {
		op = Id("document").Dot(fieldName).Op("!=").Nil().Op("&&").Op("*")
	}

	return Return().Func().Params(Id("document").Op("*").Add(jen.Qual(typeName))).Bool().Block(
		Return(op.Id("document").Dot(fieldName).Op(sign).Id("parsed")),
	)
}
