package basic

import (
	"go/types"
	"log/slog"

	. "github.com/dave/jennifer/jen"

	"github.com/antoninferrand/pergolator/codegen/generator/queries/shared"
	"github.com/antoninferrand/pergolator/codegen/generator/queries/shared/jen"
)

func GenerateBasicFieldCase(typeName *types.TypeName, fieldName string, t *types.Basic) *Statement {
	out, err := shared.ParseBasicType(typeName, t)
	if err != nil {
		slog.Error("failed to generate basic field case", slog.String("type", t.String()), slog.String("name", typeName.Name()), slog.String("field", fieldName), slog.String("error", err.Error()))
		return shared.ReturnFalseQuery(typeName)
	}

	if t.Kind() == types.Bool || t.Kind() == types.String {
		return out.Line().Return().Func().Params(Id("document").Op("*").Add(jen.Qual(typeName))).Bool().Block(
			Return(Id("document").Dot(fieldName).Op("==").Id("parsed")),
		)
	}

	/*
		switch query.Sign {
		case tree.Eq:
			return func(document *Struct) bool {
				return document.BasicInt64 == parsed
			}
		case tree.Gte:
			return func(document *Struct) bool {
				return document.BasicInt64 >= parsed
			}
		case tree.Gt:
			return func(document *Struct) bool {
				return document.BasicInt64 > parsed
			}
		case tree.Lte:
			return func(document *Struct) bool {
				return document.BasicInt64 <= parsed
			}
		case tree.Lt:
			return func(document *Struct) bool {
				return document.BasicInt64 < parsed
			}
		default:
			return pStructFalseFn
		}
	*/

	return out.Line().Switch(Id("query").Dot("Sign")).Block(
		shared.GenerateOperatorCases(typeName, fieldName, false, shared.GenerateBasicCmpStatement),
		Default().Block(shared.ReturnFalseQuery(typeName)),
	)
}
