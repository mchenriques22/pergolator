package _map

import (
	"go/types"
	"log/slog"

	. "github.com/dave/jennifer/jen"

	"github.com/mchenriques22/pergolator/codegen/generator/queries/shared"
	"github.com/mchenriques22/pergolator/codegen/generator/queries/shared/jen"
)

func GenerateMapFieldCase(typeName *types.TypeName, fieldName string, fieldType *types.Map) *Statement {
	switch el := fieldType.Elem().(type) {
	case *types.Basic:
		out, err := shared.ParseBasicType(typeName, el)
		if err != nil {
			return Comment("failed to generate case for " + fieldName + ": " + err.Error()).Add(shared.ReturnFalseQuery(typeName))
		}

		/*
			value, found := document.SimpleMap[suffix]
			if !found {
				return false
			}
			return value == parsed
		*/
		return out.Line().Return().Func().Params(Id("document").Op("*").Add(jen.Qual(typeName))).Bool().Block(
			List(Id("value"), Id("found")).Op(":=").Id("document").Dot(fieldName).Index(Id("suffix")),
			Return(Id("found").Op("&&").Id("value").Op("==").Id("parsed")),
		)
	case *types.Named:
		/*
			mapKey, mapSuffix, found := strings.Cut(suffix, ".")
			if !found {
				return func(document *Struct) bool {
					return false
				}
			}
			fn := NestedQuery(&tree.Query{
				Key:   mapSuffix,
				Value: query.Value,
			})
			return func(document *Struct) bool {
				value, mapFound := document.MapStringNested[mapKey]
				return mapFound && fn(&value)
			}
		*/
		return shared.ReturnFalseQuery(typeName)
	default:
		slog.Warn("failed to generate query_map case", slog.String("type", el.String()), slog.String("underlying", el.Underlying().String()))
		return shared.ReturnFalseQuery(typeName)
	}
}
