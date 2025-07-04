// Code generated by github.com/mchenriques22/pergolator, DO NOT EDIT.
package misc

import (
	"reflect"
	"time"

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

type FlattenRecStructPercolator struct {
	fn func(document *FlattenRecStruct) bool
}

// Percolate percolates the document with the percolator's query.
// It will return a boolean indicating if the query matches the document.
func (p *FlattenRecStructPercolator) Percolate(document *FlattenRecStruct) bool {
	return p.fn(document)
}

// NewFlattenRecStructPercolatorConstructor is a constructor of percolators.
// Every percolator it creates will use the parser provided to the constructor.
func NewFlattenRecStructPercolatorConstructor(parseFn tree.ParseFn, modifiers ...tree.Modifiers) func(query string) (*FlattenRecStructPercolator, error) {
	return func(query string) (*FlattenRecStructPercolator, error) {
		return NewFlattenRecStructPercolator(parseFn, query, modifiers...)
	}
}

// NewFlattenRecStructPercolator creates a percolator with a given query and a given parser.
// It returns an error if the parsing failed.
func NewFlattenRecStructPercolator(parseFn tree.ParseFn, query string, modifiers ...tree.Modifiers) (*FlattenRecStructPercolator, error) {
	root, err := parseFn(query)
	if err != nil {
		return nil, err
	}

	for _, modifier := range modifiers {
		root = modifier(root)
		if _, isEmpty := root.(*tree.Empty); isEmpty {
			// If the query is empty, we return a percolator that always returns false
			return &FlattenRecStructPercolator{fn: pFlattenRecStructFalseFn}, nil
		}
	}

	return &FlattenRecStructPercolator{fn: updateNodeFlattenRecStruct(root)}, nil
}

func updateNodeFlattenRecStruct(root tree.Expr) func(document *FlattenRecStruct) bool {
	switch r := root.(type) {
	case *tree.And:
		fns := make([]func(document *FlattenRecStruct) bool, 0, len(r.Children))
		for _, child := range r.Children {
			fn := updateNodeFlattenRecStruct(child)

			// Optimize the case where one of the children is always false
			if reflect.ValueOf(fn).Pointer() == reflect.ValueOf(pFlattenRecStructFalseFn).Pointer() {
				return pFlattenRecStructFalseFn
			}
			fns = append(fns, fn)
		}

		return func(document *FlattenRecStruct) bool {
			for _, fn := range fns {
				if !fn(document) {
					return false
				}
			}
			return true
		}
	case *tree.Or:
		fns := make([]func(document *FlattenRecStruct) bool, 0, len(r.Children))
		for _, child := range r.Children {
			fn := updateNodeFlattenRecStruct(child)

			// Optimize the case where one of the children is always false
			if reflect.ValueOf(fn).Pointer() == reflect.ValueOf(pFlattenRecStructFalseFn).Pointer() {
				continue
			}
			fns = append(fns, fn)
		}

		if len(fns) == 0 {
			return pFlattenRecStructFalseFn
		}

		return func(document *FlattenRecStruct) bool {
			for _, fn := range fns {
				if fn(document) {
					return true
				}
			}
			return false
		}
	case *tree.Not:
		child := updateNodeFlattenRecStruct(r.Child)
		return func(document *FlattenRecStruct) bool {
			return !child(document)
		}
	case *tree.Query:
		return pFlattenRecStructQuery(r)
	}

	return pFlattenRecStructFalseFn
}

type InlineStringPercolator struct {
	fn func(document *InlineString) bool
}

// Percolate percolates the document with the percolator's query.
// It will return a boolean indicating if the query matches the document.
func (p *InlineStringPercolator) Percolate(document *InlineString) bool {
	return p.fn(document)
}

// NewInlineStringPercolatorConstructor is a constructor of percolators.
// Every percolator it creates will use the parser provided to the constructor.
func NewInlineStringPercolatorConstructor(parseFn tree.ParseFn, modifiers ...tree.Modifiers) func(query string) (*InlineStringPercolator, error) {
	return func(query string) (*InlineStringPercolator, error) {
		return NewInlineStringPercolator(parseFn, query, modifiers...)
	}
}

// NewInlineStringPercolator creates a percolator with a given query and a given parser.
// It returns an error if the parsing failed.
func NewInlineStringPercolator(parseFn tree.ParseFn, query string, modifiers ...tree.Modifiers) (*InlineStringPercolator, error) {
	root, err := parseFn(query)
	if err != nil {
		return nil, err
	}

	for _, modifier := range modifiers {
		root = modifier(root)
		if _, isEmpty := root.(*tree.Empty); isEmpty {
			// If the query is empty, we return a percolator that always returns false
			return &InlineStringPercolator{fn: pInlineStringFalseFn}, nil
		}
	}

	return &InlineStringPercolator{fn: updateNodeInlineString(root)}, nil
}

func updateNodeInlineString(root tree.Expr) func(document *InlineString) bool {
	switch r := root.(type) {
	case *tree.And:
		fns := make([]func(document *InlineString) bool, 0, len(r.Children))
		for _, child := range r.Children {
			fn := updateNodeInlineString(child)

			// Optimize the case where one of the children is always false
			if reflect.ValueOf(fn).Pointer() == reflect.ValueOf(pInlineStringFalseFn).Pointer() {
				return pInlineStringFalseFn
			}
			fns = append(fns, fn)
		}

		return func(document *InlineString) bool {
			for _, fn := range fns {
				if !fn(document) {
					return false
				}
			}
			return true
		}
	case *tree.Or:
		fns := make([]func(document *InlineString) bool, 0, len(r.Children))
		for _, child := range r.Children {
			fn := updateNodeInlineString(child)

			// Optimize the case where one of the children is always false
			if reflect.ValueOf(fn).Pointer() == reflect.ValueOf(pInlineStringFalseFn).Pointer() {
				continue
			}
			fns = append(fns, fn)
		}

		if len(fns) == 0 {
			return pInlineStringFalseFn
		}

		return func(document *InlineString) bool {
			for _, fn := range fns {
				if fn(document) {
					return true
				}
			}
			return false
		}
	case *tree.Not:
		child := updateNodeInlineString(r.Child)
		return func(document *InlineString) bool {
			return !child(document)
		}
	case *tree.Query:
		return pInlineStringQuery(r)
	}

	return pInlineStringFalseFn
}

type InlineUint64Percolator struct {
	fn func(document *InlineUint64) bool
}

// Percolate percolates the document with the percolator's query.
// It will return a boolean indicating if the query matches the document.
func (p *InlineUint64Percolator) Percolate(document *InlineUint64) bool {
	return p.fn(document)
}

// NewInlineUint64PercolatorConstructor is a constructor of percolators.
// Every percolator it creates will use the parser provided to the constructor.
func NewInlineUint64PercolatorConstructor(parseFn tree.ParseFn, modifiers ...tree.Modifiers) func(query string) (*InlineUint64Percolator, error) {
	return func(query string) (*InlineUint64Percolator, error) {
		return NewInlineUint64Percolator(parseFn, query, modifiers...)
	}
}

// NewInlineUint64Percolator creates a percolator with a given query and a given parser.
// It returns an error if the parsing failed.
func NewInlineUint64Percolator(parseFn tree.ParseFn, query string, modifiers ...tree.Modifiers) (*InlineUint64Percolator, error) {
	root, err := parseFn(query)
	if err != nil {
		return nil, err
	}

	for _, modifier := range modifiers {
		root = modifier(root)
		if _, isEmpty := root.(*tree.Empty); isEmpty {
			// If the query is empty, we return a percolator that always returns false
			return &InlineUint64Percolator{fn: pInlineUint64FalseFn}, nil
		}
	}

	return &InlineUint64Percolator{fn: updateNodeInlineUint64(root)}, nil
}

func updateNodeInlineUint64(root tree.Expr) func(document *InlineUint64) bool {
	switch r := root.(type) {
	case *tree.And:
		fns := make([]func(document *InlineUint64) bool, 0, len(r.Children))
		for _, child := range r.Children {
			fn := updateNodeInlineUint64(child)

			// Optimize the case where one of the children is always false
			if reflect.ValueOf(fn).Pointer() == reflect.ValueOf(pInlineUint64FalseFn).Pointer() {
				return pInlineUint64FalseFn
			}
			fns = append(fns, fn)
		}

		return func(document *InlineUint64) bool {
			for _, fn := range fns {
				if !fn(document) {
					return false
				}
			}
			return true
		}
	case *tree.Or:
		fns := make([]func(document *InlineUint64) bool, 0, len(r.Children))
		for _, child := range r.Children {
			fn := updateNodeInlineUint64(child)

			// Optimize the case where one of the children is always false
			if reflect.ValueOf(fn).Pointer() == reflect.ValueOf(pInlineUint64FalseFn).Pointer() {
				continue
			}
			fns = append(fns, fn)
		}

		if len(fns) == 0 {
			return pInlineUint64FalseFn
		}

		return func(document *InlineUint64) bool {
			for _, fn := range fns {
				if fn(document) {
					return true
				}
			}
			return false
		}
	case *tree.Not:
		child := updateNodeInlineUint64(r.Child)
		return func(document *InlineUint64) bool {
			return !child(document)
		}
	case *tree.Query:
		return pInlineUint64Query(r)
	}

	return pInlineUint64FalseFn
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

type TimePercolator struct {
	fn func(document *time.Time) bool
}

// Percolate percolates the document with the percolator's query.
// It will return a boolean indicating if the query matches the document.
func (p *TimePercolator) Percolate(document *time.Time) bool {
	return p.fn(document)
}

// NewTimePercolatorConstructor is a constructor of percolators.
// Every percolator it creates will use the parser provided to the constructor.
func NewTimePercolatorConstructor(parseFn tree.ParseFn, modifiers ...tree.Modifiers) func(query string) (*TimePercolator, error) {
	return func(query string) (*TimePercolator, error) {
		return NewTimePercolator(parseFn, query, modifiers...)
	}
}

// NewTimePercolator creates a percolator with a given query and a given parser.
// It returns an error if the parsing failed.
func NewTimePercolator(parseFn tree.ParseFn, query string, modifiers ...tree.Modifiers) (*TimePercolator, error) {
	root, err := parseFn(query)
	if err != nil {
		return nil, err
	}

	for _, modifier := range modifiers {
		root = modifier(root)
		if _, isEmpty := root.(*tree.Empty); isEmpty {
			// If the query is empty, we return a percolator that always returns false
			return &TimePercolator{fn: pTimeFalseFn}, nil
		}
	}

	return &TimePercolator{fn: updateNodeTime(root)}, nil
}

func updateNodeTime(root tree.Expr) func(document *time.Time) bool {
	switch r := root.(type) {
	case *tree.And:
		fns := make([]func(document *time.Time) bool, 0, len(r.Children))
		for _, child := range r.Children {
			fn := updateNodeTime(child)

			// Optimize the case where one of the children is always false
			if reflect.ValueOf(fn).Pointer() == reflect.ValueOf(pTimeFalseFn).Pointer() {
				return pTimeFalseFn
			}
			fns = append(fns, fn)
		}

		return func(document *time.Time) bool {
			for _, fn := range fns {
				if !fn(document) {
					return false
				}
			}
			return true
		}
	case *tree.Or:
		fns := make([]func(document *time.Time) bool, 0, len(r.Children))
		for _, child := range r.Children {
			fn := updateNodeTime(child)

			// Optimize the case where one of the children is always false
			if reflect.ValueOf(fn).Pointer() == reflect.ValueOf(pTimeFalseFn).Pointer() {
				continue
			}
			fns = append(fns, fn)
		}

		if len(fns) == 0 {
			return pTimeFalseFn
		}

		return func(document *time.Time) bool {
			for _, fn := range fns {
				if fn(document) {
					return true
				}
			}
			return false
		}
	case *tree.Not:
		child := updateNodeTime(r.Child)
		return func(document *time.Time) bool {
			return !child(document)
		}
	case *tree.Query:
		return pTimeQuery(r)
	}

	return pTimeFalseFn
}

func pStructQuery(query *tree.Query) func(*Struct) bool {
	key, suffix, _ := strings.Cut(query.Key, ".")
	_ = suffix
	switch key {
	case "nested":
		fn := pNestedStructQuery(&tree1.Query{
			Key:   suffix,
			Sign:  query.Sign,
			Value: query.Value,
		})
		return func(document *Struct) bool {
			return fn(&document.Nested)
		}
	case "inline_string":
		fn := pInlineStringQuery(&tree1.Query{
			Key:   suffix,
			Sign:  query.Sign,
			Value: query.Value,
		})
		return func(document *Struct) bool {
			return fn(&document.InlineString)
		}
	case "inline_uint64":
		fn := pInlineUint64Query(&tree1.Query{
			Key:   suffix,
			Sign:  query.Sign,
			Value: query.Value,
		})
		return func(document *Struct) bool {
			return fn(&document.InlineUint64)
		}
	case "pointer_string":
		parsed := query.Value
		return func(document *Struct) bool {
			return document.PointerString != nil && *document.PointerString == parsed
		}
	case "pointer_nested":
		fn := pNestedStructQuery(&tree1.Query{
			Key:   suffix,
			Sign:  query.Sign,
			Value: query.Value,
		})
		return func(document *Struct) bool {
			return fn(document.PointerNested)
		}
	case "pointer_bool":
		parsed, err := strconv.ParseBool(query.Value)
		if err != nil {
			return pStructFalseFn
		}
		return func(document *Struct) bool {
			return document.PointerBool != nil && *document.PointerBool == parsed
		}
	case "time":
		fn := pTimeQuery(&tree1.Query{
			Key:   suffix,
			Sign:  query.Sign,
			Value: query.Value,
		})
		return func(document *Struct) bool {
			return fn(&document.Time)
		}
	default:
		fnFlattenNested := pNestedStructQuery(query)
		fnFlattenNestedEnabled := reflect.ValueOf(fnFlattenNested).Pointer() != reflect.ValueOf(pNestedStructFalseFn).Pointer()
		fnFlattenDeeplyNested := pDeeplyNestedStructQuery(query)
		fnFlattenDeeplyNestedEnabled := reflect.ValueOf(fnFlattenDeeplyNested).Pointer() != reflect.ValueOf(pDeeplyNestedStructFalseFn).Pointer()
		return func(document *Struct) bool {
			return fnFlattenNestedEnabled && fnFlattenNested(&document.FlattenNested) || fnFlattenDeeplyNestedEnabled && fnFlattenDeeplyNested(&document.FlattenDeeplyNested)
		}
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

func pFlattenRecStructQuery(query *tree.Query) func(*FlattenRecStruct) bool {
	key, suffix, _ := strings.Cut(query.Key, ".")
	_ = suffix
	switch key {
	case "another_field":
		parsed := query.Value
		return func(document *FlattenRecStruct) bool {
			return document.AnotherField == parsed
		}
	default:
		return pFlattenRecStructFalseFn
	}
}

func pFlattenRecStructFalseFn(_ *FlattenRecStruct) bool {
	return false
}

func pInlineStringQuery(query *tree.Query) func(*InlineString) bool {
	key, suffix, _ := strings.Cut(query.Key, ".")
	_ = suffix
	switch key {
	case "":
		parsed := query.Value
		converted := InlineString(parsed)
		return func(document *InlineString) bool {
			return *document == converted
		}
	default:
		return pInlineStringFalseFn
	}
}

func pInlineStringFalseFn(_ *InlineString) bool {
	return false
}

func pInlineUint64Query(query *tree.Query) func(*InlineUint64) bool {
	key, suffix, _ := strings.Cut(query.Key, ".")
	_ = suffix
	switch key {
	case "":
		parsed, err := strconv.ParseUint(query.Value, 10, 64)
		if err != nil {
			return pInlineUint64FalseFn
		}
		converted := InlineUint64(parsed)
		return func(document *InlineUint64) bool {
			return *document == converted
		}
	default:
		return pInlineUint64FalseFn
	}
}

func pInlineUint64FalseFn(_ *InlineUint64) bool {
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
		fnFlattenRecStruct := pFlattenRecStructQuery(query)
		fnFlattenRecStructEnabled := reflect.ValueOf(fnFlattenRecStruct).Pointer() != reflect.ValueOf(pFlattenRecStructFalseFn).Pointer()
		return func(document *NestedStruct) bool {
			return fnFlattenRecStructEnabled && fnFlattenRecStruct(&document.FlattenRecStruct)
		}
	}
}

func pNestedStructFalseFn(_ *NestedStruct) bool {
	return false
}

func pTimeQuery(query *tree1.Query) func(*time.Time) bool {
	parsed, err := time.Parse(time.RFC3339, query.Value)
	if err != nil {
		return pTimeFalseFn
	}
	switch query.Sign {
	case tree.Eq:
		return func(t *time.Time) bool {
			return t.Equal(parsed)
		}
	case tree.Lt:
		return func(t *time.Time) bool {
			return t.Before(parsed)
		}
	case tree.Lte:
		return func(t *time.Time) bool {
			return t.Before(parsed) || t.Equal(parsed)
		}
	case tree.Gt:
		return func(t *time.Time) bool {
			return parsed.Before(*t)
		}

	case tree.Gte:
		return func(t *time.Time) bool {
			return parsed.Before(*t) || parsed.Equal(*t)
		}
	}
	return pTimeFalseFn
}

func pTimeFalseFn(_ *time.Time) bool {
	return false
}
