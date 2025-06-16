package named

import (
	"go/types"

	. "github.com/dave/jennifer/jen"

	"github.com/mchenriques22/pergolator/codegen/generator/queries/shared"
	"github.com/mchenriques22/pergolator/codegen/generator/queries/shared/jen"
	"github.com/mchenriques22/pergolator/codegen/utils"
)

func GenerateNamedFieldCase(typeName *types.TypeName, fieldName string, fieldType *types.Named) *Statement {
	destinationType, err := utils.Parse(fieldType.String())
	if err != nil {
		return shared.ReturnFalseQuery(typeName)
	}
	return createNameQueryWithReturn(typeName, destinationType.Name(), simpleReturnStatement(fieldName, false))
}

func GenerateNamedFieldCaseFromPointer(sourceTypeName *types.TypeName, fieldName string, childType *types.TypeName) *Statement {
	return createNameQueryWithReturn(sourceTypeName, childType.Name(), simpleReturnStatement(fieldName, true))
}

func simpleReturnStatement(fieldName string, isPointer bool) *Statement {
	op := Op("&")
	if isPointer {
		op = Empty()
	}
	return Return(Id("fn").Call(op.Id("document").Dot(fieldName)))
}

func GenerateNamedFieldCaseFromSlice(sourceTypeName *types.TypeName, fieldName string, childType *types.TypeName, isPointer bool) *Statement {
	op := Op("&")
	if isPointer {
		op = Empty()
	}

	sliceReturnStmt := For(List(Id("_"), Id("value")).Op(":=").Range().Id("document").Dot(fieldName)).Block(
		If(Id("fn").Call(op.Id("value")).Block(
			Return(True()),
		))).Line().Return(False())

	return createNameQueryWithReturn(sourceTypeName, childType.Name(), sliceReturnStmt)
}

func createNameQueryWithReturn(sourceTypeName *types.TypeName, underlyingTypeName string, returnStmt *Statement) *Statement {
	return createNamedQueryFn(underlyingTypeName).Add(
		Line(),
		Return(Func().Params(Id("document").Op("*").Add(jen.Qual(sourceTypeName))).Bool().Block(
			returnStmt,
		)),
	)
}

func createNamedQueryFn(underlyingTypeName string) *Statement {
	return Id("fn").Op(":=").Id("p" + underlyingTypeName + "Query").Call(Op("&").Qual("github.com/mchenriques22/pergolator/tree", "Query").Values(Dict{
		Id("Key"):   Id("suffix"),
		Id("Sign"):  Id("query").Dot("Sign"),
		Id("Value"): Id("query").Dot("Value"),
	}))
}
