package special

import (
	"go/types"

	"github.com/dave/jennifer/jen"

	"github.com/mchenriques22/pergolator/codegen/utils"
)

var KnownTypes = map[string]func(typeName *types.TypeName, optionsGetter utils.FieldOptionsGetter) *jen.Statement{
	"time.Time": GenerateTimeQuery,
}
