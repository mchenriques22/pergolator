package shared

import (
	"go/types"

	. "github.com/dave/jennifer/jen"

	"github.com/mchenriques22/pergolator/codegen/generator/queries/shared/jen"
)

func GenerateFalseFn(typeName *types.TypeName) *Statement {
	return Func().Id("p" + typeName.Name() + "FalseFn").Params(Op("_").Op("*").Add(jen.Qual(typeName))).Bool().Block(
		Return(Lit(false)),
	)
}

func GetFalseFn(typeName *types.TypeName) *Statement {
	return Id("p" + typeName.Name() + "FalseFn")
}
