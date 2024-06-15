package queries

import (
	"go/types"
	"log/slog"

	. "github.com/dave/jennifer/jen"

	"github.com/antoninferrand/pergolator/codegen/generator/queries/basic"
	"github.com/antoninferrand/pergolator/codegen/generator/queries/inline"
	"github.com/antoninferrand/pergolator/codegen/generator/queries/map"
	"github.com/antoninferrand/pergolator/codegen/generator/queries/named"
	"github.com/antoninferrand/pergolator/codegen/generator/queries/pointer"
	"github.com/antoninferrand/pergolator/codegen/generator/queries/shared"
	"github.com/antoninferrand/pergolator/codegen/generator/queries/shared/jen"
	"github.com/antoninferrand/pergolator/codegen/generator/queries/slice"
	"github.com/antoninferrand/pergolator/codegen/generator/queries/special"
	"github.com/antoninferrand/pergolator/codegen/utils"
)

const (
	// skipTag is used to skip a field in the query generation. As a result, the field cannot be used in a queries.
	// This is the same value as the one used in json tags to skip a field during JSON marshalling.
	// We can consider making it configurable in the future.
	skipTag = "-"
)

/*
func NestedStructQuery(query *tree.Query) func (*NestedStruct) bool{
	key, suffix, _ := strings.Cut(query.Key, ".")
	_ = suffix
	switch key {
	case "value":
		parsed := query.Value
		return func(document *NestedStruct) bool {
			return document.Value == parsed
		}
	case "nested_value":
		fn := DeeplyNestedStructQuery(&tree.Query{
			Key:   suffix,
			Value: query.Value,
		})
		return func(nestedStruct *NestedStruct) bool {
			return fn(&nestedStruct.NestedValue)
		}
	default:
		return func(document *NestedStruct) bool {
			return false
		}
	}
}
*/

func GenerateQueryStatements(typeName *types.TypeName, _type types.Type, optionsGetter utils.FieldOptionsGetter) *Statement {
	caseStmts, defaultStmt := Empty(), shared.ReturnFalseQuery(typeName)
	if typeName.Type() != nil {
		genFn, ok := special.KnownTypes[typeName.Type().String()]
		if ok {
			return Line().Func().Id("p" + typeName.Name() + "Query").Params(Id("query").Op("*").Qual("github.com/antoninferrand/pergolator/tree", "Query")).Func().Params(Op("*").Add(jen.Qual(typeName))).Bool().Block(genFn(typeName, optionsGetter)).
				Line().Line().Add(shared.GenerateFalseFn(typeName)).Line()
		}
	}

	switch typeCategory := _type.(type) {
	case *types.Struct:
		caseStmts, defaultStmt = generateStructQueryCases(typeName, optionsGetter, typeCategory)
	case *types.Basic:
		caseStmts, defaultStmt = inline.GenerateBasicQueryStatements(typeName, typeCategory)
	default:
		slog.Warn("unsupported type", slog.String("type", typeCategory.String()), slog.String("name", typeName.Name()))
	}

	return queryStatementWrapper(typeName, caseStmts, defaultStmt)
}

func queryStatementWrapper(typeName *types.TypeName, caseStmts, defaultStmt *Statement) *Statement {
	return Line().Func().Id("p"+typeName.Name()+"Query").Params(Id("query").Op("*").Qual("github.com/antoninferrand/pergolator/tree", "Query")).Func().Params(Op("*").Add(jen.Qual(typeName))).Bool().Block(
		Id("key").Op(",").Id("suffix").Op(",").Id("_").Op(":=").Qual("strings", "Cut").Call(Id("query").Dot("Key"), Lit(".")),
		Id("_").Op("=").Id("suffix"),
		Switch(Id("key")).Block(
			caseStmts,
			Default().Block(
				defaultStmt,
			),
		),
	).Line().Line().Add(shared.GenerateFalseFn(typeName)).Line()
}

func generateStructQueryCases(typeName *types.TypeName, optionsGetter utils.FieldOptionsGetter, t *types.Struct) (caseStmts *Statement, defaultStmt *Statement) {
	var cases, flattenedStmts, flattenedReturnStmt []*Statement
	for i := 0; i < t.NumFields(); i++ {
		field := t.Field(i)
		options := optionsGetter(typeName, field, t.Tag(i))
		if shouldSkipField(field, options) {
			continue
		}

		switch typedField := field.Type().(type) {
		case *types.Named:
			if options.Flatten {
				parsed, err := utils.Parse(typedField.String())
				if err != nil {
					continue
				}

				flattenedStmts = append(flattenedStmts, Id("fn"+field.Name()).Op(":=").Id("p"+parsed.Name()+"Query").Call(Id("query")))
				flattenedStmts = append(flattenedStmts, Id("fn"+field.Name()+"Enabled").Op(":=").Qual("reflect", "ValueOf").Call(Id("fn"+field.Name())).Dot("Pointer").Call().Op("!=").Qual("reflect", "ValueOf").Call(shared.GetFalseFn(parsed)).Dot("Pointer").Call())
				flattenedReturnStmt = append(flattenedReturnStmt, Id("fn"+field.Name()+"Enabled").Op("&&").Id("fn"+field.Name()).Call(Op("&").Id("document").Dot(field.Name())))
				continue
			}
		}

		cases = append(cases, Case(List(jen.Wrap(options.Aliases, Lit)...)).Block(
			generateStructQueryFieldCase(typeName, field.Name(), field.Type()),
		))
	}

	caseStmts = jen.Join(cases, Line())
	if len(flattenedStmts) == 0 {
		return caseStmts, Return(shared.GetFalseFn(typeName))
	}
	return caseStmts, jen.Join([]*Statement{
		jen.Join(flattenedStmts, Line()),
		Return().Func().Params(Id("document").Op("*").Add(jen.Qual(typeName))).Bool().Block(
			Return(jen.Join(flattenedReturnStmt, Op("||"))),
		),
	}, Line())
}

func generateStructQueryFieldCase(typeName *types.TypeName, fieldName string, fieldType types.Type) *Statement {
	switch t := fieldType.(type) {
	case *types.Basic:
		return basic.GenerateBasicFieldCase(typeName, fieldName, t)
	case *types.Slice:
		return slice.GenerateSliceFieldCase(typeName, fieldName, t)
	case *types.Pointer:
		return pointer.GeneratePointerFieldCase(typeName, fieldName, t)
	case *types.Named:
		return named.GenerateNamedFieldCase(typeName, fieldName, t)
	case *types.Map:
		return _map.GenerateMapFieldCase(typeName, fieldName, t)
	default:
		slog.Info("field type not handled", slog.String("type", fieldType.String()), slog.String("underlying", t.Underlying().String()), slog.String("field_name", fieldName))
		return shared.ReturnFalseQuery(typeName)
	}
}

func shouldSkipField(field *types.Var, options utils.FieldOptions) bool {
	return !field.Exported() || len(options.Aliases) == 1 && options.Aliases[0] == skipTag || len(options.Aliases) == 0
}
