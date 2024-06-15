package codegen

import (
	"bytes"
	"io"
	"log/slog"

	. "github.com/dave/jennifer/jen"
	"golang.org/x/tools/imports"

	"github.com/antoninferrand/pergolator/codegen/generator/percolator"
	"github.com/antoninferrand/pergolator/codegen/generator/queries"
	"github.com/antoninferrand/pergolator/codegen/generator/queries/shared/jen"
	"github.com/antoninferrand/pergolator/codegen/utils"
)

// Run generates the percolator and the queries used by the percolator.
func Run(w io.Writer, goPackage string, descriptorPath string, sourceTypePaths []string, maxDepth int64, renameFieldsToSnakeCase bool) error {
	sourceTypes, err := utils.ParseAll(sourceTypePaths...)
	if err != nil {
		return err
	}

	recursiveSourceTypes, err := utils.GetAllTypes(sourceTypes, maxDepth)
	if err != nil {
		return err
	}

	tagGetter, err := utils.GetTagGetter(descriptorPath, renameFieldsToSnakeCase)
	if err != nil {
		return err
	}

	// Render the code to create the percolator
	var buffer bytes.Buffer
	err = percolator.Write(&buffer, goPackage, recursiveSourceTypes)
	if err != nil {
		return err
	}

	// Render the code to create the queries used by the percolator
	jen.SetDestPackage(goPackage)
	statement := Empty()
	for _, typeName := range recursiveSourceTypes {
		structDefinitionType, err := utils.GetStructDefinition(typeName)
		if err != nil {
			slog.Warn("failed to get struct type", slog.String("error", err.Error()))
			continue
		}

		statement.Add(queries.GenerateQueryStatements(typeName, structDefinitionType, tagGetter))
	}
	err = statement.Render(&buffer)
	if err != nil {
		return err
	}

	// Sort and deduplicate the imports
	formatted, err := imports.Process("", buffer.Bytes(), nil)
	if err != nil {
		return err
	}

	_, err = w.Write(formatted)
	return err
}
