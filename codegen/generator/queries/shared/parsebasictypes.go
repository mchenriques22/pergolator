package shared

import (
	"fmt"
	"go/types"
	"strconv"
	"strings"

	. "github.com/dave/jennifer/jen"
)

func ParseBasicType(typeName *types.TypeName, t *types.Basic) (*Statement, error) {
	switch t.Kind() {
	case types.String:
		return Id("parsed").Op(":=").Id("query").Op(".").Id("Value"), nil
	case types.Int:
		parsingStatement := List(Id("tmp"), Err()).Op(":=").Qual("strconv", "ParseInt").Call(Id("query").Op(".").Id("Value"), Lit(10), Lit(64))
		return parsingStatement.Add(Line(), IfErrNotNilStatement(typeName)).Line().Id("parsed").Op(":=").Op(t.String()).Call(Id("tmp")), nil
	case types.Int8, types.Int16, types.Int32:
		bitSize, err := strconv.Atoi(strings.TrimPrefix(t.String(), "int"))
		if err != nil {
			return nil, fmt.Errorf("failed to parse bit size: %w", err)
		}
		parsingStatement := List(Id("tmp"), Err()).Op(":=").Qual("strconv", "ParseInt").Call(Id("query").Op(".").Id("Value"), Lit(10), Lit(bitSize))
		return parsingStatement.Add(Line(), IfErrNotNilStatement(typeName)).Line().Id("parsed").Op(":=").Op(t.String()).Call(Id("tmp")), nil
	case types.Int64:
		parsingStatement := List(Id("parsed"), Err()).Op(":=").Qual("strconv", "ParseInt").Call(Id("query").Op(".").Id("Value"), Lit(10), Lit(64))
		return parsingStatement.Add(Line(), IfErrNotNilStatement(typeName)), nil
	case types.Uint:
		parsingStatement := List(Id("tmp"), Err()).Op(":=").Qual("strconv", "ParseUint").Call(Id("query").Op(".").Id("Value"), Lit(10), Lit(64))
		return parsingStatement.Add(Line(), IfErrNotNilStatement(typeName)).Line().Id("parsed").Op(":=").Op(t.String()).Call(Id("tmp")), nil
	case types.Uint8, types.Uint16, types.Uint32:
		bitSize, err := strconv.Atoi(strings.TrimPrefix(t.String(), "uint"))
		if err != nil {
			return nil, fmt.Errorf("failed to parse bit size: %w", err)
		}
		parsingStatement := List(Id("tmp"), Err()).Op(":=").Qual("strconv", "ParseUint").Call(Id("query").Op(".").Id("Value"), Lit(10), Lit(bitSize))
		return parsingStatement.Add(Line(), IfErrNotNilStatement(typeName)).Line().Id("parsed").Op(":=").Op(t.String()).Call(Id("tmp")), nil
	case types.Uint64:
		parsingStatement := List(Id("parsed"), Err()).Op(":=").Qual("strconv", "ParseUint").Call(Id("query").Op(".").Id("Value"), Lit(10), Lit(64))
		return parsingStatement.Add(Line(), IfErrNotNilStatement(typeName)), nil
	case types.Float32:
		parsingStatement := List(Id("tmp"), Err()).Op(":=").Qual("strconv", "ParseFloat").Call(Id("query").Op(".").Id("Value"), Lit(32))
		return parsingStatement.Add(Line(), IfErrNotNilStatement(typeName)).Line().Id("parsed").Op(":=").Op(t.String()).Call(Id("tmp")), nil
	case types.Float64:
		parsingStatement := List(Id("parsed"), Err()).Op(":=").Qual("strconv", "ParseFloat").Call(Id("query").Op(".").Id("Value"), Lit(64))
		return parsingStatement.Add(Line(), IfErrNotNilStatement(typeName)), nil
	case types.Bool:
		parsingStatement := List(Id("parsed"), Err()).Op(":=").Qual("strconv", "ParseBool").Call(Id("query").Op(".").Id("Value"))
		return parsingStatement.Add(Line(), IfErrNotNilStatement(typeName)), nil
	default:
		return nil, fmt.Errorf("basic type not handled: %v", t.String())
	}
}

func ReturnFalseQuery(typeName *types.TypeName) *Statement {
	return Return(GetFalseFn(typeName))
}

func IfErrNotNilStatement(typeName *types.TypeName) *Statement {
	return If(Err().Op("!=").Nil()).Block(
		ReturnFalseQuery(typeName),
	)
}
