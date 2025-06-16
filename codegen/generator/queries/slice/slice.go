package slice

import (
	"fmt"
	"go/types"
	"log/slog"

	. "github.com/dave/jennifer/jen"

	"github.com/mchenriques22/pergolator/codegen/generator/queries/named"
	"github.com/mchenriques22/pergolator/codegen/generator/queries/shared"
	"github.com/mchenriques22/pergolator/codegen/generator/queries/shared/jen"
	"github.com/mchenriques22/pergolator/codegen/utils"
)

func GenerateSliceFieldCase(typeName *types.TypeName, fieldName string, t *types.Slice) *Statement {
	switch el := t.Elem().(type) {
	case *types.Basic:
		out, err := shared.ParseBasicType(typeName, el)
		if err != nil {
			return Comment(fmt.Sprintf("failed to generate case for %s: %s", fieldName, err.Error())).Add(shared.ReturnFalseQuery(typeName))
		}

		if el.Kind() == types.Bool {
			return out.Line().Add(generateSliceCmpStatement(typeName, fieldName, "==", false))
		}

		return out.Line().Switch(Id("query").Dot("Sign")).Block(
			shared.GenerateOperatorCases(typeName, fieldName, false, generateSliceCmpStatement),
			Default().Block(shared.ReturnFalseQuery(typeName)),
		)
	case *types.Pointer:
		switch eld := el.Elem().(type) {
		case *types.Basic:
			out, err := shared.ParseBasicType(typeName, eld)
			if err != nil {
				return Comment(fmt.Sprintf("failed to generate case for %s: %s", fieldName, err.Error())).Add(shared.ReturnFalseQuery(typeName))
			}

			if eld.Kind() == types.Bool {
				return out.Line().Add(generateSliceCmpStatement(typeName, fieldName, "==", true))
			}

			return out.Line().Switch(Id("query").Dot("Sign")).Block(
				shared.GenerateOperatorCases(typeName, fieldName, true, generateSliceCmpStatement),
				Default().Block(shared.ReturnFalseQuery(typeName)),
			)
		case *types.Named:
			destinationType, err := utils.Parse(eld.String())
			if err != nil {
				slog.Warn("failed to generate query_slice case", slog.String("type", eld.String()), slog.String("underlying", eld.Underlying().String()))
				return shared.ReturnFalseQuery(typeName)
			}
			return named.GenerateNamedFieldCaseFromSlice(typeName, fieldName, destinationType, true)
		default:
			slog.Warn("failed to generate query_slice case", slog.String("type", eld.String()), slog.String("underlying", eld.Underlying().String()))
			return shared.ReturnFalseQuery(typeName)
		}

	case *types.Named:
		destinationType, err := utils.Parse(el.String())
		if err != nil {
			slog.Warn("failed to generate query_slice case", slog.String("type", el.String()), slog.String("underlying", el.Underlying().String()))
			return shared.ReturnFalseQuery(typeName)
		}
		return named.GenerateNamedFieldCaseFromSlice(typeName, fieldName, destinationType, false)

	default:
		slog.Warn("failed to generate case", slog.String("type", t.Elem().String()), slog.String("underlying", t.Elem().Underlying().String()))
		return shared.ReturnFalseQuery(typeName)
	}
}

func generateSliceCmpStatement(typeName *types.TypeName, fieldName string, sign string, isPointer bool) *Statement {
	op := Empty()
	if isPointer {
		op = Id("value").Op("!=").Nil().Op("&&").Op("*")
	}

	return Return().Func().Params(Id("document").Op("*").Add(jen.Qual(typeName))).Bool().Block(
		For(List(Id("_"), Id("value")).Op(":=").Range().Id("document").Dot(fieldName)).Block(
			If(op.Id("value").Op(sign).Id("parsed")).Block(
				Return(True()),
			),
		).Line().Return(False()),
	)
}
