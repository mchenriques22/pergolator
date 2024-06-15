package utils

import (
	"fmt"
	"go/types"
	"log/slog"
	"strings"

	"golang.org/x/tools/go/packages"
)

// Parse parses a source type path (examples: "github.com/antoninferrand/pergolator/tests/basic.Struct", "http/http.Request") and returns a types.TypeName.
func Parse(sourceTypePath string) (*types.TypeName, error) {
	idx := strings.LastIndexByte(sourceTypePath, '.')
	if idx == -1 {
		return nil, fmt.Errorf(`expected qualified type as "pkg/path.MyType, got: %s`, sourceTypePath)
	}

	// TODO: Determine the package name
	jdx := strings.LastIndexByte(sourceTypePath[:idx], '/')

	return types.NewTypeName(0, types.NewPackage(sourceTypePath[:idx], sourceTypePath[jdx+1:idx]), sourceTypePath[idx+1:], nil), nil
}

// ParseAll parses all source type paths (examples: "github.com/antoninferrand/pergolator/tests/basic.Struct", "http/http.Request") and returns a slice of types.TypeName.
func ParseAll(sourceTypePaths ...string) ([]*types.TypeName, error) {
	sourceTypes := make([]*types.TypeName, 0, len(sourceTypePaths))
	for _, sourceTypePath := range sourceTypePaths {
		sourceType, err := Parse(sourceTypePath)
		if err != nil {
			return nil, fmt.Errorf("failed to parse source type: %v", err)
		}
		sourceTypes = append(sourceTypes, sourceType)
	}

	return sourceTypes, nil
}

// GetStructDefinition retrieves the struct definition for a given type name.
// For that it needs to load the package.
func GetStructDefinition(typeName *types.TypeName) (types.Type, error) {
	// 2. Inspect package and use type checker to infer imported types
	pkg, err := loadPackage(typeName.Pkg().Path())
	if err != nil {
		return nil, err
	}

	// 3. Lookup the given source type name in the package declarations
	obj := pkg.Types.Scope().Lookup(typeName.Name())
	if obj == nil {
		return nil, fmt.Errorf("%s not found in declared types of %s", typeName.Name(), pkg)
	}

	// 4. We check if it is a declared type
	if _, ok := obj.(*types.TypeName); !ok {
		return nil, fmt.Errorf("%v is not a named type", obj)
	}

	return obj.Type().Underlying(), nil
}

func loadPackage(path string) (*packages.Package, error) {
	cfg := &packages.Config{Mode: packages.NeedTypes | packages.NeedImports}

	loadedPackages, err := packages.Load(cfg, path)
	if err != nil {
		return nil, fmt.Errorf("failed to load whole package for inspection: %v", err)
	}

	for _, loadedPackage := range loadedPackages {
		for _, pkgErr := range loadedPackage.Errors {
			slog.Warn("errors while loading package for inspection. This can be due to fields being renamed or removed from a struct. In some cases it could cause partially generated percolators, please double check the code generated.", slog.String("error", pkgErr.Error()))
		}
	}

	if len(loadedPackages) == 0 {
		return nil, fmt.Errorf("no packages found in %s", path)
	}

	return loadedPackages[0], nil
}
