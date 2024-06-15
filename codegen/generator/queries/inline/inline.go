package inline

import (
	"go/types"

	. "github.com/dave/jennifer/jen"

	"github.com/antoninferrand/pergolator/codegen/generator/queries/shared"
	"github.com/antoninferrand/pergolator/codegen/generator/queries/shared/jen"
)

/*
func InlineValueQuery(query *tree.Query) func (*InlineValue) bool {
	key, suffix, _ := strings.Cut(query.Key, ".")
	_ = suffix
	switch key {
	case "":
		parsed := InlineValue(query.Value)
		return func(document *InlineValue) bool {
			return *document == parsed
		}
	default:
		return func(document *InlineValue) bool {
			return false
		}
	}
}
*/

func GenerateBasicQueryStatements(currentTypeName *types.TypeName, t *types.Basic) (*Statement, *Statement) {
	out, err := shared.ParseBasicType(currentTypeName, t)
	if err != nil {
		return shared.ReturnFalseQuery(currentTypeName), shared.ReturnFalseQuery(currentTypeName)
	}

	return Case(Lit("")).Block(
			out,
			Id("converted").Op(":=").Add(jen.Qual(currentTypeName)).Call(Id("parsed")),
			Return(Func().Params(Id("document").Op("*").Add(jen.Qual(currentTypeName))).Bool().Block(
				Return(Op("*").Id("document").Op("==").Id("converted"))),
			),
		),
		shared.ReturnFalseQuery(currentTypeName)
}
