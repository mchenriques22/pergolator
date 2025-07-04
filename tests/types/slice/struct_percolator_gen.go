// Code generated by github.com/mchenriques22/pergolator, DO NOT EDIT.
package slice

import (
	"reflect"

	"github.com/mchenriques22/pergolator/tree"
)

type StructPercolator struct {
	fn func(document *Struct) bool
}

// Percolate percolates the document with the percolator's query.
// It will return a boolean indicating if the query matches the document.
func (p *StructPercolator) Percolate(document *Struct) bool {
	return p.fn(document)
}

// NewStructPercolatorConstructor is a constructor of percolators.
// Every percolator it creates will use the parser provided to the constructor.
func NewStructPercolatorConstructor(parseFn tree.ParseFn, modifiers ...tree.Modifiers) func(query string) (*StructPercolator, error) {
	return func(query string) (*StructPercolator, error) {
		return NewStructPercolator(parseFn, query, modifiers...)
	}
}

// NewStructPercolator creates a percolator with a given query and a given parser.
// It returns an error if the parsing failed.
func NewStructPercolator(parseFn tree.ParseFn, query string, modifiers ...tree.Modifiers) (*StructPercolator, error) {
	root, err := parseFn(query)
	if err != nil {
		return nil, err
	}

	for _, modifier := range modifiers {
		root = modifier(root)
		if _, isEmpty := root.(*tree.Empty); isEmpty {
			// If the query is empty, we return a percolator that always returns false
			return &StructPercolator{fn: pStructFalseFn}, nil
		}
	}

	return &StructPercolator{fn: updateNodeStruct(root)}, nil
}

func updateNodeStruct(root tree.Expr) func(document *Struct) bool {
	switch r := root.(type) {
	case *tree.And:
		fns := make([]func(document *Struct) bool, 0, len(r.Children))
		for _, child := range r.Children {
			fn := updateNodeStruct(child)

			// Optimize the case where one of the children is always false
			if reflect.ValueOf(fn).Pointer() == reflect.ValueOf(pStructFalseFn).Pointer() {
				return pStructFalseFn
			}
			fns = append(fns, fn)
		}

		return func(document *Struct) bool {
			for _, fn := range fns {
				if !fn(document) {
					return false
				}
			}
			return true
		}
	case *tree.Or:
		fns := make([]func(document *Struct) bool, 0, len(r.Children))
		for _, child := range r.Children {
			fn := updateNodeStruct(child)

			// Optimize the case where one of the children is always false
			if reflect.ValueOf(fn).Pointer() == reflect.ValueOf(pStructFalseFn).Pointer() {
				continue
			}
			fns = append(fns, fn)
		}

		if len(fns) == 0 {
			return pStructFalseFn
		}

		return func(document *Struct) bool {
			for _, fn := range fns {
				if fn(document) {
					return true
				}
			}
			return false
		}
	case *tree.Not:
		child := updateNodeStruct(r.Child)
		return func(document *Struct) bool {
			return !child(document)
		}
	case *tree.Query:
		return pStructQuery(r)
	}

	return pStructFalseFn
}

type DeeplyNestedStructPercolator struct {
	fn func(document *DeeplyNestedStruct) bool
}

// Percolate percolates the document with the percolator's query.
// It will return a boolean indicating if the query matches the document.
func (p *DeeplyNestedStructPercolator) Percolate(document *DeeplyNestedStruct) bool {
	return p.fn(document)
}

// NewDeeplyNestedStructPercolatorConstructor is a constructor of percolators.
// Every percolator it creates will use the parser provided to the constructor.
func NewDeeplyNestedStructPercolatorConstructor(parseFn tree.ParseFn, modifiers ...tree.Modifiers) func(query string) (*DeeplyNestedStructPercolator, error) {
	return func(query string) (*DeeplyNestedStructPercolator, error) {
		return NewDeeplyNestedStructPercolator(parseFn, query, modifiers...)
	}
}

