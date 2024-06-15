package utils

import (
	"fmt"
	"go/types"
	"log/slog"
	"sort"
	"strings"
)

var (
	knownTypes = map[string]struct{}{}
)

func getChildTypes(typeName *types.TypeName, depth int64, maxDepth int64) ([]*types.TypeName, error) {
	if depth >= maxDepth {
		return nil, nil
	}

	if _, ok := knownTypes[typeName.String()]; ok {
		return nil, nil
	}
	knownTypes[typeName.String()] = struct{}{}

	structDefinitionType, err := GetStructDefinition(typeName)
	if err != nil {
		return nil, fmt.Errorf("failed to get struct type: %v", err)
	}

	structType, ok := structDefinitionType.(*types.Struct)
	if !ok {
		slog.Debug("Skipping non-struct type", slog.String("sourceTypeName", typeName.Name()), slog.String("type", structDefinitionType.String()), slog.String("underlyingType", structDefinitionType.Underlying().String()))
		return nil, nil
	}

	var out []*types.TypeName
	for i := 0; i < structType.NumFields(); i++ {
		field := structType.Field(i)
		if !strings.Contains(field.Type().String(), ".") || !field.Exported() {
			continue
		}

		// TODO: Support anonymous fields
		if field.Anonymous() {
			continue
		}

		var childType *types.TypeName
		switch u := field.Type().(type) {
		case *types.Pointer:
			childType = getUnderlyingType(u.Elem())
		case *types.Slice:
			childType = getUnderlyingType(u.Elem())
		case *types.Map:
			childType = getUnderlyingType(u.Elem())
		case *types.Named:
			childType = u.Obj()

		default:
			continue
		}

		if childType == nil {
			continue
		}

		// TODO: Maybe we can do better than that
		if !childType.Exported() {
			continue
		}

		out = append(out, childType)

		childFieldTypes, err := getChildTypes(childType, depth+1, maxDepth)
		if err != nil {
			slog.Warn("failed to get child types", slog.String("error", err.Error()))
		}

		out = append(out, childFieldTypes...)

	}

	return out, nil
}

func GetAllTypes(sourceTypes []*types.TypeName, maxDepth int64) ([]*types.TypeName, error) {
	dynamicSourceTypes := make(map[string]map[string]*types.TypeName)
	for _, sourceType := range sourceTypes {
		_, ok := dynamicSourceTypes[sourceType.Pkg().Path()]
		if !ok {
			dynamicSourceTypes[sourceType.Pkg().Path()] = make(map[string]*types.TypeName)
		}

		dynamicSourceTypes[sourceType.Pkg().Path()][sourceType.Name()] = sourceType
		childTypes, err := getChildTypes(sourceType, 0, maxDepth)
		if err != nil {
			return nil, fmt.Errorf("failed to get all types: %v", err)
		}
		for _, childType := range childTypes {
			dynamicSourceTypes[sourceType.Pkg().Path()][childType.Name()] = childType
		}
	}

	sortedSourceTypes := make([]*types.TypeName, 0, len(dynamicSourceTypes))
	for _, dynamicSourceType := range dynamicSourceTypes {
		for _, sourceType := range dynamicSourceType {
			sortedSourceTypes = append(sortedSourceTypes, sourceType)
		}
	}
	sort.SliceStable(sortedSourceTypes, func(i, j int) bool {
		return sortedSourceTypes[i].String() < sortedSourceTypes[j].String()
	})

	return sortedSourceTypes, nil
}

func getUnderlyingType(t types.Type) *types.TypeName {
	switch u := t.(type) {
	case *types.Named:
		return u.Obj()
	case *types.Pointer:
		return getUnderlyingType(u.Elem())
	case *types.Slice:
		return getUnderlyingType(u.Elem())
	case *types.Map:
		return getUnderlyingType(u.Elem())
	default:
		return nil
	}
}
