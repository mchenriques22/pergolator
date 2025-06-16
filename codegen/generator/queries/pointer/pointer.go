package pointer

import (
	"fmt"
	"go/types"
	"log/slog"

	. "github.com/dave/jennifer/jen"

	"github.com/mchenriques22/pergolator/codegen/generator/queries/named"
	"github.com/mchenriques22/pergolator/codegen/generator/queries/shared"
	"github.com/mchenriques22/pergolator/codegen/utils"
)

func GeneratePointerFieldCase(typeName *types.TypeName, fieldName string, t *types.Pointer) *Statement {
	switch u := t.Elem().(type) {
	case *types.Basic:
		return GenerateBasicPointerFieldCase(typeName, fieldName, u)
	case *types.Named:
		destinationType, err := utils.Parse(u.String())
		if err != nil {
			slog.Warn("failed to generate pointer field case", slog.String("entire_type", t.String()), slog.String("type", fmt.Sprintf("%T", u)), slog.String("name", typeName.Name()), slog.String("error", err.Error()))
			return shared.ReturnFalseQuery(typeName)
		}
		return named.GenerateNamedFieldCaseFromPointer(typeName, fieldName, destinationType)
	default:
		slog.Warn("failed to generate pointer field case", slog.String("entire_type", t.String()), slog.String("type", fmt.Sprintf("%T", u)), slog.String("name", typeName.Name()))
		return shared.ReturnFalseQuery(typeName)
	}
}

func GenerateBasicPointerFieldCase(typeName *types.TypeName, fieldName string, t *types.Basic) *Statement {
	out, err := shared.ParseBasicType(typeName, t)
	if err != nil {
		return Comment(fmt.Sprintf("failed to generate case for %s: %s", fieldName, err.Error()))
	}

	return out.Line().Add(shared.GenerateBasicCmpStatement(typeName, fieldName, "==", true))
}