// NewDeeplyNestedStructPercolator creates a percolator with a given query and a given parser.
// It returns an error if the parsing failed.
func NewDeeplyNestedStructPercolator(parseFn tree.ParseFn, query string, modifiers ...tree.Modifiers) (*DeeplyNestedStructPercolator, error) {
	root, err := parseFn(query)
	if err != nil {
		return nil, err
	}

	for _, modifier := range modifiers {
		root = modifier(root)
		if _, isEmpty := root.(*tree.Empty); isEmpty {
			// If the query is empty, we return a percolator that always returns false
			return &DeeplyNestedStructPercolator{fn: pDeeplyNestedStructFalseFn}, nil
		}
	}

	return &DeeplyNestedStructPercolator{fn: updateNodeDeeplyNestedStruct(root)}, nil
}

func updateNodeDeeplyNestedStruct(root tree.Expr) func(document *DeeplyNestedStruct) bool {
	switch r := root.(type) {
	case *tree.And:
		fns := make([]func(document *DeeplyNestedStruct) bool, 0, len(r.Children))
		for _, child := range r.Children {
			fn := updateNodeDeeplyNestedStruct(child)

			// Optimize the case where one of the children is always false
			if reflect.ValueOf(fn).Pointer() == reflect.ValueOf(pDeeplyNestedStructFalseFn).Pointer() {
				return pDeeplyNestedStructFalseFn
			}
			fns = append(fns, fn)
		}

		return func(document *DeeplyNestedStruct) bool {
			for _, fn := range fns {
				if !fn(document) {
					return false
				}
			}
			return true
		}
	case *tree.Or:
		fns := make([]func(document *DeeplyNestedStruct) bool, 0, len(r.Children))
		for _, child := range r.Children {
			fn := updateNodeDeeplyNestedStruct(child)

			// Optimize the case where one of the children is always false
			if reflect.ValueOf(fn).Pointer() == reflect.ValueOf(pDeeplyNestedStructFalseFn).Pointer() {
				continue
			}
			fns = append(fns, fn)
		}

		if len(fns) == 0 {
			return pDeeplyNestedStructFalseFn
		}

		return func(document *DeeplyNestedStruct) bool {
			for _, fn := range fns {
				if fn(document) {
					return true
				}
			}
			return false
		}
	case *tree.Not:
		child := updateNodeDeeplyNestedStruct(r.Child)
		return func(document *DeeplyNestedStruct) bool {
			return !child(document)
		}
	case *tree.Query:
		return pDeeplyNestedStructQuery(r)
	}

	return pDeeplyNestedStructFalseFn
}

type NestedStructPercolator struct {
	fn func(document *NestedStruct) bool
}

// Percolate percolates the document with the percolator's query.
// It will return a boolean indicating if the query matches the document.
func (p *NestedStructPercolator) Percolate(document *NestedStruct) bool {
	return p.fn(document)
}

// NewNestedStructPercolatorConstructor is a constructor of percolators.
// Every percolator it creates will use the parser provided to the constructor.
func NewNestedStructPercolatorConstructor(parseFn tree.ParseFn, modifiers ...tree.Modifiers) func(query string) (*NestedStructPercolator, error) {
	return func(query string) (*NestedStructPercolator, error) {
		return NewNestedStructPercolator(parseFn, query, modifiers...)
	}
}

// NewNestedStructPercolator creates a percolator with a given query and a given parser.
// It returns an error if the parsing failed.
func NewNestedStructPercolator(parseFn tree.ParseFn, query string, modifiers ...tree.Modifiers) (*NestedStructPercolator, error) {
	root, err := parseFn(query)
	if err != nil {
		return nil, err
	}

	for _, modifier := range modifiers {
		root = modifier(root)
		if _, isEmpty := root.(*tree.Empty); isEmpty {
			// If the query is empty, we return a percolator that always returns false
			return &NestedStructPercolator{fn: pNestedStructFalseFn}, nil
		}
	}

	return &NestedStructPercolator{fn: updateNodeNestedStruct(root)}, nil
}

