package percolator

import (
	"embed"
	"errors"
	"go/types"
	"io"
	"sort"
	"text/template"
)

var (
	//go:embed templates/*
	files embed.FS
)

type BodyArg struct {
	Type       string
	TypePrefix string
}

type HeaderArgs struct {
	Package string
	Imports []string
}

type TemplateInput struct {
	HeaderArgs HeaderArgs
	BodyArgs   []BodyArg
}

func Write(w io.Writer, pkgName string, typeNames []*types.TypeName) error {
	if len(typeNames) == 0 {
		return errors.New("no types")
	}

	percolatorTemplate, err := template.ParseFS(files, "templates/*.gotmpl")
	if err != nil {
		return err
	}

	imports := getAllExternalImports(pkgName, typeNames)
	templateInput := TemplateInput{
		HeaderArgs: HeaderArgs{
			Package: pkgName,
			Imports: imports,
		},
	}

	for _, typeName := range typeNames {
		prefix := ""
		if typeName.Pkg().Name() != pkgName {
			prefix = typeName.Pkg().Name() + "."
		}

		templateInput.BodyArgs = append(templateInput.BodyArgs, BodyArg{
			Type:       typeName.Name(),
			TypePrefix: prefix,
		})
	}

	return percolatorTemplate.ExecuteTemplate(w, "percolators", templateInput)
}

func getAllExternalImports(pkgName string, typeNames []*types.TypeName) []string {
	imports := make(map[string]struct{})
	for _, typeName := range typeNames {
		if typeName.Pkg().Name() != pkgName {
			imports[typeName.Pkg().Path()] = struct{}{}
		}
	}

	result := make([]string, 0, len(imports))
	for imp := range imports {
		result = append(result, imp)
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i] < result[j]
	})

	return result
}
