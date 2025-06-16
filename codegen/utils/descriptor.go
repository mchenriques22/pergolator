package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"go/types"
	"io"
	"log/slog"
	"os"
	"slices"
	"strings"

	"github.com/fatih/structtag"
	"github.com/iancoleman/strcase"
)

var errTagNotFound = errors.New("tag not found")

type FieldOptions struct {
	Aliases []string
	Flatten bool
}

type FieldOptionsGetter func(typeName *types.TypeName, field *types.Var, unparsedTag string) FieldOptions

func getDescriptor(descriptorPath string) (map[string]map[string]string, error) {
	file, err := os.Open(descriptorPath)
	if err != nil {
		return nil, err
	}

	descriptorBytes, err := io.ReadAll(file)
	if err != nil {
		return nil, errors.Join(err, file.Close())
	}
	err = file.Close()
	if err != nil {
		return nil, err
	}

	var descriptor map[string]map[string]string
	err = json.Unmarshal(descriptorBytes, &descriptor)
	if err != nil {
		return nil, err
	}

	return descriptor, nil
}

func GetTagGetter(descriptorPath string, renameFieldsToSnakeCase bool) (FieldOptionsGetter, error) {
	descriptor, err := getDescriptor(descriptorPath)
	if err != nil {
		descriptor = make(map[string]map[string]string)
	}

	return func(typeName *types.TypeName, field *types.Var, unparsedTag string) FieldOptions {
		defaultFieldName := field.Name()
		if renameFieldsToSnakeCase {
			defaultFieldName = strcase.ToSnake(defaultFieldName)
		}

		options := FieldOptions{}
		mergeTagToOptions(&options, unparsedTag, defaultFieldName)

		overwrites := descriptor[typeName.Name()][field.Name()]
		if overwrites != "" {
			mergeTagToOptions(&options, overwrites, defaultFieldName)
		}

		return options
	}, nil
}

func mergeTagToOptions(options *FieldOptions, unparsedTag string, defaultName string) {
	fieldTag, err := extractTag(unparsedTag)
	if err != nil {
		if !errors.Is(err, errTagNotFound) {
			slog.Debug("failed to extract tag", slog.String("error", err.Error()), slog.String("tag", unparsedTag))
		}
		fieldTag = &structtag.Tag{Name: defaultName}
	}

	aliases := make([]string, 0)
	for _, alias := range fieldTag.Options {
		if alias == "" || strings.ContainsRune(alias, '!') {
			continue
		}
		aliases = append(aliases, alias)
	}

	aliases = append(aliases, fieldTag.Name)

	options.Aliases = aliases
	options.Flatten = slices.Contains(fieldTag.Options, "!flatten")
}

func extractTag(rawTags string) (tag *structtag.Tag, err error) {
	tags, err := structtag.Parse(rawTags)
	if err != nil {
		return nil, fmt.Errorf("failed to parse tags %w", err)
	}
	fieldTag, err := tags.Get("pergolator")
	if err == nil {
		return fieldTag, nil
	}

	return tags.Get("json")
}