func updateNodeNestedStruct(root tree.Expr) func(document *NestedStruct) bool {
	switch r := root.(type) {
	case *tree.And:
		fns := make([]func(document *NestedStruct) bool, 0, len(r.Children))
		for _, child := range r.Children {
			fn := updateNodeNestedStruct(child)

			// Optimize the case where one of the children is always false
			if reflect.ValueOf(fn).Pointer() == reflect.ValueOf(pNestedStructFalseFn).Pointer() {
				return pNestedStructFalseFn
			}
			fns = append(fns, fn)
		}

		return func(document *NestedStruct) bool {
			for _, fn := range fns {
				if !fn(document) {
					return false
				}
			}
			return true
		}
	case *tree.Or:
		fns := make([]func(document *NestedStruct) bool, 0, len(r.Children))
		for _, child := range r.Children {
			fn := updateNodeNestedStruct(child)

			// Optimize the case where one of the children is always false
			if reflect.ValueOf(fn).Pointer() == reflect.ValueOf(pNestedStructFalseFn).Pointer() {
				continue
			}
			fns = append(fns, fn)
		}

		if len(fns) == 0 {
			return pNestedStructFalseFn
		}

		return func(document *NestedStruct) bool {
			for _, fn := range fns {
				if fn(document) {
					return true
				}
			}
			return false
		}
	case *tree.Not:
		child := updateNodeNestedStruct(r.Child)
		return func(document *NestedStruct) bool {
			return !child(document)
		}
	case *tree.Query:
		return pNestedStructQuery(r)
	}

	return pNestedStructFalseFn
}

func pStructQuery(query *tree.Query) func(*Struct) bool {
	key, suffix, _ := strings.Cut(query.Key, ".")
	_ = suffix
	switch key {
	case "slice_string":
		parsed := query.Value
		switch query.Sign {
		case tree.Eq:
			return func(document *Struct) bool {
				for _, value := range document.SliceString {
					if value == parsed {
						return true
					}
				}
				return false
			}
		case tree.Gte:
			return func(document *Struct) bool {
				for _, value := range document.SliceString {
					if value >= parsed {
						return true
					}
				}
				return false
			}
		case tree.Gt:
			return func(document *Struct) bool {
				for _, value := range document.SliceString {
					if value > parsed {
						return true
					}
				}
				return false
			}
		case tree.Lte:
			return func(document *Struct) bool {
				for _, value := range document.SliceString {
					if value <= parsed {
						return true
					}
				}
				return false
			}
		case tree.Lt:
			return func(document *Struct) bool {
				for _, value := range document.SliceString {
					if value < parsed {
						return true
					}
				}
				return false
			}
		default:
			return pStructFalseFn
		}
	case "slice_int16":
		tmp, err := strconv.ParseInt(query.Value, 10, 16)
		if err != nil {
			return pStructFalseFn
		}
		parsed := int16(tmp)
		switch query.Sign {
		case tree.Eq:
			return func(document *Struct) bool {
				for _, value := range document.SliceInt16 {
					if value == parsed {
						return true
					}
				}
				return false
			}
		case tree.Gte:
			return func(document *Struct) bool {
				for _, value := range document.SliceInt16 {
					if value >= parsed {
						return true
					}
				}
				return false
			}
		case tree.Gt:
			return func(document *Struct) bool {
				for _, value := range document.SliceInt16 {
					if value > parsed {
						return true
					}
				}
				return false
			}
		case tree.Lte:
			return func(document *Struct) bool {
				for _, value := range document.SliceInt16 {
					if value <= parsed {
						return true
					}
				}
				return false
			}
		case tree.Lt:
			return func(document *Struct) bool {
				for _, value := range document.SliceInt16 {
					if value < parsed {
						return true
					}
				}
				return false
			}
		default:
			return pStructFalseFn
		}
	case "slice_bool":
		parsed, err := strconv.ParseBool(query.Value)
		if err != nil {
			return pStructFalseFn
		}
		return func(document *Struct) bool {
			for _, value := range document.SliceBool {
				if value == parsed {
					return true
				}
			}
			return false
		}
	case "slice_nested":
		fn := pNestedStructQuery(&tree1.Query{
			Key:   suffix,
			Sign:  query.Sign,
			Value: query.Value,
		})
		return func(document *Struct) bool {
			for _, value := range document.SliceNested {
				if fn(&value) {
					return true
				}
			}
			return false
		}
	case "slice_pointer_bool":
		parsed, err := strconv.ParseBool(query.Value)
		if err != nil {
			return pStructFalseFn
		}
		return func(document *Struct) bool {
			for _, value := range document.SlicePointerBool {
				if value != nil && *value == parsed {
					return true
				}
			}
			return false
		}
	case "slice_pointer_basic":
		parsed := query.Value
		switch query.Sign {
		case tree.Eq:
			return func(document *Struct) bool {
				for _, value := range document.SlicePointerBasic {
					if value != nil && *value == parsed {
						return true
					}
				}
				return false
			}
		case tree.Gte:
			return func(document *Struct) bool {
				for _, value := range document.SlicePointerBasic {
					if value != nil && *value >= parsed {
						return true
					}
				}
				return false
			}
		case tree.Gt:
			return func(document *Struct) bool {
				for _, value := range document.SlicePointerBasic {
					if value != nil && *value > parsed {
						return true
					}
				}
				return false
			}
		case tree.Lte:
			return func(document *Struct) bool {
				for _, value := range document.SlicePointerBasic {
					if value != nil && *value <= parsed {
						return true
					}
				}
				return false
			}
		case tree.Lt:
			return func(document *Struct) bool {
				for _, value := range document.SlicePointerBasic {
					if value != nil && *value < parsed {
						return true
					}
				}
				return false
			}
		default:
			return pStructFalseFn
		}
	case "slice_pointer_int16":
		tmp, err := strconv.ParseInt(query.Value, 10, 16)
		if err != nil {
			return pStructFalseFn
		}
		parsed := int16(tmp)
		switch query.Sign {
		case tree.Eq:
			return func(document *Struct) bool {
				for _, value := range document.SlicePointerInt16 {
					if value != nil && *value == parsed {
						return true
					}
				}
				return false
			}
		case tree.Gte:
			return func(document *Struct) bool {
				for _, value := range document.SlicePointerInt16 {
					if value != nil && *value >= parsed {
						return true
					}
				}
				return false
			}
		case tree.Gt:
			return func(document *Struct) bool {
				for _, value := range document.SlicePointerInt16 {
					if value != nil && *value > parsed {
						return true
					}
				}
				return false
			}
		case tree.Lte:
			return func(document *Struct) bool {
				for _, value := range document.SlicePointerInt16 {
					if value != nil && *value <= parsed {
						return true
					}
				}
				return false
			}
		case tree.Lt:
			return func(document *Struct) bool {
				for _, value := range document.SlicePointerInt16 {
					if value != nil && *value < parsed {
						return true
					}
				}
				return false
			}
		default:
			return pStructFalseFn
		}
	case "slice_pointer_nested":
		fn := pNestedStructQuery(&tree1.Query{
			Key:   suffix,
			Sign:  query.Sign,
			Value: query.Value,
		})
		return func(document *Struct) bool {
			for _, value := range document.SlicePointerNested {
				if fn(value) {
					return true
				}
			}
			return false
		}
	default:
		return pStructFalseFn
	}
}

