package jen

import (
	"go/types"
	"log/slog"

	"github.com/dave/jennifer/jen"
)

var (
	destPackage string
)

func SetDestPackage(pkg string) {
	destPackage = pkg
}

func Qual(typeName *types.TypeName) *jen.Statement {
	if typeName.Pkg().Name() == destPackage {
		return jen.Id(typeName.Name())
	}

	slog.Debug("qualifying type", slog.String("type", typeName.String()), slog.String("pkg", typeName.Pkg().Name()), slog.String("dest_package", destPackage))

	return jen.Qual(typeName.Pkg().Path(), typeName.Name())
}