func pStructFalseFn(_ *Struct) bool {
	return false
}

func pDeeplyNestedStructQuery(query *tree.Query) func(*DeeplyNestedStruct) bool {
	key, suffix, _ := strings.Cut(query.Key, ".")
	_ = suffix
	switch key {
	case "field":
		parsed := query.Value
		return func(document *DeeplyNestedStruct) bool {
			return document.Field == parsed
		}
	default:
		return pDeeplyNestedStructFalseFn
	}
}

func pDeeplyNestedStructFalseFn(_ *DeeplyNestedStruct) bool {
	return false
}

func pNestedStructQuery(query *tree.Query) func(*NestedStruct) bool {
	key, suffix, _ := strings.Cut(query.Key, ".")
	_ = suffix
	switch key {
	case "value":
		parsed := query.Value
		return func(document *NestedStruct) bool {
			return document.Value == parsed
		}
	case "nested_value":
		fn := pDeeplyNestedStructQuery(&tree1.Query{
			Key:   suffix,
			Sign:  query.Sign,
			Value: query.Value,
		})
		return func(document *NestedStruct) bool {
			return fn(&document.NestedValue)
		}
	default:
		return pNestedStructFalseFn
	}
}

func pNestedStructFalseFn(_ *NestedStruct) bool {
	return false
}
